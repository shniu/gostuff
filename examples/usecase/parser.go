package usecase

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

// Parse excel, return records

var currencyMaps = map[string]string{
	"Fiat":      "USD",
	"BTC Asset": "BTC",
	"BCH Asset": "BCH",
	"LTC Asset": "LTC",
	"ETH Asset": "ETH",
}

func parseExcel() (groups []*Group) {
	path := "/Users/dfg/Desktop/softledger.xlsx"
	fmt.Println("File path is ", path)

	// Read excel
	f, err := excelize.OpenFile(path)
	if err != nil {
		panic(err)
	}

	sheetList := f.GetSheetList()
	// fmt.Println("Size: ", len(sheetList))
	for i := range sheetList {
		sheetName := strings.Trim(sheetList[i], " ")
		// fmt.Println("=>", i, "th sheet is ", sheetName)

		currency, ok := currencyMaps[sheetName]
		if !ok {
			panic("No Currency")
		}

		isStartLine := true
		order := 1
		resident := ""
		var entries []*Entry
		rows, _ := f.GetRows(sheetList[i])
		for i2 := range rows {
			// fmt.Println(i2, " rows is ", rows[i2], "and row size is ", len(rows[i2]))

			if isStartLine {
				// fmt.Println("=== Reaching an start line, then skip ===", rows[i2])
				// Next is not a start line
				isStartLine = false
				continue
			}

			if len(rows[i2]) > 0 && rows[i2][0] == "Account Name" {
				// fmt.Println("Skip", rows[i2])
				continue
			}

			// 表示结束位置，创建 Group
			if len(rows[i2]) == 0 || strings.Trim(rows[i2][0], " ") == "" {
				group := &Group{Entries: entries, Len: len(entries), Currency: currency, Resident: resident, Order: order}
				groups = append(groups, group)
				entries = nil

				// fmt.Println("=== Reaching an end line, then skip ===", rows[i2])
				// Next is a start line
				isStartLine = true
				order++
				resident = ""
				continue
			}

			entry := &Entry{
				AccountName:   rows[i2][0],
				AccountNumber: strings.Trim(rows[i2][1], " "),
				Debit:         strings.Replace(rows[i2][3], ",", "", -1),
				Credit:        strings.Replace(rows[i2][4], ",", "", -1),
			}
			entries = append(entries, entry)

			// Resident: resident
			if resident == "" {
				resident = NonUAE
			}
			if len(rows[i2]) >= 6 {
				if strings.Trim(rows[i2][5], " ") == UAE {
					resident = UAE
				}
			}
		}

		// 所有行遍历完后，把最后的记录加入进来
		if entries != nil {
			group := &Group{Entries: entries, Len: len(entries), Currency: currency, Resident: resident, Order: order}
			groups = append(groups, group)
			entries = nil
			resident = ""
		}

		order = 1

		// fmt.Println("")
	}

	return
}
