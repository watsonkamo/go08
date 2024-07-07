package piscine

func StrLen(s []int) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func Median(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	for i := 0; i < StrLen(arr)-1; i++ {
		for j := 0; j < StrLen(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr[2]
}
