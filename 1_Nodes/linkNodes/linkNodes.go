package linkNodes

import "fmt"

type Node struct {
	Data int        //数据域
    NextNode *Node  //地址域
}


//创建头节点指针
func CreateHeadNode(data int) *Node{
	var node =&Node{data,nil}
	return node  //放回对象地址
}

//头插法
func CreateHead(data int) *Node {
	var node = &Node{data,nil}
	hNode = node
	return node
}

func AddNodeT(data int)*Node {
	var newNode = &Node{data,nil}
	newNode.NextNode =  hNode
	hNode = newNode
	return newNode
}

//链表的遍历
func ShowNodesT() {
	node := hNode

	for {
		//打印节点的数据域信息
		fmt.Println("节点信息",node.Data)

		if node.NextNode == nil {
			break
		}else {
			node = node.NextNode //实现指针的位移
		}
	}
}

//通过尾插法添加新的节点

func AddNodeH(data int,node *Node) *Node{
	var newNode = &Node{data,nil}

	node.NextNode =newNode
	return newNode

}


//链表的遍历
func ShowNodes(head *Node) {
	node := head

	for {
		//打印节点的数据域信息
		fmt.Println(node.Data)

		if node.NextNode == nil {
			break
		}else {
			node = node.NextNode //实现指针的位移
		}
	}
}


//计算链表的总长度
func LinkLen(head *Node) int {
	var cnt =1
	node := head

	for {
		if node.NextNode == nil {
			break
		}else {
			node = node.NextNode
			cnt++
		}
	}
	return cnt
}


//通过全局变量记录当前链表的头节点
var hNode *Node

//按照下标插入新的节点
func InsertNodeWithIndex(data int,index int,node *Node) *Node {//头插法

	if index == 0 {
		var insertedNode = &Node{data,nil}

		insertedNode.NextNode = node
		hNode = insertedNode
		return insertedNode
	}else if index >= LinkLen(hNode)-1 {
		//当插入的节点的下标大于等于节点的长度-1时，就相当于在链表上插入了一个新节点
		node := hNode
		for {
			if node.NextNode == nil {
				//说明该节点是尾节点
				var insertedNode = &Node{data,nil}
				node.NextNode = insertedNode
				
				break
			}else {
				node = node.NextNode
			}
		}
	}else {
		head := node 
		var cnt = 1
		for {
			if cnt == index {
				var insertedNode = &Node{data,nil}
				
				insertedNode.NextNode = head.NextNode
				head.NextNode = insertedNode
			}else {
				head = head.NextNode
				cnt++
			}
		}
	}
	return nil
	
}

//修改指定下标节点的数据
func UpdateNodeByIndex(data int,index int,head *Node) {
	if index == 0 {
		hNode.Data = data
	}else {
		node := head
		var cnt = 1
		for {
			if node.NextNode == nil {
				break
			}else {
				if cnt == index {
					node.NextNode.Data = data
					break
				}else {
					node = node.NextNode
					cnt++
				}

			}
		
		}
	}
}

//删除指定下标的节点
func DeleteNodeByIndex(index int,head *Node) *Node {

	if index == 0 {
		node := head
		hNode = node.NextNode
		return hNode
	}else {
	     node := head
	     var cnt = 1

	     for {
	     	if node.NextNode == nil {
	     		break
			}else {
				if cnt == index {
					//删除这个节点
					node.NextNode = node.NextNode.NextNode
					break
				}else {
					node = node.NextNode
					cnt++
				}
	        }
		 }
	     return nil
    }
}