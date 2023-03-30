package lasagna

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, val := range layers {
		if val == "noodles" {
			noodles += 50
		} else if val == "sauce" {
			sauce += 0.2
		}
	}
	return
}
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(list []float64, portions int) []float64 {
	result := make([]float64, len(list))
	for index, val := range list {
		result[index] = val * float64(portions) / 2
	}
	return result
}
