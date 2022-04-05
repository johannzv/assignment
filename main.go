package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	err, products := ReadProducts("products.json")

	if err != nil {
		panic("load products: could not load products")
	}

	err, inventory := ReadInventory("inventory.json")

	if err != nil {
		panic("load products: could no load inventory")
	}

	wh := NewWarehouse(products, inventory)

	router := gin.Default()

	router.GET("/product", func(context *gin.Context) {

		products := wh.GetProducts()
		context.IndentedJSON(http.StatusOK, products)

	})

	router.POST("/product/:id/buy", func(context *gin.Context) {

		productId := context.Param("id")

		if err != nil {
			context.JSON(http.StatusBadRequest, "")
			return
		}

		err = wh.SellProduct(productId)

		if err != nil {
			context.JSON(http.StatusBadRequest, "")
			return
		}

		context.JSON(http.StatusOK, "purchase successful")
		return
	})

	err = router.Run("localhost:8080")

	if err != nil {
		panic("failed to start server")
	}
}
