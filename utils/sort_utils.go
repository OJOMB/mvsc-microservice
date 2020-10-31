package utils

// Bubblesort sorting algo much easy extremely slow
func Bubblesort(list []int) []int {
	for running := true; running; {
		running = false
		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				running = true
			}
		}
	}
	return list
}
