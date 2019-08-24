package main

import (
	"coursera/stanford-algorithms/util"
	"fmt"
	"strconv"
)

//In this programming problem and the next you'll code up the clustering algorithm from lecture for computing a max-spacing kk-clustering.
//
//This file describes a distance function (equivalently, a complete graph with edge costs). It has the following format:
//
//[number_of_nodes]
//
//[edge 1 node 1] [edge 1 node 2] [edge 1 cost]
//
//[edge 2 node 1] [edge 2 node 2] [edge 2 cost]
//
//...
//
//There is one edge (i,j)(i,j) for each choice of 1≤i<j≤n, where nn is the number of nodes.
//
//For example, the third line of the file is "1 3 5250", indicating that the distance between nodes 1 and 3 (equivalently, the cost of the edge (1,3)) is 5250. You can assume that distances are positive, but you should NOT assume that they are distinct.
//
//Your task in this problem is to run the clustering algorithm from lecture on this data set, where the target number kk of clusters is set to 4. What is the maximum spacing of a 4-clustering?
//
//ADVICE: If you're not getting the correct answer, try debugging your algorithm using some small test cases. And then post them to the discussion forum!

type Edge struct {
	weight  int
	vertex0 int
	vertex1 int
}

func initGraph() (int, []Edge) {
	lines, err := util.ReadTXT("/Users/xkahj/Documents/code/go/src/coursera/stanford-algorithms/part3/week2/clustering1.txt", "\n")
	if err != nil {
		fmt.Println("read txt error---", err)
	}
	vNum, err := strconv.Atoi(lines[0])
	if err != nil {
		fmt.Println("get vNum error---", err)
	}
	var edges []Edge
	for i, line := range lines {
		if i == 0 {
			continue
		}
		edge, err := util.StringToInt(util.SplitAndTrim(line, " "))
		if err != nil {
			fmt.Println("get edge error---", err)
		}
		edges = append(edges, Edge{
			weight:  edge[2],
			vertex0: edge[0] - 1,
			vertex1: edge[1] - 1,
		})

	}
	return vNum, edges
}

func edgeSort(arr []Edge) {
	var s = make([]Edge, len(arr)/2+1)
	if len(arr) < 2 {
		return
	}
	mid := len(arr) / 2

	edgeSort(arr[:mid])
	edgeSort(arr[mid:])

	if arr[mid-1].weight <= arr[mid].weight {
		return
	}
	copy(s, arr[:mid])
	l, r := 0, mid
	for i := 0; ; i++ {
		if s[l].weight <= arr[r].weight {
			arr[i] = s[l]
			l++

			if l == mid {
				break
			}
		} else {
			arr[i] = arr[r]
			r++
			if r == len(arr) {
				copy(arr[i+1:], s[l:mid])
				break
			}
		}
	}
	return
}


