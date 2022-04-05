package main

import (
	"fmt"
	"strconv"
)

type warehouseProduct struct {
	Name      string           `json:"name"`
	Articles  []ContainArticle `json:"articles"`
	Quantity  int              `json:"quantity"`
	ProductId string           `json:"productId"`
}

type Warehouse struct {
	inventory []Article
	product   []warehouseProduct
}

func NewWarehouse(products Products, inventory Inventory) *Warehouse {
	var wp []warehouseProduct
	for i, pr := range products.Products {
		wp = append(wp, warehouseProduct{
			Name:     pr.Name,
			Articles: pr.ContainArticles,
			// create an productId for each product
			ProductId: strconv.Itoa(i),
		})
	}
	return &Warehouse{
		inventory: inventory.Inventory,
		product:   wp,
	}
}

// GetProducts - returns all products in the warehouse. Uses the inventory to calculate the quantity/stock for each product.
func (wh *Warehouse) GetProducts() []warehouseProduct {
	var warehouseProductsWithStockCount = []warehouseProduct{}

	for _, product := range wh.product {
		// find out how many products the warehouse contain based on the number of articles in the warehouse
		var numberOfProductsBasedOnArticles = make([]int, len(product.Articles))

		for i, article := range product.Articles {
			numberOfProductsBasedOnArticles[i] = wh.GetStockForArticleByArticleId(article.ArtId) / article.AmountOf
		}

		// the smallest number, is the stock count for this product
		stockCount := getMin(numberOfProductsBasedOnArticles)

		warehouseProductsWithStockCount = append(warehouseProductsWithStockCount, warehouseProduct{
			Name:      product.Name,
			Articles:  product.Articles,
			Quantity:  stockCount,
			ProductId: product.ProductId,
		})

	}
	return warehouseProductsWithStockCount
}

// GetStockForArticleByArticleId - iterates through the articles in the warehouse and returns the stock count for it. Returns 0 if the the article is not in stock or not in the warehouse.
func (wh *Warehouse) GetStockForArticleByArticleId(artId string) int {
	for _, article := range wh.inventory {
		if article.ArtId == artId {
			return article.Stock
		}
	}
	return 0
}

// GetProductByProductId - iterates through the products in the warehouse and returns the product if it can be found. Error if it does not exist
func (wh *Warehouse) GetProductByProductId(productId string) (warehouseProduct, error) {
	for _, product := range wh.product {
		if product.ProductId == productId {
			return product, nil
		}
	}
	return warehouseProduct{}, fmt.Errorf("cannot find product")
}

// SellProduct - will try to sell a product and return an error if the product is not in the warehouse or if some articles are not in stock
func (wh *Warehouse) SellProduct(productId string) error {

	// create a temp version of the warehouse, so it can be rolled back
	var tempStock = wh.inventory

	// check if the product is in the warehouse
	product, err := wh.GetProductByProductId(productId)

	if err != nil {
		return err
	}

	// go through is article that this product contains
	for _, article := range product.Articles {

		// check if the stock count is less than what this project requires
		if wh.GetStockForArticleByArticleId(article.ArtId) < article.AmountOf {
			return fmt.Errorf("warehous error: stock to low")
		}

		// the article is in stock, and the temporary updated inventory can be changed
		for i, artTemp := range tempStock {
			if article.ArtId == artTemp.ArtId {
				tempStock[i].Stock = artTemp.Stock - article.AmountOf
			}
		}
	}

	// every article is in stock, the inventory can be updated
	wh.inventory = tempStock

	return nil
}

// getMin - returns the value for the lowest entry in the slice
func getMin(array []int) int {
	min := array[0]
	for _, val := range array {
		if val < min {
			min = val
		}
	}
	return min
}
