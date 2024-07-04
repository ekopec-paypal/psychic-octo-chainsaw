package scratch

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

// GetCustomerExchangesHistory the previously created exchanges for the given customer
func GetCustomerExchangesHistory(orders []*Order) (*ExchangeOrders, error) {

	exchangeOrderHistory := &ExchangeOrders{}
	customerOrders, err := OrdersByCustomer()

	findFirstOrderId(customerOrders)

	if err != nil {
		return nil, err
	}

	// We only need exchanges, so let's filter
	exchangeOrders := []*Order{}
	for _, order := range customerOrders {
		if HasTag(order.Tags, ExchangeTag) {
			exchangeOrders = append(exchangeOrders, order)
		}
	}

	// No exchanges! Just return empty
	if len(exchangeOrders) == 0 {
		return exchangeOrderHistory, nil
	}

	orderMap := make(map[int64]*Order)
	for _, order := range customerOrders {
		orderMap[order.OrderNumber] = order
	}

	result := findParentOrders(orderMap, exchangeOrders)

	exchangeHistory := make(map[int64]*ExchangeOrders)
	for _, order := range result {
		exchangeHistory[order.OrderNo] = order
	}

	//for k, order := range exchangeHistory {
	//	fmt.Printf("exchangeHistory: %+v\t\t%+v\t\t%+v\n\n", k, order.OrderNo, order.NumberExchange)
	//	for _, orderChild := range order.OrdersExchangeHist {
	//		fmt.Printf("exchangeHistory: %+v\n\n", orderChild)
	//	}
	//}

	return exchangeOrderHistory, nil
}

func findParentOrderNumber(note string) (int64, error) {
	pattern := `Created from[^()]+\((\d+)\)`
	re := regexp.MustCompile(pattern)

	match := re.FindStringSubmatch(note)
	if len(match) > 1 {
		var parentOrderNumber int64
		_, err := fmt.Sscanf(match[1], "%d", &parentOrderNumber)
		if err != nil {
			return 0, err
		}
		fmt.Printf("parentOrderNumber %+v\n\n", parentOrderNumber)
		return parentOrderNumber, nil
	}
	return 0, fmt.Errorf("no parent order number found")
}

func findParentOrders(orderMap map[int64]*Order, orderExchanges []*Order) []*ExchangeOrders {
	var result []*ExchangeOrders
	visited := make(map[int64]bool)

	for _, exchange := range orderExchanges {
		if visited[exchange.OrderNumber] {
			continue
		}
		var hist []*ExchangeOrdersHist
		currentOrder := exchange
		numberExchange := int64(0)
		for currentOrder != nil && currentOrder.Note != "" {
			parentOrderNumber, _ := findParentOrderNumber(currentOrder.Note)
			if parentOrderNumber != 0 {
				numberExchange++
			}
			fmt.Printf("OrderNumber: %+v, *parentOrderNumber %+v\n", currentOrder.OrderNumber, parentOrderNumber)
			hist = append(hist, &ExchangeOrdersHist{
				OrderID:       currentOrder.OrderNumber,
				ParentOrderID: parentOrderNumber,
			})
			visited[currentOrder.OrderNumber] = true
			currentOrder = orderMap[parentOrderNumber]

		}
		if currentOrder != nil {
			hist = append(hist, &ExchangeOrdersHist{
				OrderID:       currentOrder.OrderNumber,
				ParentOrderID: 0,
			})
			visited[currentOrder.OrderNumber] = true

		}

		result = append(result, &ExchangeOrders{
			OrderNo:            exchange.OrderNumber,
			OrdersExchangeHist: hist,
			NumberExchange:     numberExchange,
		})
	}
	for _, order := range result {
		fmt.Printf("Order: %+v NumberExchange: %+v\n", order.OrderNo, order.NumberExchange)
		for _, orderChild := range order.OrdersExchangeHist {
			fmt.Printf("\t\tOrderHist: %+v\n", orderChild)
		}
	}

	return result
}

func findFirstOrderId(orders []*Order) []*Order {
	var customerOrder []*Order
	seen := make(map[string]bool)
	for _, order := range orders {
		for _, lineItem := range order.LineItems {
			if len(lineItem.Properties) > 0 {
				var prop []map[string]interface{}
				if properties, ok := lineItem.Properties[`refund_transactions`].(string); ok {
					err := json.Unmarshal([]byte(properties), &prop)
					if err != nil {
						continue
					}
					for _, property := range prop {
						if value, exists := property["order_id"]; exists {
							orderID := strconv.FormatInt(int64(value.(float64)), 10)
							customerID := strconv.FormatInt(order.Customer.ID, 10)
							key := orderID + ":" + customerID
							// check if this combination already exists OrderID and CustomerID
							if seen[key] {
								continue
							}
							// Tag as seen
							seen[key] = true
							item := &Order{
								ID:       int64(value.(float64)),
								Customer: &Customer{ID: order.Customer.ID},
							}
							customerOrder = append(customerOrder, item)
						} else {
							fmt.Printf("Order ID not found in one of the extra fields for item with id\n")
						}

					}
				}

			}
		}
	}

	return customerOrder
}

func IsExchangedLimitReached(m map[string]interface{}, keys ...string) (interface{}, bool) {
	currentMap := m

	for _, key := range keys {
		if value, exists := currentMap[key]; exists {
			// If the value is a map, continue to the next level
			if nestedMap, ok := value.(map[string]interface{}); ok {
				currentMap = nestedMap
			} else {
				// If we can't go deeper, return the value
				return value, true
			}
		} else {
			// Key not found
			return nil, false
		}
	}
	return nil, false
}
