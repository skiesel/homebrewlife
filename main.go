// file, err := beerxml.NewBeerXmlFromFile("recipes.xml")
// if err != nil {
// 	panic(err)
// }

// for _, recipe := range file.Recipes {
// 	fmt.Println(recipe.Name)
// }

// Read in beerxml file for recipes and ingredients
// What recipes are being made
// Sum ingredients needed

// Conversions := map[string]map[string]float64 {
// 	"lb" : map[string]float64 {
// 		"oz" : 16.,
// 	},
// 	"oz" : map[string]float64 {
// 		"lb" : 1./16.,
// 	},
// }

//https://cloud.google.com/appengine/docs/go/googlecloudstorageclient/read-write-to-cloud-storage#required_imports

// "github.com/stone/beerxml"

package main

import (
	"fmt"
	"github.com/draffensperger/golp"
)

func main() {

	type Package struct {
		Unit   string
		Amount float64
		Price  float64
	}

	type Ingredient struct {
		Name     string
		Packages []Package
	}

	PriceList := map[string]Ingredient{
		"2Row": Ingredient{
			Name: "2Row",
			Packages: []Package{
				Package{
					Unit:   "lb",
					Amount: 10,
					Price:  11.29,
				},
				Package{
					Unit:   "lb",
					Amount: 5,
					Price:  5.89,
				},
				Package{
					Unit:   "lb",
					Amount: 1,
					Price:  1.39,
				},
			},
		},
	}

	Totals := map[string]float64{
		"2Row": 17,
	}

	variableCount := 0
	for ingredient, _ := range Totals {
		variableCount += len(PriceList[ingredient].Packages)
	}

	lp := golp.NewLP(0, variableCount)

	currentVariable := 0
	variablePrices := []float64{}
	for ingredient, total := range Totals {
		entries := []golp.Entry{}
		for _, pack := range PriceList[ingredient].Packages {
			entries = append(entries, golp.Entry{currentVariable, pack.Amount})
			variablePrices = append(variablePrices, pack.Price)
			lp.SetInt(currentVariable, true)
			currentVariable++
		}
		lp.AddConstraintSparse(entries, golp.GE, total-0.0000001)
	}
	lp.SetObjFn(variablePrices)
	lp.Solve()

	fmt.Printf("Final Price: %v\n", lp.Objective())

	solvedVariables := lp.Variables()
	currentVariable = 0
	for ingredient, _ := range Totals {
		for _, pack := range PriceList[ingredient].Packages {
			if solvedVariables[currentVariable] > 0 {
				fmt.Printf("%v%v of %v x %v\n", pack.Amount, pack.Unit, ingredient, solvedVariables[currentVariable])
			}
			currentVariable++
		}
	}
}
