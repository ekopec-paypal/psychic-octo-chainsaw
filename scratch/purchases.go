package scratch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func SearchPurchases() ([]interface{}, error) {
	// detect if query is an email
	var (
		err error
	)
	filename := "./psychic-octo-chainsaw/order.json"

	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("err json file %+v\n\n", err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Printf("err close json file %+v\n\n", err)
		}
	}(jsonFile)

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error Purchases reading file:", err)
	}

	var ordersResponse struct {
		Orders []*Order `json:"orders"`
	}
	err = json.Unmarshal(byteValue, &ordersResponse.Orders)
	if err != nil {
		fmt.Printf("error Purchases decoding orders %+v\n", err)
		return nil, err
	}
	fmt.Printf("ordersResponse details %+v\n\n\n", ordersResponse.Orders)

	_, err = GetCustomerExchangesHistory(ordersResponse.Orders)
	if err != nil {
		// tratar erro
		fmt.Printf("exchangedSKUs err =>> %+v\n", err)
	}

	filename2 := "./psychic-octo-chainsaw/details.json"

	jsonFile2, err := os.Open(filename2)
	if err != nil {
		fmt.Printf("err json file %+v\n\n", err)
	}
	defer func(jsonFile2 *os.File) {
		err := jsonFile2.Close()
		if err != nil {
			fmt.Printf("err close json file %+v\n\n", err)
		}
	}(jsonFile2)

	byteValue2, err := ioutil.ReadAll(jsonFile2)
	if err != nil {
		fmt.Println("Error Purchases reading file:", err)
	}

	var myDetails map[string]interface{}
	fmt.Printf("Details : %+v\n", len(myDetails))
	err = json.Unmarshal(byteValue2, &myDetails)
	if err != nil {
		fmt.Printf("error Purchases decoding orders %+v\n", err)
		return nil, err
	}

	if value, ok := IsExchangedLimitReached(myDetails, "exchange_orders", "exchange_limit_reached"); ok {
		fmt.Printf("found %+v \n", value)

	} else {
		fmt.Printf("Not found %+v\n", value)
	}

	return nil, err
}
