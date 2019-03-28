package LinkNodes

import "fmt"

type KV struct {
	Key string
	Value string
}

type Node struct {
	Data KV
	NextNode *Node
}

func CreateHead() *Node {
	var node = &Node{KV{"头节点","头节点"},nil}
    return node
}


func AddNode(data KV,node *Node) *Node {
	var newNode = &Node{data,nil}
	node.NextNode = newNode
	return newNode

}

//遍历链表

func ShowNode(Key string,head *Node) {
	node := head

	for {
		//fmt.Println(node.Data)
		if node.Data.Key == Key {
			fmt.Println(node.Data.Value)
	    }

		if node.NextNode == nil {
			break
		}else {
			node = node.NextNode
		}
	}
}

func GetTailNode(Head *Node) *Node {
	node := Head
	for {
		if node.NextNode == nil {
			return node
		}else {
			node =node.NextNode
		}
	}
}

