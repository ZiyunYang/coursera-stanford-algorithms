package main

import (
	"coursera/stanford-algorithms/util"
	"fmt"
)

//Q1：
//In this programming problem and the next you'll code up the knapsack algorithm from lecture.
//
//This file describes a knapsack instance, and it has the following format:
//
//[knapsack_size][number_of_items]
//
//[value_1] [weight_1]
//
//[value_2] [weight_2]
//
//...
//
//For example, the third line of the file is "50074 659", indicating that the second item has value 50074 and size 659, respectively.
//
//You can assume that all numbers are positive. You should assume that item weights and the knapsack capacity are integers.
//
//In the box below, type in the value of the optimal solution.
//
//Q2：
//This problem also asks you to solve a knapsack instance, but a much bigger one.
//
//This file describes a knapsack instance, and it has the following format:
//
//[knapsack_size][number_of_items]
//
//[value_1] [weight_1]
//
//[value_2] [weight_2]
//
//...
//
//For example, the third line of the file is "50074 834558", indicating that the second item has value 50074 and size 834558, respectively. As before, you should assume that item weights and the knapsack capacity are integers.
//
//This instance is so big that the straightforward iterative implemetation uses an infeasible amount of time and space. So you will have to be creative to compute an optimal solution. One idea is to go back to a recursive implementation, solving subproblems --- and, of course, caching the results to avoid redundant work --- only on an "as needed" basis. Also, be sure to think about appropriate data structures for storing and looking up solutions to subproblems.
//
//In the box below, type in the value of the optimal solution.



func main(){
	fmt.Println(knapsack("/Users/xkahj/Documents/code/go/src/coursera/stanford-algorithms/part3/week4/knapsack1.txt"))
	fmt.Println(knapsack("/Users/xkahj/Documents/code/go/src/coursera/stanford-algorithms/part3/week4/knapsack_big.txt"))
}

func knapsack(path string) int{
	value, weight, n, w := initPack(path)
	type s1 []int
	type s2 []s1
	var A = s2{
		make(s1, w+1),
		make(s1, w+1),
	}
	for j := 0; j < w+1; j++ {
		A[0][j] = 0

	}
	for i := 1; i < n+1; i++ {
		for j := 0; j < w+1; j++ {
			if weight[i] > j {
				A[1][j] = A[0][j]
			} else {
				if value[i]+A[0][j-weight[i]] > A[0][j] {
					A[1][j] = value[i] + A[0][j-weight[i]]
				} else {
					A[1][j] = A[0][j]
				}
			}
		}
		A[0] = A[1]
		A[1] = make(s1, w+1)

	}
	return A[0][w]
}
func initPack(path string) ([]int, []int, int, int) {
	lines, err := util.ReadTXT(path, "\n")
	if err != nil {
		fmt.Println("read file error---", err)
	}
	items, err := util.StringToInt(util.SplitAndTrim(lines[0], " "))
	if err != nil {
		fmt.Println("get num and weight err---", err)
	}
	w, n := items[0], items[1]
	var value = make([]int, n+1)
	var weight = make([]int, n+1)
	for i, line := range lines {
		if i == 0 {
			continue
		}
		items, err := util.StringToInt(util.SplitAndTrim(line, " "))
		if err != nil {
			fmt.Println("num conversion err---", err)
		}
		value[i] = items[0]
		weight[i] = items[1]
	}
	return value, weight, n, w
}
