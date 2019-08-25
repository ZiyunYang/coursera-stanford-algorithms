package main

import (
	pq "coursera/stanford-algorithms/priority-queue"
	"coursera/stanford-algorithms/util"
	"fmt"
	"strconv"
)

//In this programming problem and the next you'll code up the greedy algorithm from the lectures on Huffman coding.
//
//This file describes an instance of the problem. It has the following format:
//
//[number_of_symbols]
//
//[weight of symbol #1]
//
//[weight of symbol #2]
//
//...
//
//For example, the third line of the file is "6852892," indicating that the weight of the second symbol of the alphabet is 6852892. (We're using weights instead of frequencies, like in the "A More Complex Example" video.)
//
//Your task in this problem is to run the Huffman coding algorithm from lecture on this data set. What is the maximum length of a codeword in the resulting Huffman code?
//
//ADVICE: If you're not getting the correct answer, try debugging your algorithm using some small test cases. And then post them to the discussion forum!
//
//Continuing the previous problem, what is the minimum length of a codeword in your Huffman code?

type Node struct {
	//weight int
	left   *Node
	right  *Node
}

func main() {
	_, q := initNode()
	var tree *Node
	for {
		l := q.Extract()
		r := q.Extract()
		parent := &Node{
			left:   l.Value.(*Node),
			right:  r.Value.(*Node),
			//weight: l.Priority + r.Priority,
		}

		// log for debug, weight field is needed during debug
		//fmt.Println(parent.weight)
		//if l.Value.(*Node).left!=nil&&l.Value.(*Node).right!=nil{
		//	fmt.Println(l.Value.(*Node).weight,l.Value.(*Node).left.weight,l.Value.(*Node).right.weight)
		//}else{
		//	fmt.Println(l.Value.(*Node).weight)
		//}
		//if r.Value.(*Node).left!=nil&&r.Value.(*Node).right!=nil{
		//	fmt.Println(r.Value.(*Node).weight,r.Value.(*Node).left.weight,r.Value.(*Node).right.weight)
		//}else{
		//	fmt.Println(r.Value.(*Node).weight)
		//}
		//fmt.Println("-------------------------------------------------------")
		q.Insert(pq.Item{
			Priority: l.Priority + r.Priority,
			Value:    parent,
		})
		if q.Len() == 1 {
			tree=q.Data()[0].(pq.Item).Value.(*Node)
			break
		}

	}

	fmt.Println("min--",min(tree)-1)
	fmt.Println("max--",max(tree)-1)

}

func initNode() (int, *pq.PQ) {
	q := pq.NewMin()
	lines, err := util.ReadTXT("/Users/xkahj/Documents/code/go/src/coursera/stanford-algorithms/part3/week3/huffman.txt", "\n")
	if err != nil {
		fmt.Println("read file error")
	}
	n, err := strconv.Atoi(lines[0])
	if err != nil {
		fmt.Println("get sum err---", err)
	}
	for i, line := range lines {
		if i == 0 {
			continue
		}
		v, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("transfer to weight err---", err)
		}
		q.Insert(pq.Item{
			Priority: v,
			Value:    &Node{},
		})
	}
	return n, q
}

func min(root *Node)int{
	if root==nil{
		return 0
	}
	if root.left == nil && root.right == nil {
		return 1
	}
	minleft := 0
	minright := 0
	mindepth := 0
	if root.left!=nil{
		minleft=min(root.left)+1
	}
	if root.right!=nil{
		minright=min(root.right)+1
	}
	if root.left==nil{
		mindepth=minright
	}else if root.right==nil{
		mindepth=minleft
	}else{
		if minleft<minright{
			mindepth=minleft
		}else{
			mindepth=minright
		}
	}
	return mindepth

}

func max(root *Node)int{
	maxdepth:=0
	maxl:=0
	maxr:=0
	if root==nil{
		return 0
	}
	maxl=max(root.left)
	maxr=max(root.right)
	if maxl>maxr{
		maxdepth=maxl+1
	}else{
		maxdepth=maxr+1
	}
	return maxdepth
}