package main

import "fmt"

type queue struct {
	size  int
	front int
	back  int
}

var que = make([]int, 10)

func (q *queue) add(ele int) int {
	if q.back == q.size-1 {
		fmt.Println("overflow")
		return 0
	}
	if q.front == -1 {
		q.front = 0
	}
	q.back++
	que[q.back] = ele

	return 0
}
func (q *queue) del() int {
	if q.front == -1 || q.front > q.back {
		fmt.Println("underflow")
		return 0
	}
	que[q.front] = que[q.back]
	que[q.back] = 0
	q.back--
	return 0
}

func dis() {
	fmt.Println(que)
}

func main() {
	a := queue{10, -1, -1}
	var ele int
	var input int
	for {

		fmt.Println("1.add 2.del 3. display 4.exit")
		fmt.Scanln(&input)

		if input == 1 {
			fmt.Println("enter element")
			fmt.Scanln(&ele)
			a.add(ele)
		} else if input == 2 {
			a.del()
		} else if input == 3 {
			dis()

		} else if input == 4 {
			break
		}

	}

}
