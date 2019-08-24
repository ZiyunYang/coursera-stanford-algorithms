package main

import "fmt"

func main() {

	// q1
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

	// q2
	clusterNum, vertexMap, hashTable := initHash()
	//init cluster
	clusterID2 := make([]int, clusterNum)
	size2 := make([]int, clusterNum)
	for i := 0; i < clusterNum; i++ {
		clusterID2[i] = i
		size2[i] = 1
	}
	count := clusterNum
	//hamming distance=0
	for _, vertexs := range hashTable {
		if len(vertexs) > 1 {
			for i := 0; i < len(vertexs)-1; i++ {
				for j := 1; j < len(vertexs); j++ {
					count = union(vertexs[i], vertexs[j], count, size2, clusterID2)
				}
			}
		}
	}

	for i := 0; i < clusterNum; i++ {
		key := vertexMap[i]
		// hamming distance=1
		for j := 0; j < len(key); j++ {
			newKey := invertOneBit(key, j)
			//fmt.Println(j, "       ", newKey)
			if hashTable[newKey] != nil {
				for _, v := range hashTable[newKey] {
					count = union(i, v, count, size2, clusterID2)
				}
			}
		}
		// hamming distance=2
		for m := 0; m < len(key)-1; m++ {
			for n := m + 1; n < len(key); n++ {
				newKey := invertTwoBit(key, m, n)
				if hashTable[newKey] != nil {
					for _, v := range hashTable[newKey] {
						count = union(i, v, count, size2, clusterID2)
					}
				}
			}
		}
	}
	fmt.Println("count---",count)
}