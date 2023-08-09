package main

import "fmt"

type node struct {
	prev *node
	data int
	next *node
}

var head *node = nil
var tail *node = nil

func (n *node) create(val int) *node {

	n.data = val
	n.prev = nil
	n.next = nil
	return n
}
func (n *node) insertf(val int) {
	na := n.create(val)
	if head == nil {
		head = na
		tail = na
	} else {
		na.next = head
		head.prev = na
		head = na
	}
}
func (i *node) insertend(val int) {
	n := i.create(val)
	if tail == nil {
		tail = n
	} else {
		tail.next = n
		n.prev = tail
		tail = n
	}
}
func (i *node) deletef() {
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

func (i *node) deleteend() {
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
func (i *node) displayf() {
	for temp := head; temp != nil; temp = temp.next {
		fmt.Printf("%d ", temp.data)
	}
	fmt.Println()
}

func main() {
	a := node{}
	a.insertf(10)
	a.insertf(20)
	a.insertend(30)
	a.insertend(40)
	a.deleteend()
	a.deletef()
	a.displayf()
}
