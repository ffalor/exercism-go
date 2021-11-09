package lasagna

// PreparationTime returns the estimated time it takes to prepare the lasagna
func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		avgPrepTime = 2
	}
	return len(layers) * avgPrepTime
}

// Quantities returns the number of noodles and sauce needed to make the lasagna
func Quantities(layers []string) (int, float64) {
	noodlesPerLayer := 50
	saucePerLayer := 0.2

	noodles := 0
	sauce := 0.0

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += noodlesPerLayer
		} else if layers[i] == "sauce" {
			sauce += saucePerLayer
		}
	}

	return noodles, sauce
}

// AddSecretIngredient adds friends secret ingredient to my lasagna
func AddSecretIngredient(friends []string, mine []string) []string {

	s := make([]string, 0, len(mine)+1)

	s = append(s, mine...)

	return append(s, friends[len(friends)-1:]...)
}

// ScaleRecipe scales the recipe by the given factor
func ScaleRecipe(ammountPerPortion []float64, desiredPortions int) []float64 {
	s := make([]float64, 0, len(ammountPerPortion))

	for i := 0; i < len(ammountPerPortion); i++ {
		scaleFactor := float64(desiredPortions) / 2
		s = append(s, ammountPerPortion[i]*scaleFactor)
	}

	return s
}
