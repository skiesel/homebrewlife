package main

import (
	"encoding/json"
	"os"
)

type RecipeListItem struct {
	Name string
	Quantity int64
}

type RecipeList []RecipeListItem

func readRecipeListFromFile(filename string) *RecipeList {
	list := &RecipeList{}
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