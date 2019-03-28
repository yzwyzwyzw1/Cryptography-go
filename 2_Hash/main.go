package main

import (
	"awesomeProject/2_Hash/MyHashCode"
)

func main() {
	MyHashCode.CreateBuckets()
    MyHashCode.AddKeyValue("aaaas","hello?")

	MyHashCode.AddKeyValue("aaaas","world!")
    MyHashCode.GetValueByKey("aaaas")
}
