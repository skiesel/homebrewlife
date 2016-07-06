package main

import (
	"encoding/json"
	"os"
)

type Package struct {
	Unit   string
	Amount float64
	Price  float64
}

type Ingredient struct {
	Name     string
	Packages []Package
}

type PriceList []Ingredient

func readPriceListFromFile(filename string) *PriceList {
	list := &PriceList{}
	file, err := os.Open(filename) 
	if err != nil {
		panic(err)
	}
	err = json.NewDecoder(file).Decode(list)
	if err != nil {
		panic(err)
	}
	return list
}