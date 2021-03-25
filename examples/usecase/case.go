package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

func bootstrap() {
	fmt.Println("use case")

	// 解析 excel 表格，获取到原始数据集

	// 遍历原始数据集，转换每一组消息为请求 body

	client := &http.Client{}

	// 获取访问接口的token
	prepareAccessToken(client)

	// 使用 token 调用 推送数据的接口
}

func prepareAccessToken(client *http.Client) *TokenResponse {
	requestTokenBody := &RequestTokenBody{
		GrantType:    "client_credentials",
		Audience:     "https://api.softledger.com",
		TenantUUID:   "3a06b187-627d-4fb0-aceb-d8c44c86453e",
		ClientId:     "1LJKbDlq2R0IY3hD36mzdQMWexG6cm6a",
		ClientSecret: "gkR9ip4SnMihkUwH8BzMWVMVaf6MP3GA6MH9no1KhuaRgZ5v7ti7SDDhVuD3t4Xa",
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
	fmt.Println("Access token is ", tokenResponse.AccessToken)
	return tokenResponse
}
