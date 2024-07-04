package scratch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func OrdersByCustomer() ([]*Order, error) {
	//call json with all orders
	jsonFile, err := os.Open("./psychic-octo-chainsaw/orders.json")
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	var ordersResponse struct {
		Orders []*Order `json:"orders"`
	}
	err = json.Unmarshal(byteValue, &ordersResponse.Orders)
	if err != nil {
		fmt.Printf("error exchangedSKUs decoding orders\n")
		return nil, err
	}
	fmt.Printf("total %+v\n\n\n\n", len(ordersResponse.Orders))
	return ordersResponse.Orders, nil
}
