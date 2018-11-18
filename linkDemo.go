package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//定义链表节点
type Node struct {
	Value int
	Next *Node
}

// 创建链表
func CreateLink() *Node {

	input := bufio.NewScanner(os.Stdin)
	var head,temp *Node
	head = &Node{Value:0,Next:nil}
	temp = head

	for input.Scan() {
		strValue := input.Text()
		value, err := strconv.Atoi(strValue)
		if err != nil {
			fmt.Println("出现了非数字输入！程序结束")
			return nil
		}else {
			if value == 0 {
				return head
			} else {
				pnew := &Node{Value:value, Next:nil}
				temp.Next = pnew
				temp = pnew
			}
		}
	}
	return nil
}

// 打印链表
func print(head *Node) {

	var p *Node

	p = head.Next
	for p != nil {
		fmt.Print(p.Value," ")
		p = p.Next
	}


}

func main() {

	head := CreateLink()
	print(head)


}
