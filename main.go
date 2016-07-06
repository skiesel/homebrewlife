package main

import (
	"flag"
)

var (
	beerXMLFile = flag.String("beerxml", "recipes.xml", "The Beer XML File containing recipes")
	recipeListFile = flag.String("recipelist", "recipelist.json", "The list of recipe names to purchase for")
	priceListFile = flag.String("pricedb", "pricelist.json", "The price list")
)

func main() {
	flag.Parse()

	// read in the beerxml file containing recipes
	beerXML := readBeerXMLFromFile(*beerXMLFile)
	if beerXML == nil {
		panic("Unable to read beer xml file")
	}

	// read in the list of recipes to purchase ingredients for
	recipeList := readRecipeListFromFile(*recipeListFile)
	if recipeList == nil {
		panic("Unable to read recipe list file")
	}

	// read in the list of prices for ingredients
	priceList := readPriceListFromFile(*priceListFile)
	if priceList == nil {
		panic("Unable to read price database file")
	}

	buildIngredientIndexFromPriceList(priceList)

	totals := getIngredientTotalsForRecipes(recipeList, beerXML)
	
	// for ingredient, total := range totals {
	// 	fmt.Printf("%v %v\n", total, ingredient)
	// }
	// fmt.Println()

	optimizePurchaseForRecipes(totals, priceList)
}
