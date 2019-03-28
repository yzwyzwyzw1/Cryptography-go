package MyHashCode

import "awesomeProject/2_Hash/LinkNodes"

var buckets [16]*LinkNodes.Node

func CreateBuckets() {
	for i:=0;i<16;i++ {
		buckets[i] = LinkNodes.CreateHead()//为每个bucket头节点对象
	}
}

//自编的hash算法
func HashCode(Key string) int {
//此哈希散列算法可以将不同长度的Key 换算为【0-15】的整数
    var sum = 0
    for i:=0;i<len(Key);i++ {
    	sum+=int(Key[i])
	}
    return sum%16


}

//添加键值对
func AddKeyValue(Key string,Value string) {
	//通过Hash数列，将Key换算为0到15的整数

	var pos = HashCode(Key)
	var bucketsHeadNode = buckets[pos]

	bucketsHeadNode = LinkNodes.GetTailNode(bucketsHeadNode)


	var kv = LinkNodes.KV{Key,Value}
	LinkNodes.AddNode(kv,bucketsHeadNode)


}

//按键取值
func GetValueByKey(Key string) {
	var pos =HashCode(Key)
	var bucketsHeadNode = buckets[pos]

	//列表的遍历
	LinkNodes.ShowNode(Key,bucketsHeadNode)
}