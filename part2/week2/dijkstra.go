package main

import (
	"coursera/stanford-algorithms/graph"
	"coursera/stanford-algorithms/util"
	"fmt"
	pq "github.com/arnauddri/algorithms/data-structures/priority-queue"
	"strconv"
)

//In this programming problem you'll code up Dijkstra's shortest-path algorithm.
//Download the text file here. (Right click and save link as).
//The file contains an adjacency list representation of an undirected weighted graph with 200 vertices labeled 1 to 200. Each row consists of the node tuples that are adjacent to that particular vertex along with the length of that edge. For example, the 6th row has 6 as the first entry indicating that this row corresponds to the vertex labeled 6. The next entry of this row "141,8200" indicates that there is an edge between vertex 6 and vertex 141 that has length 8200. The rest of the pairs of this row indicate the other vertices adjacent to vertex 6 and the lengths of the corresponding edges.
//
//Your task is to run Dijkstra's shortest-path algorithm on this graph, using 1 (the first vertex) as the source vertex, and to compute the shortest-path distances between 1 and every other vertex of the graph. If there is no path between a vertex  v  and vertex 1, we'll define the shortest-path distance between 1 and  v  to be 1000000.
//
//You should report the shortest-path distances to the following ten vertices, in order: 7,37,59,82,99,115,133,165,188,197. You should encode the distances as a comma-separated string of integers. So if you find that all ten of these vertices except 115 are at distance 1000 away from vertex 1 and 115 is 2000 distance away, then your answer should be 1000,1000,1000,1000,1000,2000,1000,1000,1000,1000. Remember the order of reporting DOES MATTER, and the string should be in the same order in which the above ten vertices are given. Please type your answer in the space provided.
//
//IMPLEMENTATION NOTES: This graph is small enough that the straightforward  O(mn)  time implementation of Dijkstra's algorithm should work fine. OPTIONAL: For those of you seeking an additional challenge, try implementing the heap-based version. Note this requires a heap that supports deletions, and you'll probably need to maintain some kind of mapping between vertices and their positions in the heap

func makeGraph(path string) (*graph.UnGraph, error) {
	g := graph.NewUndirected()
	rows, err := util.ReadTXT(path, "\n")
	if err != nil {
		fmt.Println("1---", err)
		return g, err
	}
	for _, row := range rows {
		vertices := util.SplitAndTrim(row, "	")
		source, err := strconv.Atoi(vertices[0])
		if err != nil {
			fmt.Println("2---", err)
			return g, err
		}

		for i := 1; i < len(vertices); i++ {
			item := util.SplitAndTrim(vertices[i], ",")
			if len(item) != 2 {
				fmt.Println("3---", err)
				return g, err
			}
			vertex, err := strconv.Atoi(item[0])
			if err != nil {
				fmt.Println("4---", err)
				return g, err
			}
			dis, err := strconv.Atoi(item[1])
			if err != nil {
				fmt.Println("5---", err)
				return g, err
			}
			err = g.AddEdge(graph.VertexId(source), graph.VertexId(vertex), dis)
			if err != nil {
				//fmt.Println("source---", source)
				//fmt.Println("vertex---", vertex)
				//fmt.Println("dis---", dis)
				//fmt.Println("g---", g)
				fmt.Println("6---", err)
				return g, err
			}
		}
	}
	return g, nil
}

func ShortestPath(g *graph.UnGraph, source graph.VertexId) (map[graph.VertexId]graph.VertexId, map[graph.VertexId]int) {
	visited := make(map[graph.VertexId]bool, g.VerticesCount())
	dist := make(map[graph.VertexId]int)
	prev := make(map[graph.VertexId]graph.VertexId)
	Q := pq.NewMin()
	vertices := g.VerticesIter()

	dist[source] = 0
	for vertex := range vertices {
		if source != vertex {
			dist[vertex] = 1000000
			prev[vertex] = 0
		}
		Q.Insert(*pq.NewItem(vertex, dist[vertex]))
	}

	for Q.Len() > 0 {
		u := Q.Extract().Value.(graph.VertexId)
		visited[u] = true

		for neighbour := range g.GetNeighbours(u).VerticesIter() {
			if !visited[neighbour] {
				alt := dist[u] + g.GetEdge(u, neighbour)

				if alt < dist[neighbour] {
					dist[neighbour] = alt
					prev[neighbour] = u
					Q.ChangePriority(neighbour, alt)
				}
			}
		}
	}
	return prev, dist
}

func main() {
	g, err := makeGraph("/Users/xkahj/Documents/code/go/src/coursera/stanford-algorithms/part2/week2/dijkstraData.txt")
	if err != nil {
		fmt.Println("err1---", err)
	}
	_, dist := ShortestPath(g, graph.VertexId(1))
	fmt.Println("7,37,59,82,99,115,133,165,188,197")
	fmt.Println("7---", dist[graph.VertexId(7)])
	fmt.Println("37---", dist[graph.VertexId(37)])
	fmt.Println("59---", dist[graph.VertexId(59)])
	fmt.Println("82---", dist[graph.VertexId(82)])
	fmt.Println("99---", dist[graph.VertexId(99)])
	fmt.Println("115---", dist[graph.VertexId(115)])
	fmt.Println("133---", dist[graph.VertexId(133)])
	fmt.Println("165---", dist[graph.VertexId(165)])
	fmt.Println("188---", dist[graph.VertexId(188)])
	fmt.Println("197---", dist[graph.VertexId(197)])
}
