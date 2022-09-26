package main

import (
	serv "SimpleStockExchange/ExecuteOrder"
	model "SimpleStockExchange/Models"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CreateRoutes() {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/test", testGET)
	router.POST("/order", orderPOST)

	router.Run(":8080")
}

func testGET(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Test succeeded")
}

func orderPOST(c *gin.Context) {
	var newOrder model.Order

	// Call BindJSON to bind the received JSON to newOrder
	if err := c.BindJSON(&newOrder); err != nil {
		return
	}

	chn := make(chan string)
	go serv.InitialProcess(newOrder, chn)
	status := <-chn

	c.IndentedJSON(http.StatusOK, status)
}
