package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, val := range birdsPerDay {
		total += val
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	start := (week - 1) * 7
	return TotalBirdCount(birdsPerDay[start : start+7])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for index, _ := range birdsPerDay {
		//fmt.Printf("index: %d \n", index)
		if index%2 == 0 {
			birdsPerDay[index]++
		}
	}
	return birdsPerDay
}
