package main

import "fmt"

type node struct {
	prev *node
	data int
	next *node
}

var head *node = nil
var tail *node = nil

func create(val int) *node {
	var n node
	n.data = val
	n.prev = nil
	n.next = nil
	return &n
}
func insertf(val int) {
	n := create(val)
	if head == nil {
		head = n
		tail = n
	} else {
		n.next = head
		head.prev = n
		head = n
	}
}
func insertend(val int) {
	n := create(val)
	if tail == nil {
		tail = n
	} else {
		tail.next = n
		n.prev = tail
		tail = n
	}
}
func deletef() {
	if head == nil {
		return
	}
	if head.next == nil {
		head = nil
		println("nil")
	} else {
		head = head.next
		head.prev = nil
	}

}

func deleteend() {
	if tail == nil {
		return
	}
	if tail.prev == nil {
		tail = nil
		println("nil")
	} else {
		tail = tail.prev
		tail.next = nil
	}
}
func displayf() {
	for temp := head; temp != nil; temp = temp.next {
		fmt.Printf("%d ", temp.data)
	}
	fmt.Println()
}

func main() {
	insertf(10)
	insertf(20)
	insertend(30)
	insertend(40)
	deleteend()
	deletef()
	displayf()
}
