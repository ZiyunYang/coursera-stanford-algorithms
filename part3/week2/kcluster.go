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

func main() {
	clusterNum, edges := initGraph()
	clusterID := make([]int, 500)
	size := make([]int, 500)
	var space_dist int
	for i := 0; i < clusterNum; i++ {
		clusterID[i] = i
		size[i] = 1
	}
	edgeSort(edges)
	for _, edge := range edges {
		if clusterNum == 4 {
			space_dist = edge.weight
		}
		clusterNum=union(edge.vertex0, edge.vertex1, clusterNum, size, clusterID)

	}
	fmt.Println("space_dist---", space_dist)
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

func find(p int, clusterID []int) int {
	for {
		if p == clusterID[p] {
			return p
		}
		clusterID[p] = clusterID[clusterID[p]]
		p = clusterID[p]
	}

}

func union(p, q, count int, sz, clusterID []int) int{
	i := find(p, clusterID)
	j := find(q, clusterID)
	if i == j {
		return count
	}
	if sz[i] < sz[j] {
		clusterID[i] = j
		sz[j] += sz[i]
	} else {
		clusterID[j] = i
		sz[i] += sz[j]
	}
	count--
	return count


}
