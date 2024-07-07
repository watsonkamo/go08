package piscine

func Compact(ptr *[]string) int {
	count := 0
	for _, str := range *ptr {
		if str != "" {
			(*ptr)[count] = str
			count++
		}
	}
	*ptr = (*ptr)[:count]
	return count
}
