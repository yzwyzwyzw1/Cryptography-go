package main

//只有函数名大写，才能算是共有函数类型，能够被外部函数调用

import (
	"awesomeProject/1_Nodes/linkNodes"
	"fmt"
)

func main(){
	fmt.Println("Hello World")

	//创建头节点
	head := linkNodes.CreateHeadNode(1)
	//fmt.Print(head.Data)

	node := linkNodes.AddNodeH(2,head)
	node = linkNodes.AddNodeH(3,node)
	linkNodes.ShowNodes(head)


	fmt.Printf("链表总长度：%d\n",linkNodes.LinkLen(head))

	fmt.Println("根据序号插入数据")

	head = linkNodes.InsertNodeWithIndex(100,0,head)
	fmt.Printf("链表总长度：%d",linkNodes.LinkLen(head))

	fmt.Println("修改指定下标节点的数据")
	linkNodes.UpdateNodeByIndex(100,2,head)
	linkNodes.ShowNodes(head)

	fmt.Println("删除指定下标节点的数据")
	head = linkNodes.DeleteNodeByIndex(0,head)
	linkNodes.ShowNodes(head)

	fmt.Println("头插法遍历链表")
	linkNodes.CreateHead(1)
	linkNodes.AddNodeT(2)
	linkNodes.AddNodeT(3)
	linkNodes.AddNodeT(4)
	linkNodes.ShowNodesT()

	}



