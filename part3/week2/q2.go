package main

import (
	"coursera/stanford-algorithms/util"
	"fmt"
	"strconv"
	"strings"
)

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
