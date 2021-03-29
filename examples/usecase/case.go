package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	//"gopkg.in/yaml.v2"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// private yaml
type AccountConfig struct {
	GrantType    string `yaml:"grantType"`
	Audience     string `yaml:"audience"`
	TenantUUID   string `yaml:"tenantUUID"`
	ClientId     string `yaml:"clientId"`
	ClientSecret string `yaml:"clientSecret"`
}

func NewConfigFromFile(path string) *AccountConfig {
	yamlConfig, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	a := new(AccountConfig)
	err = yaml.Unmarshal(yamlConfig, a)
	if err != nil {
		panic(err)
	}

	return a
}

type RequestTokenBody struct {
	GrantType    string `json:"grant_type"`
	Audience     string `json:"audience"`
	TenantUUID   string `json:"tenantUUID"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
	Scope       string `json:"scope"`
}

/// Ledger

type LedgerAccount struct {
	Id     int    `json:"_id"`
	Name   string `json:"name"`
	Number string `json:"number"`
}

type LedgerAccountResponse struct {
	TotalItems int             `json:"totalItems"`
	Data       []LedgerAccount `json:"data"`
}

type Entry struct {
	AccountName   string
	AccountNumber string
	Debit         string
	Credit        string
}

type Group struct {
	Entries  []*Entry
	Len      int
	Currency string
	Resident string
	Order    int
}

// Create journal
type CreateJournalRequest struct {
	Status       string        `json:"status"`
	EntryType    string        `json:"entryType"`
	SourceLedger string        `json:"sourceLedger"`
	Reference    string        `json:"reference"`
	Currency     string        `json:"currency"`
	Notes        string        `json:"notes"`
	Transactions []Transaction `json:"Transactions"`
}

type Transaction struct {
	Description     string `json:"description"`
	Debit           string `json:"debit"`
	Credit          string `json:"credit"`
	TransactionDate string `json:"transactionDate"`
	PostedDate      string `json:"postedDate"`
	LedgerAccountId int    `json:"LedgerAccountId"`
	LocationId      int    `json:"LocationId"`
}

// Cache LedgerAccounts
var ledgerAccounts map[string]*LedgerAccount

var accessToken string
// var accessToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ik16WTFPRVEzUWpZMk1EazVPRGxCTWtJMk9FTXhOVGxEUmpVeFJVWkJRa1pDUTBZek5UY3dOQSJ9.eyJodHRwczovL2FwaS5zb2Z0bGVkZ2VyLmNvbS90ZW5hbnRVVUlEIjoiM2EwNmIxODctNjI3ZC00ZmIwLWFjZWItZDhjNDRjODY0NTNlIiwiaHR0cHM6Ly9hcGkuc29mdGxlZGdlci5jb20vY2xpZW50TmFtZSI6Ik1hdHJpeCAtIFNhbmRib3giLCJpc3MiOiJodHRwczovL3NvZnRsZWRnZXIuZXUuYXV0aDAuY29tLyIsInN1YiI6IjFMSktiRGxxMlIwSVkzaEQzNm16ZFFNV2V4RzZjbTZhQGNsaWVudHMiLCJhdWQiOiJodHRwczovL2FwaS5zb2Z0bGVkZ2VyLmNvbSIsImlhdCI6MTYxNjk5Mjk4MSwiZXhwIjoxNjE3MDc5MzgxLCJhenAiOiIxTEpLYkRscTJSMElZM2hEMzZtemRRTVdleEc2Y202YSIsInNjb3BlIjoiY3J5cHRvIGNyZWF0ZTpDb3N0Q2VudGVyIGNyZWF0ZTpKb3VybmFsRW50cnkgY3JlYXRlOkxlZGdlckFjY291bnQgY3JlYXRlOkxvY2F0aW9uIGNyZWF0ZTpQcm9kdWN0IGRlbGV0ZTpDb3N0Q2VudGVyIGRlbGV0ZTpMZWRnZXJBY2NvdW50IGRlbGV0ZTpMb2NhdGlvbiBkZWxldGU6UHJvZHVjdCBsaXN0OkNvc3RDZW50ZXJzIGxpc3Q6TGVkZ2VyQWNjb3VudHMgbGlzdDpMb2NhdGlvbnMgbGlzdDpQcm9kdWN0cyB1cGRhdGU6Q29zdENlbnRlciB1cGRhdGU6TGVkZ2VyQWNjb3VudCB1cGRhdGU6TG9jYXRpb24gdXBkYXRlOlByb2R1Y3QgY3JlYXRlOkJpbGwgY3JlYXRlOlZlbmRvciBkZWxldGU6QmlsbCBkZWxldGU6VmVuZG9yIGxpc3Q6QmlsbHMgbGlzdDpWZW5kb3JzIHBheTpCaWxscyB1cGRhdGU6QmlsbCB1cGRhdGU6VmVuZG9yIHZpZXc6QVBBZ2luZyBjcmVhdGU6QmFua0FjY291bnQgY3JlYXRlOkNhc2hSZWNlaXB0IGRlbGV0ZTpCYW5rQWNjb3VudCBsaXN0OkJhbmtBY2NvdW50cyBsaXN0OkNhc2hSZWNlaXB0cyByZWNvbmNpbGU6QmFua0FjY291bnQgdXBkYXRlOkNhc2hSZWNlaXB0IGFwcGx5OkNhc2hSZWNlaXB0IGNyZWF0ZTpDdXN0b21lciBjcmVhdGU6SW52b2ljZSBkZWxldGU6Q3VzdG9tZXIgZGVsZXRlOkludm9pY2UgaXNzdWU6SW52b2ljZSBsaXN0OkN1c3RvbWVycyBsaXN0Okludm9pY2VzIHVwZGF0ZTpDdXN0b21lciB1cGRhdGU6SW52b2ljZSB2aWV3OkFSQWdpbmcgdm9pZDpJbnZvaWNlIHZpZXc6UHJvZmlsZSB1cGRhdGU6RmluYW5jaWFscyB2aWV3OkRhc2hib2FyZCB2aWV3OkZpbmFuY2lhbHMgdmlldzpHZW5lcmFsTGVkZ2VyIHZpZXc6Sm91cm5hbFJlcG9ydCB2aWV3OkxlZGdlcnMgY2xvc2U6TGVkZ2VycyBsaXN0OkF1ZGl0TG9ncyBsaXN0OlNldHRpbmdzIGxpc3Q6VXNlcnMgdXBkYXRlOlNldHRpbmdzIHVwZGF0ZTpVc2VycyBhY2NlcHQ6U2FsZXNRdW90ZSBhZGp1c3Q6U3RvY2sgY3JlYXRlOkl0ZW0gY3JlYXRlOlB1cmNoYXNlT3JkZXIgY3JlYXRlOlNhbGVzT3JkZXIgY3JlYXRlOldhcmVob3VzZSBkZWxldGU6SXRlbSBkZWxldGU6UHVyY2hhc2VPcmRlciBkZWxldGU6U2FsZXNPcmRlciBkZWxldGU6V2FyZWhvdXNlIGZ1bGZpbGw6SXRlbSBpc3N1ZTpQdXJjaGFzZU9yZGVyIGlzc3VlOlNhbGVzUXVvdGUgbGlzdDpJdGVtcyBsaXN0OlB1cmNoYXNlT3JkZXJzIGxpc3Q6U2FsZXNPcmRlcnMgbGlzdDpXYXJlaG91c2VzIHJlY2VpdmU6SXRlbXMgcmVqZWN0OlNhbGVzUXVvdGUgdXBkYXRlOkl0ZW0gdXBkYXRlOlB1cmNoYXNlT3JkZXIgdXBkYXRlOlNhbGVzT3JkZXIgdXBkYXRlOldhcmVob3VzZSBsaXN0OkpvYnMgY3JlYXRlOkpvYnMgY3JlYXRlOkpvYiB1cGRhdGU6Sm9iIGRlbGV0ZTpKb2IgYXBwcm92ZTpKb3VybmFsRW50cnkgdXBkYXRlOkpvdXJuYWxFbnRyeSBkZWxldGU6Sm91cm5hbEVudHJ5IHVwZGF0ZTpQcm9kdWN0aW9uIGxpc3Q6UHJvZHVjdGlvbnMgZGVsZXRlOlByb2R1Y3Rpb24gY3JlYXRlOlByb2R1Y3Rpb24gbGlzdDpPdmVyaGVhZHMgdXBkYXRlOk92ZXJoZWFkIGRlbGV0ZTpPdmVyaGVhZCBjcmVhdGU6T3ZlcmhlYWQgbGlzdDpQYXltZW50cyBjcmVhdGU6UGF5bWVudCB1cGRhdGU6UGF5bWVudCBhcHByb3ZlOlBheW1lbnQgdm9pZDpQYXltZW50IGRlbGV0ZTpQYXltZW50IGFwcHJvdmU6QmlsbCB2aWV3OkNyeXB0b0ludGVncmF0aW9ucyBjb21wbGV0ZTpQcm9kdWN0aW9uIiwiZ3R5IjoiY2xpZW50LWNyZWRlbnRpYWxzIiwicGVybWlzc2lvbnMiOlsiY3J5cHRvIiwiY3JlYXRlOkNvc3RDZW50ZXIiLCJjcmVhdGU6Sm91cm5hbEVudHJ5IiwiY3JlYXRlOkxlZGdlckFjY291bnQiLCJjcmVhdGU6TG9jYXRpb24iLCJjcmVhdGU6UHJvZHVjdCIsImRlbGV0ZTpDb3N0Q2VudGVyIiwiZGVsZXRlOkxlZGdlckFjY291bnQiLCJkZWxldGU6TG9jYXRpb24iLCJkZWxldGU6UHJvZHVjdCIsImxpc3Q6Q29zdENlbnRlcnMiLCJsaXN0OkxlZGdlckFjY291bnRzIiwibGlzdDpMb2NhdGlvbnMiLCJsaXN0OlByb2R1Y3RzIiwidXBkYXRlOkNvc3RDZW50ZXIiLCJ1cGRhdGU6TGVkZ2VyQWNjb3VudCIsInVwZGF0ZTpMb2NhdGlvbiIsInVwZGF0ZTpQcm9kdWN0IiwiY3JlYXRlOkJpbGwiLCJjcmVhdGU6VmVuZG9yIiwiZGVsZXRlOkJpbGwiLCJkZWxldGU6VmVuZG9yIiwibGlzdDpCaWxscyIsImxpc3Q6VmVuZG9ycyIsInBheTpCaWxscyIsInVwZGF0ZTpCaWxsIiwidXBkYXRlOlZlbmRvciIsInZpZXc6QVBBZ2luZyIsImNyZWF0ZTpCYW5rQWNjb3VudCIsImNyZWF0ZTpDYXNoUmVjZWlwdCIsImRlbGV0ZTpCYW5rQWNjb3VudCIsImxpc3Q6QmFua0FjY291bnRzIiwibGlzdDpDYXNoUmVjZWlwdHMiLCJyZWNvbmNpbGU6QmFua0FjY291bnQiLCJ1cGRhdGU6Q2FzaFJlY2VpcHQiLCJhcHBseTpDYXNoUmVjZWlwdCIsImNyZWF0ZTpDdXN0b21lciIsImNyZWF0ZTpJbnZvaWNlIiwiZGVsZXRlOkN1c3RvbWVyIiwiZGVsZXRlOkludm9pY2UiLCJpc3N1ZTpJbnZvaWNlIiwibGlzdDpDdXN0b21lcnMiLCJsaXN0Okludm9pY2VzIiwidXBkYXRlOkN1c3RvbWVyIiwidXBkYXRlOkludm9pY2UiLCJ2aWV3OkFSQWdpbmciLCJ2b2lkOkludm9pY2UiLCJ2aWV3OlByb2ZpbGUiLCJ1cGRhdGU6RmluYW5jaWFscyIsInZpZXc6RGFzaGJvYXJkIiwidmlldzpGaW5hbmNpYWxzIiwidmlldzpHZW5lcmFsTGVkZ2VyIiwidmlldzpKb3VybmFsUmVwb3J0IiwidmlldzpMZWRnZXJzIiwiY2xvc2U6TGVkZ2VycyIsImxpc3Q6QXVkaXRMb2dzIiwibGlzdDpTZXR0aW5ncyIsImxpc3Q6VXNlcnMiLCJ1cGRhdGU6U2V0dGluZ3MiLCJ1cGRhdGU6VXNlcnMiLCJhY2NlcHQ6U2FsZXNRdW90ZSIsImFkanVzdDpTdG9jayIsImNyZWF0ZTpJdGVtIiwiY3JlYXRlOlB1cmNoYXNlT3JkZXIiLCJjcmVhdGU6U2FsZXNPcmRlciIsImNyZWF0ZTpXYXJlaG91c2UiLCJkZWxldGU6SXRlbSIsImRlbGV0ZTpQdXJjaGFzZU9yZGVyIiwiZGVsZXRlOlNhbGVzT3JkZXIiLCJkZWxldGU6V2FyZWhvdXNlIiwiZnVsZmlsbDpJdGVtIiwiaXNzdWU6UHVyY2hhc2VPcmRlciIsImlzc3VlOlNhbGVzUXVvdGUiLCJsaXN0Okl0ZW1zIiwibGlzdDpQdXJjaGFzZU9yZGVycyIsImxpc3Q6U2FsZXNPcmRlcnMiLCJsaXN0OldhcmVob3VzZXMiLCJyZWNlaXZlOkl0ZW1zIiwicmVqZWN0OlNhbGVzUXVvdGUiLCJ1cGRhdGU6SXRlbSIsInVwZGF0ZTpQdXJjaGFzZU9yZGVyIiwidXBkYXRlOlNhbGVzT3JkZXIiLCJ1cGRhdGU6V2FyZWhvdXNlIiwibGlzdDpKb2JzIiwiY3JlYXRlOkpvYnMiLCJjcmVhdGU6Sm9iIiwidXBkYXRlOkpvYiIsImRlbGV0ZTpKb2IiLCJhcHByb3ZlOkpvdXJuYWxFbnRyeSIsInVwZGF0ZTpKb3VybmFsRW50cnkiLCJkZWxldGU6Sm91cm5hbEVudHJ5IiwidXBkYXRlOlByb2R1Y3Rpb24iLCJsaXN0OlByb2R1Y3Rpb25zIiwiZGVsZXRlOlByb2R1Y3Rpb24iLCJjcmVhdGU6UHJvZHVjdGlvbiIsImxpc3Q6T3ZlcmhlYWRzIiwidXBkYXRlOk92ZXJoZWFkIiwiZGVsZXRlOk92ZXJoZWFkIiwiY3JlYXRlOk92ZXJoZWFkIiwibGlzdDpQYXltZW50cyIsImNyZWF0ZTpQYXltZW50IiwidXBkYXRlOlBheW1lbnQiLCJhcHByb3ZlOlBheW1lbnQiLCJ2b2lkOlBheW1lbnQiLCJkZWxldGU6UGF5bWVudCIsImFwcHJvdmU6QmlsbCIsInZpZXc6Q3J5cHRvSW50ZWdyYXRpb25zIiwiY29tcGxldGU6UHJvZHVjdGlvbiJdfQ.ohnr3HzsrnIm2vhPSREa87_RkkTb8T5S60wSM4VSVoNzT02Rm3dWr78DzFZT9pU2HYv2UlHpxosmfSoxNQOVn84u2xAyngtktqKG81Tijp5CX0QMK61cwiE6KrqmhLteDDcpB7SkjiC_Rz6WV8D4naAKpUKgog6GEWfJBqy85zvGD3vKSWNljhq6hCZ-KTJb3kfj0IRB97yh2dRdU11Dm7OrxY15_GT5j54tfoDoImiTXYYNJQ2cEenB06fptfH0hLLg-0SeWHlP1ioguQIiEqSHIz5ycBOYTgOKG2ibNsxlt-VN9jFrjwyAsTSpGnJSke6WwXscJd_yZnmqv-QV0g"

const (
	NonUAE = "LUX"
	UAE    = "UAE"
)

func bootstrap() {
	accountConfig := NewConfigFromFile("/Users/dfg/workspace/go/src/github.com/shniu/gostuff/examples/usecase/secret.priv.yaml")

	// 解析 excel 表格，获取到原始数据集
	fmt.Println("=> Step 1. Parse excel files to get the original data set")
	groups := parseExcel()
	printGroups(groups)

	//if true {
	//	panic("Fail fast")
	//}

	client := &http.Client{}

	// 获取 API Token
	fmt.Println("=> Step 2. Request an API Token")
	tokenResponse := prepareAccessToken(client, accountConfig)

	// 使用 Token 请求到所有的 LedgerAccount
	ledgerAccounts = getAllLedgerAccounts(client, tokenResponse)
	for i := range ledgerAccounts {
		fmt.Println(i, ledgerAccounts[i])
	}

	var success, failure = 0, 0

	// 遍历原始数据集，转换每一组消息为请求 body，并发送请求
	fmt.Println("=> Step 3. Build create journal request, then sending it")
	var reference = 20000001
	for i := range groups {

		valid := true
		journalReq := initCreateJournalRequest(strconv.Itoa(reference), groups[i].Currency, "102439")
		for j := range groups[i].Entries {
			// fmt.Println("***", groups[i].Currency, " order is ", groups[i].Order, groups[i].Entries[j])
			t := Transaction{}
			t.LocationId = 19
			t.Description = groups[i].Resident

			if groups[i].Entries[j].Debit == "" {
				t.Debit = "0.00"
			} else {
				t.Debit = groups[i].Entries[j].Debit
			}

			if groups[i].Entries[j].Credit == "" {
				t.Credit = "0.00"
			} else {
				t.Credit = groups[i].Entries[j].Credit
			}

			t.TransactionDate = "2021-03-26"
			t.PostedDate = "2021-03-27"

			accountNumber := groups[i].Entries[j].AccountNumber
			ledgerAccount, ok := ledgerAccounts[accountNumber]
			if !ok {
				// panic("Not found ledger account id")
				fmt.Println("@@@@@ error: Not found ledger account id, ", groups[i].Currency, " order is ", groups[i].Order, groups[i].Entries[j])
				valid = false
				break
			}
			t.LedgerAccountId = ledgerAccount.Id

			journalReq.Transactions = append(journalReq.Transactions, t)
		}

		if valid {
			// doSend
			fmt.Println("Do send =>", journalReq)
			// doSendJournal(client, tokenResponse, journalReq)
			time.Sleep(time.Duration(500) * time.Millisecond)

			success++
		} else {
			failure++
		}

		reference += 2
	}

	fmt.Println("Finished, success", success, "failure", failure)
}

func initCreateJournalRequest(reference, currency, notes string) *CreateJournalRequest {
	return &CreateJournalRequest{
		Status:       "posted",
		EntryType:    "Standard",
		SourceLedger: "Financial",
		Reference:    reference,
		Currency:     currency,
		Notes:        notes,
		Transactions: []Transaction{},
	}
}

func doSendJournal(client *http.Client, token *TokenResponse, journalReq *CreateJournalRequest) {
	reqBytes, err := json.Marshal(journalReq)
	// fmt.Println(err)
	req, _ := http.NewRequest(http.MethodPost, "https://eu-api.softledger.com/api/journals", bytes.NewBuffer(reqBytes))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("$$ Error: Create journal failed.")
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Create journal succeed, Status", resp.Status, "response body", string(body))
}

func prepareAccessToken(client *http.Client, config *AccountConfig) *TokenResponse {
	if accessToken != "" {
		return &TokenResponse{AccessToken: accessToken}
	}

	requestTokenBody := &RequestTokenBody{
		GrantType:    config.GrantType,
		Audience:     config.Audience,
		TenantUUID:   config.TenantUUID,
		ClientId:     config.ClientId,
		ClientSecret: config.ClientSecret,
	}

	tokenBytes, _ := json.Marshal(requestTokenBody)
	req, _ := http.NewRequest(
		"POST",
		"https://softledger.eu.auth0.com/oauth/token",
		bytes.NewBuffer(tokenBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic("Error while getting access token.")
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	var tokenResponse = new(TokenResponse)
	json.Unmarshal(body, &tokenResponse)
	// fmt.Println("Access token is ", tokenResponse.AccessToken)
	return tokenResponse
}

func getAllLedgerAccounts(client *http.Client, token *TokenResponse) map[string]*LedgerAccount {

	req, _ := http.NewRequest(http.MethodGet, "https://eu-api.softledger.com/api/ledger_accounts", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	queries := req.URL.Query()
	queries.Add("limit", "500")
	req.URL.RawQuery = queries.Encode()

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Get all ledger accounts, response Status ", resp.Status, string(body))

	var ledgerAccountResponse = new(LedgerAccountResponse)
	json.Unmarshal(body, &ledgerAccountResponse)

	fmt.Println("Response of LedgerAccount, totalItems ", ledgerAccountResponse.TotalItems)

	if ledgerAccountResponse.TotalItems == 0 {
		panic("No LedgerAccount.")
	}

	accounts := make(map[string]*LedgerAccount)
	for i := range ledgerAccountResponse.Data {
		// fmt.Println(ledgerAccountResponse.Data[i])
		accounts[ledgerAccountResponse.Data[i].Number] = &ledgerAccountResponse.Data[i]
	}

	return accounts
}

func printGroups(groups []*Group) {
	for k := range groups {
		fmt.Println(k, groups[k].Len, groups[k].Currency, groups[k].Resident, groups[k].Order)
		for ki := range groups[k].Entries {
			fmt.Println("\t", groups[k].Entries[ki])
		}
	}

	fmt.Println("Got groups, size is", len(groups))
}
