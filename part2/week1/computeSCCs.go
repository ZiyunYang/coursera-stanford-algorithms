package main

import (
	"coursera/stanford-algorithms/graph"
	"coursera/stanford-algorithms/util"
	"errors"
	"fmt"
	"github.com/arnauddri/algorithms/data-structures/stack"
)


//The file contains the edges of a directed graph. Vertices are labeled as positive integers from 1 to 875714. Every row indicates an edge, the vertex label in first column is the tail and the vertex label in second column is the head (recall the graph is directed, and the edges are directed from the first column vertex to the second column vertex). So for example, the  11th  row looks liks : "2 47646". This just means that the vertex with label 2 has an outgoing edge to the vertex with label 47646
//
//Your task is to code up the algorithm from the video lectures for computing strongly connected components (SCCs), and to run this algorithm on the given graph.
//
//Output Format: You should output the sizes of the 5 largest SCCs in the given graph, in decreasing order of sizes, separated by commas (avoid any spaces). So if your algorithm computes the sizes of the five largest SCCs to be 500, 400, 300, 200 and 100, then your answer should be "500,400,300,200,100". If your algorithm finds less than 5 SCCs, then write 0 for the remaining terms. Thus, if your algorithm computes only 3 SCCs whose sizes are 400, 300, and 100, then your answer should be "400,300,100,0,0".
//
//WARNING: This is the most challenging programming assignment of the course. Because of the size of the graph you may have to manage memory carefully. The best way to do this depends on your programming language and environment, and we strongly suggest that you exchange tips for doing this on the discussion forums.


func MakeGraph(path, split string) (*graph.DirGraph, error) {
	g := graph.NewDirected()
	rows, err := util.ReadTXT(path, split)
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		items := util.SplitAndTrim(row, " ")
		if len(items) != 2 {
			return nil, errors.New("every edge should contains 2 vertices ")
		}
		vertices, err := util.StringToInt(items)
		if err != nil {
			fmt.Println("err1---", err)
			return nil, err
		}
		if err := g.AddEdge(graph.VertexId(vertices[0]), graph.VertexId(vertices[1]), 1); err != nil {
			fmt.Println("err3---", err, graph.VertexId(vertices[0]), graph.VertexId(vertices[1]))
			return nil, err
		}
	}
	return g, nil
}

func reversePostOrder(g *graph.DirGraph) *stack.Stack {
	s := stack.New()
	visited := make(map[graph.VertexId]bool, g.VerticesCount())
	vertices := g.VerticesIter()
	for s.Len() != g.VerticesCount() {
		vertex := <-vertices
		reverseDFS(g, vertex, s, visited)
	}
	return s
}

func reverseDFS(g *graph.DirGraph, from graph.VertexId, s *stack.Stack, visited map[graph.VertexId]bool) {
	if !visited[from] {
		visited[from] = true
		neighbours := g.GetNeighbours(from).VerticesIter()
		for neighbour := range neighbours {
			reverseDFS(g, neighbour, s, visited)
		}
		s.Push(from)
	}
}

func dfs(g *graph.DirGraph, from graph.VertexId, explored map[graph.VertexId]bool, sccNum *int) {
	if !explored[from] {
		*sccNum++
		explored[from] = true
		neighbours := g.GetSuccessors(from).VerticesIter()
		for neighbour := range neighbours {
			dfs(g, neighbour, explored, sccNum)
		}
	}
}

func computeSCC(g *graph.DirGraph) (int, []int) {
	reverseG := g.Reverse()
	explored := make(map[graph.VertexId]bool, g.VerticesCount())
	s := reversePostOrder(reverseG)
	sccNums := make([]int, 5)
	count := 0
	for s.Len() > 0 {
		v := s.Pop().(graph.VertexId)
		if !explored[v] {
			var sccNum int
			dfs(g, v, explored, &sccNum)
			if len(sccNums) < 5 {
				sccNums = append(sccNums, sccNum)
			} else {
				min, index := findMin(sccNums)
				if sccNum > min {
					sccNums[index] = sccNum
				}
			}
			count++
		}
	}
	return count, sccNums
}

func findMin(sccNum []int) (min int, index int) {
	min = sccNum[0]
	for i, num := range sccNum {
		if num < min {
			min = num
			index = i
		}
	}
	return
}

func main() {
	g, err := MakeGraph("/Users/xkahj/Documents/code/go/src/coursera/stanford-algorithms/part2/week1/SCC.txt", "\n")
	fmt.Println("err4---", err)
	count, sccNums := computeSCC(g)
	fmt.Println("count---", count)
	fmt.Println("sccNums---", sccNums)
}
