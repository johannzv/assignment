package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Products struct {
	Products []Product `json:"products"`
}

type Product struct {
	Name            string           `json:"name"`
	ContainArticles []ContainArticle `json:"contain_articles"`
}

type ContainArticle struct {
	ArtId    string `json:"art_id"`
	AmountOf int    `json:"amount_of,string"`
}

type Inventory struct {
	Inventory []Article `json:"inventory"`
}

type Article struct {
	ArtId string `json:"art_id"`
	Name  string `json:"name"`
	Stock int    `json:"stock,string"`
}

// ReadProducts - Read products from the specified path and unmarshalls the content.
func ReadProducts(path string) (error, Products) {

	jsonFile, err := os.Open(path)
	defer jsonFile.Close()

	if err != nil {
		return err, Products{}
	}

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return err, Products{}
	}

	var products Products

	err = json.Unmarshal(byteValue, &products)

	if err != nil {
		return err, products
	}

	return nil, products
}

// ReadInventory - Read inventory from the specified path and unmarshalls the content.
func ReadInventory(path string) (error, Inventory) {

	jsonFile, err := os.Open(path)
	defer jsonFile.Close()

	if err != nil {
		return err, Inventory{}
	}

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return err, Inventory{}
	}

	var inventory Inventory

	err = json.Unmarshal(byteValue, &inventory)

	if err != nil {
		return err, inventory
	}

	return nil, inventory
}
