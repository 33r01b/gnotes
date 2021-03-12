package issue

func NumberPipeline(counter int) []int {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x <= counter; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	result := make([]int, 0)

	for x := range squares {
		result = append(result, x)
	}

	return result
}
