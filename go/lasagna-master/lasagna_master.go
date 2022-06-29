package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time <= 0 {
		return len(layers) * 2
	} else {
		return len(layers) * time
	}
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64
	for _, value := range layers {
		if value == "noodles" {
			noodles += 50
		} else if value == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaled []float64
	for i := 0; i < len(quantities); i++ {
		scaled = append(scaled, (quantities[i]/2)*float64(portions))
	}
	return scaled
}
