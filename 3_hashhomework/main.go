package main

import (
	"fmt"
	"time"
)

var hNode *Node

var cNode *Node

type Node struct {
	Data     int
	NextNode *Node
}

func CreateHead(data int) *Node {
	var node = &Node{data, nil}
	hNode = node
	cNode = node
	return node
}

func AddNode(data int) *Node {
	var newNode = &Node{data, nil}
	cNode.NextNode = newNode

	cNode = newNode
	return newNode
}

func ShowNode() {
	node := hNode

	cnt := 1
	for {
		fmt.Println(node.Data)
		if cnt == 4 {

			fmt.Println("节点被剔除", node.NextNode.Data)
			node.NextNode = node.NextNode.NextNode
			cnt = 0
		}
		cnt++
		time.Sleep(1 * time.Second)
		if node.NextNode == nil {
			break
		} else {
			node = node.NextNode
		}

		if node.NextNode == node.NextNode.NextNode {
			fmt.Println(node.Data)
			fmt.Println("节点被剔除", node.Data)
			fmt.Println("结束程序")
			break
		}
	}

}

func main() {

	head := CreateHead(1)
	AddNode(2)
	AddNode(3)
	AddNode(4)
	AddNode(5)
	tail := AddNode(6)

	tail.NextNode = head

	ShowNode()
}
