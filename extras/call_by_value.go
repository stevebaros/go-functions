package extras

func NotModifying(number int) int {
	return  number*number
}

func Modifying(number *int) int {
	*number = *number * *number
	return *number
}