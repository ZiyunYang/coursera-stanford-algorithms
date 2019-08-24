package main

import (
	"coursera/stanford-algorithms/util"
	"fmt"
	"strconv"
)

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


