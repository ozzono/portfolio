package merge

func MergeCH(left, right, c chan int) {
	defer close(c)
	valLeft, okLeft := <-left
	valRight, okRight := <-right
	for okLeft && okRight {
		if valLeft < valRight {
			c <- valLeft
			valLeft, okLeft = <-left
		} else {
			c <- valRight
			valRight, okRight = <-right
		}
	}

	for okLeft {
		c <- valLeft
		valLeft, okLeft = <-left
	}

	for okRight {
		c <- valRight
		valRight, okRight = <-right
	}
}

func MergeSort(arr []int, ch chan int) {
	if len(arr) < 2 {
		ch <- arr[0]
		defer close(ch)
		return
	}

	left := make(chan int)
	right := make(chan int)
	go MergeSort(arr[len(arr)/2:], left)
	go MergeSort(arr[:len(arr)/2], right)
	go MergeCH(left, right, ch)
}
