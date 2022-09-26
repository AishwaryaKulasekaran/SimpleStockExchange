package executeOrder

import (
	model "SimpleStockExchange/Models"
	utils "SimpleStockExchange/Utilities"
	"strings"
)

var OrderArr = &model.Queue{Order: []model.Order{}, Size: 10}

func InitialProcess(newOrder model.Order, c chan string) {

	var orderStatus string
	if len(OrderArr.Order) == 0 {
		orderStatus = queueOrder(newOrder)
	} else {
		orderStatus = executeOrder(newOrder)
	}

	c <- orderStatus
}

func executeOrder(newOrder model.Order) string {

	var orderStatus string

	typeFlag := executeTypeCheck(newOrder.Type)
	priceFlag := executePriceCheck(newOrder)
	quantityFlag := executeQuantityCheck(newOrder)

	//Check for price and quanity check and perform action accordingly
	if typeFlag && priceFlag && quantityFlag {
		orderStatus = priceAndQuantityMatch()
	} else if typeFlag && priceFlag {
		orderStatus = onlyPriceMatch(newOrder)
	} else {
		orderStatus = queueOrder(newOrder)
	}

	return orderStatus
}

func priceAndQuantityMatch() string {
	var err, orderStatus string
	var latOrder []model.Order

	err, latOrder = popOrder(OrderArr.Order, len(OrderArr.Order))
	OrderArr.Order = latOrder
	if err != "" {
		orderStatus = err
	}
	orderStatus = "Order is successful"

	return orderStatus
}

func onlyPriceMatch(order model.Order) string {
	return queueOrder(order)
}

func popOrder(ord []model.Order, len int) (string, []model.Order) {
	var errStatus string
	ns, err := utils.RemoveElement(ord, len-1)
	if err != nil {
		errStatus = err.Error()
		return errStatus, nil
	}
	OrderArr.Order = ns
	return "", ns
}

func executeQuantityCheck(ord model.Order) bool {
	chk := (ord.Quantity == OrderArr.Order[len(OrderArr.Order)-1].Quantity)
	return chk
}

func executePriceCheck(ord model.Order) bool {

	chk := (ord.LimitPrice == OrderArr.Order[len(OrderArr.Order)-1].LimitPrice)
	return chk
}

func executeTypeCheck(orderType string) bool {

	chk := (OrderArr.Order[len(OrderArr.Order)-1].Type != strings.ToLower(orderType))
	return chk
}

func queueOrder(ord model.Order) string {
	var queueStatus string
	if len(OrderArr.Order) == OrderArr.Size {
		queueStatus = "Maximum queue size reached. Please try again later"
	} else {
		OrderArr.Order = append(OrderArr.Order, ord)
		queueStatus = "Order is queued"
	}
	return queueStatus
}
