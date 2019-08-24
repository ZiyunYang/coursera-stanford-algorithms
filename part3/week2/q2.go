package main

import (
	"coursera/stanford-algorithms/util"
	"fmt"
	"strconv"
	"strings"
)

//In this question your task is again to run the clustering algorithm from lecture, but on a MUCH bigger graph. So big, in fact, that the distances (i.e., edge costs) are only defined implicitly, rather than being provided as an explicit list.
//
//The format is:
//
//[# of nodes] [# of bits for each node's label]
//
//[first bit of node 1] ... [last bit of node 1]
//
//[first bit of node 2] ... [last bit of node 2]
//
//...
//
//For example, the third line of the file "0 1 1 0 0 1 1 0 0 1 0 1 1 1 1 1 1 0 1 0 1 1 0 1" denotes the 24 bits associated with node #2.
//
//The distance between two nodes uu and vv in this problem is defined as the Hamming distance--- the number of differing bits --- between the two nodes' labels. For example, the Hamming distance between the 24-bit label of node #2 above and the label "0 1 0 0 0 1 0 0 0 1 0 1 1 1 1 1 1 0 1 0 0 1 0 1" is 3 (since they differ in the 3rd, 7th, and 21st bits).
//
//The question is: what is the largest value of kk such that there is a kk-clustering with spacing at least 3? That is, how many clusters are needed to ensure that no pair of nodes with all but 2 bits in common get split into different clusters?
//
//NOTE: The graph implicitly defined by the data file is so big that you probably can't write it out explicitly, let alone sort the edges by cost. So you will have to be a little creative to complete this part of the question. For example, is there some way you can identify the smallest distances without explicitly looking at every pair of nodes?

func invert(bit string) string {
	if bit == "0" {
		return "1"
	} else {
		return "0"
	}
}

func invertOneBit(key string, num int) string {
	return key[:num] + invert(key[num:num+1]) + key[num+1:]

}

func invertTwoBit(key string, m, n int) string {
	return key[:m] + invert(key[m:m+1]) + key[m+1:n] + invert(key[n:n+1]) + key[n+1:]
}

func initHash() (int, map[int]string, map[string][]int) {
	lines, err := util.ReadTXT("/Users/xkahj/Documents/code/go/src/coursera/stanford-algorithms/part3/week2/clustering_big.txt", "\n")
	if err != nil {
		fmt.Println("read txt error---", err)
	}
	vNum, err := strconv.Atoi(util.SplitAndTrim(lines[0], " ")[0])
	if err != nil {
		fmt.Println("get vNum error---", err)
	}
	hashTable := make(map[string][]int)
	vertexMap := make(map[int]string)
	for i, line := range lines {
		if i == 0 {
			continue
		}
		key := strings.Replace(line, " ", "", -1)
		//key := line
		if hashTable[key] == nil {
			hashTable[key] = []int{i - 1}
		} else {
			old := hashTable[key]
			hashTable[key] = append(old, i-1)
		}
		vertexMap[i-1] = key
	}
	return vNum, vertexMap, hashTable
}

//func find(p int, clusterID []int) int {
//	for {
//		if p == clusterID[p] {
//			return p
//		}
//		clusterID[p] = clusterID[clusterID[p]]
//		p = clusterID[p]
//	}
//
//}
//
//func union(p, q, count int, sz, clusterID []int) int {
//	i := find(p, clusterID)
//	j := find(q, clusterID)
//	if i == j {
//		return count
//	}
//	if sz[i] < sz[j] {
//		clusterID[i] = j
//		sz[j] += sz[i]
//	} else {
//		clusterID[j] = i
//		sz[i] += sz[j]
//	}
//	count--
//	return count
//
//}
