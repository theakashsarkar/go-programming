package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)

	for ele := intList.Front(); ele != nil; ele = ele.Next() {
		fmt.Println(ele.Value.(int))
	}

}
