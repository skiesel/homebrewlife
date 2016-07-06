package main

import (
	"github.com/blevesearch/bleve"
	"github.com/stone/beerxml"
	"os"
	"fmt"
)

var (
	index bleve.Index
)

func buildIngredientIndexFromBeerXML(beerXML *beerxml.BeerXml) {
	os.RemoveAll("IngredientsIndex")

	var err error
	mapping := bleve.NewIndexMapping()
    index, err = bleve.New("IngredientsIndex", mapping)
    if err != nil {
        panic(err)
    }

    for i := range beerXML.Recipes {
		for j := range beerXML.Recipes[i].Hops {
			data := struct {
		        Name string
		    }{
		        Name: beerXML.Recipes[i].Hops[j].Name,
		    }
		    index.Index(data.Name, data)
		}
		for j := range beerXML.Recipes[i].Fermentables {
			data := struct {
		        Name string
		    }{
		        Name: beerXML.Recipes[i].Fermentables[j].Name,
		    }
		    index.Index(data.Name, data)
		}
		for j := range beerXML.Recipes[i].Miscs {
			data := struct {
		        Name string
		    }{
		        Name: beerXML.Recipes[i].Miscs[j].Name,
		    }
		    index.Index(data.Name, data)
		}
		for j := range beerXML.Recipes[i].Yeasts {
			data := struct {
		        Name string
		    }{
		        Name: beerXML.Recipes[i].Yeasts[j].Name,
		    }
		    index.Index(data.Name, data)
		}
	}
}

func buildIngredientIndexFromPriceList(priceList *PriceList) {
	os.RemoveAll("IngredientsIndex")

	var err error
	mapping := bleve.NewIndexMapping()
    index, err = bleve.New("IngredientsIndex", mapping)
    if err != nil {
        panic(err)
    }

    for i := range *priceList {
		data := struct {
	        Name string
	    }{
	        Name: (*priceList)[i].Name,
	    }
	    index.Index(data.Name, data)
	}
}

func searchForIngredient(ingredient string) string {
    query := bleve.NewMatchQuery(ingredient)
    search := bleve.NewSearchRequest(query)
    searchResults, err := index.Search(search)
    if err != nil {
        panic(err)
    }
    if len(searchResults.Hits) == 0 {
    	panic("No price found for: " + ingredient)
    }
    if searchResults.Hits[0].Score < 1. {
    	fmt.Printf("Poor match found : %s -> %s (%g)\n", ingredient, searchResults.Hits[0].ID, searchResults.Hits[0].Score)	
    }

    return searchResults.Hits[0].ID
}