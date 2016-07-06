package main

import (
	"fmt"
	"github.com/draffensperger/golp"
)

func optimizePurchaseForRecipes(totals map[string]float64, priceList *PriceList) {
	priceMap := map[string]Ingredient{}
	for i := range *priceList {
		priceMap[(*priceList)[i].Name] = (*priceList)[i]
	}

	variableCount := 0
	ingredientList := []string{}
	mappedIngredientList := []string{}
	for ingredient, _ := range totals {
		ingredientList = append(ingredientList, ingredient)
		mappedIngredientList = append(mappedIngredientList, searchForIngredient(ingredient))
		variableCount += len(priceMap[ingredient].Packages)
	}

	lp := golp.NewLP(0, variableCount)

	missingPrice := false

	currentVariable := 0
	variablePrices := []float64{}
	for i, ingredient := range ingredientList {
		mappedIngredient := mappedIngredientList[i]
		total := totals[ingredient]
		if _, ok := priceMap[mappedIngredient]; !ok {
			fmt.Printf("Could not find price for: %v\n", ingredient)
			missingPrice = true
			continue
		}

		entries := []golp.Entry{}
		for _, pack := range priceMap[mappedIngredient].Packages {
			entries = append(entries, golp.Entry{currentVariable, pack.Amount})
			variablePrices = append(variablePrices, pack.Price)
			lp.SetInt(currentVariable, true)
			currentVariable++
		}
		lp.AddConstraintSparse(entries, golp.GE, total)
	}

	if missingPrice {
		panic("Unable to run optimization with incomplete price list")
	}


	lp.SetObjFn(variablePrices)
	lp.Solve()

	solvedVariables := lp.Variables()
	currentVariable = 0
	for i := range ingredientList {
		mappedIngredient := mappedIngredientList[i]
		for _, pack := range priceMap[mappedIngredient].Packages {
			if solvedVariables[currentVariable] > 0 {
				price := float64(solvedVariables[currentVariable]) * pack.Price
				fmt.Printf("%v%v of %v x %v (%.2f)\n", pack.Amount, pack.Unit, mappedIngredient, solvedVariables[currentVariable], price)
			}
			currentVariable++
		}
	}

	fmt.Printf("\nFinal Price: %.2f\n", lp.Objective())
}