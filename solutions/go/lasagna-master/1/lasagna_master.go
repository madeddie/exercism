// Package lasagnamaster does somethibg.
package lasagnamaster

// PreparationTime calculates PreparationTime.
func PreparationTime(layers []string, prepTime int) int {
	if prepTime == 0 {
		prepTime = 2
	}
	return prepTime * len(layers)
}

// Quantities calculates Quantities.
func Quantities(layers []string) (int, float64) {
	var noodlesGrams int
	var sauceLiters float64

	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodlesGrams += 50
		case "sauce":
			sauceLiters += 0.2
		}
	}
	return noodlesGrams, sauceLiters
}

// AddSecretIngredient adds a secret ingredient.
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// ScaleRecipe scales the recipe.
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaledQuantities []float64

	for _, quantity := range quantities {
		scaledQuantities = append(scaledQuantities, (quantity/2)*float64(portions))
	}
	return scaledQuantities
}
