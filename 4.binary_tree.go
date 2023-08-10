package main

import (
	"fmt"
	"time"
)

type nodet struct {
	data  int
	right *nodet
	left  *nodet
}

func (n *nodet) insert(root *nodet, val int) *nodet {
	if root == nil {
		new := nodet{val, nil, nil}
		fmt.Println("the element is inserted", val)
		return &new
	}
	
	return root
}
func (n *nodet) search(root *nodet, val int) *nodet {
	start := time.Now()
	for root != nil && root.data != val {
		if val < root.data {
			root = root.left
		} else {
			root = root.right
		}
	}
	timetaken := time.Since(start)
	fmt.Println("time taken", timetaken)
	return root
}
func main() {
	var root *nodet = nil
	root = root.insert(root, 10)
	root = root.insert(root, 30)
	root = root.insert(root, 5)
	root = root.insert(root, 31)
	root = root.insert(root, 40)
	r := root.search(root, 31)
	fmt.Println(r.left.data)
}
