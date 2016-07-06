package main

import (
	"github.com/stone/beerxml"
)

const (
	KG_TO_OZ = 35.2739619
	KG_TO_LB = 2.20462
)

func readBeerXMLFromFile(filename string) *beerxml.BeerXml {
	file, err := beerxml.NewBeerXmlFromFile(filename)
	if err != nil {
		panic(err)
	}
	return file
}

func addOrIncrement(key string, val float64, m *map[string]float64) {
	if _, ok := (*m)[key]; !ok {
		(*m)[key] = val
	} else {
		(*m)[key] = (*m)[key] + val
	}
}

func kgToOz(kg float64) float64 {
	return kg * KG_TO_OZ
}

func kgToLb(kg float64) float64 {
	return kg * KG_TO_LB
}

func getIngredientTotalsForRecipes(recipeList *RecipeList, beerXML *beerxml.BeerXml) map[string]float64 {
	totals := map[string]float64{}

	recipeLookup := map[string]int64{}
	for _, recipe := range *recipeList {
		recipeLookup[recipe.Name] = recipe.Quantity
	}

	for i := range beerXML.Recipes {
		if numRecipes, ok := recipeLookup[beerXML.Recipes[i].Name]; ok {
			for j := range beerXML.Recipes[i].Hops {
				amount := kgToOz(float64(beerXML.Recipes[i].Hops[j].Amount) * float64(numRecipes))
				addOrIncrement(beerXML.Recipes[i].Hops[j].Name, amount, &totals)
			}
			for j := range beerXML.Recipes[i].Fermentables {
				amount := kgToLb(float64(beerXML.Recipes[i].Fermentables[j].Amount) * float64(numRecipes))
				addOrIncrement(beerXML.Recipes[i].Fermentables[j].Name, amount, &totals)
			}
			for j := range beerXML.Recipes[i].Miscs {
				amount := float64(numRecipes) //Just assume for now that we need 1 item each time we see it
				addOrIncrement(beerXML.Recipes[i].Miscs[j].Name, amount, &totals)
			}
			for j := range beerXML.Recipes[i].Yeasts {
				amount := float64(numRecipes) //Just assume for now that we need 1 pack each time we see it
				addOrIncrement(beerXML.Recipes[i].Yeasts[j].Name, amount, &totals)
			}
		}
	}

	return totals
}