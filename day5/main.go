package main

/*
create a map and then an array of numbers that you should not come across

Task:
Find which arrays fit
For all the arrays that are correct find the middle
sum all the middles
*/
func main() {
	var sum int
	sum += sum
}

func createMap() map[int][]int {
	m := make(map[int][]int)
	return m
}

func checkArray(m map[int][]int, arr []int) int {
	var middle int
	noFlyList := make(map[int]bool)
	for _, n := range arr {
		if noFlyList[n] {
			return 0
		}
		for _, i := range m[n] {
			noFlyList[i] = true
		}
	}
	middle = arr[(len(arr)/2)-1]
	return middle
}
