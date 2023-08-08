package main

import "fmt"

type node struct {
	prev *node
	data int
	next *node
}
var head *node=nil
func create(val int) *node {
	var n node
	n.data = val
	n.prev = nil
	n.next = nil
	return &n
}
func insertbg(val int){
	n := create(val)
	if(head==nil){
		head=n
	}else{
		n.next=head
		head.prev=n
		head =n
	}
}
func deletef(){
	if head==nil{
		return
	}
	if head .next == nil{
		head=nil
		println("nill")
	}else{
		head=head.next
		head.prev=nil
	}
	
}
func displayf(){
	for temp := head;temp != nil;temp = temp.next{
	   fmt.Printf("%d ", temp.data)
	}
	fmt.Println()
 }


func main() {
	insertbg(10)
	insertbg(20)
	deletef()
	deletef()
	displayf()
}
