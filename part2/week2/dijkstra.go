package main

import (
	"coursera/stanford-algorithms/graph"
	"coursera/stanford-algorithms/util"
	"fmt"
	pq "github.com/arnauddri/algorithms/data-structures/priority-queue"
	"strconv"
)

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
