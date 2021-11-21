package web

import (
	"sort"
)

type Graph interface {
	IsVertexExist(value string) bool
	AddVertex(value string) string
	GetVertex(value string) string
	AddEdge(v1, v2 string) Edge
	GetEdge(v1, v2 string) Edge
	GetEdges(value string) []Edge
	GetPaths(start, end string) [][]string
	GetAdjacencyVerticesList() map[string][]string
}

type GraphImpl struct {
	vertices    map[string]string
	edgesMatrix map[string]map[string]Edge
}

func NewGraph() Graph {
	return &GraphImpl{
		vertices:    make(map[string]string),
		edgesMatrix: make(map[string]map[string]Edge),
	}
}

func (graph *GraphImpl) IsVertexExist(value string) bool {
	_, ok := graph.vertices[value]
	return ok
}

func (graph *GraphImpl) AddVertex(value string) string {

	if graph.IsVertexExist(value) {
		return graph.GetVertex(value)
	}

	graph.vertices[value] = value

	for m := range graph.vertices {
		for n := range graph.vertices {
			if v, ok := graph.edgesMatrix[m]; !ok {
				graph.edgesMatrix[m] = map[string]Edge{n: nil}
			} else if _, ok := v[n]; !ok {
				graph.edgesMatrix[m][n] = nil
			}
		}
	}

	return value
}

func (graph *GraphImpl) GetVertex(value string) string {
	return graph.vertices[value]
}

func (graph *GraphImpl) AddEdge(v1, v2 string) Edge {

	if v1 == v2 {
		return nil
	}

	if !graph.IsVertexExist(v1) {
		graph.AddVertex(v1)
	}

	if !graph.IsVertexExist(v2) {
		graph.AddVertex(v2)
	}

	edge := NewEdgeImpl([]string{graph.GetVertex(v1), graph.GetVertex(v2)})

	graph.edgesMatrix[v1][v2] = edge
	graph.edgesMatrix[v2][v1] = edge

	return edge
}

func (graph *GraphImpl) GetEdge(v1, v2 string) Edge {

	if !graph.IsVertexExist(v1) || !graph.IsVertexExist(v2) {
		return nil
	}

	dict := graph.GetAdjacencyVerticesList()
	if _, ok := dict[v1]; !ok {
		return nil
	}

	if _, ok := dict[v2]; !ok {
		return nil
	}

	return graph.edgesMatrix[v1][v2]
}

func (graph *GraphImpl) GetEdges(value string) []Edge {

	if !graph.IsVertexExist(value) {
		return nil
	}

	edges := []Edge{}

	dict := graph.GetAdjacencyVerticesList()
	if adj, ok := dict[value]; !ok || len(adj) == 0 {
		return nil
	}

	for i := 0; i < len(dict[value]); i++ {
		edge := graph.edgesMatrix[value][dict[value][i]]
		edges = append(edges, edge)
	}

	return edges
}

func (graph *GraphImpl) GetPaths(v1, v2 string) [][]string {
	dfs := NewDFS(graph, v1, v2)
	return dfs.Search()
}

func (graph *GraphImpl) GetAdjacencyVerticesList() map[string][]string {

	dict := make(map[string][]string)

	for k := range graph.vertices {
		if !graph.IsVertexExist(k) {
			continue
		}
		dict[k] = []string{}
	}

	if len(dict) == 0 {
		return nil
	}

	for m, v := range graph.edgesMatrix {
		for n, edge := range v {
			if edge == nil {
				continue
			}
			if _, ok := dict[m]; ok {
				dict[m] = append(dict[m], n)
			}
		}
		sort.Slice(dict[m], func(i, j int) bool {
			return string(dict[m][i]) < string(dict[m][j])
		})
	}

	return dict
}

func (graph *GraphImpl) GetVertices() map[string]string {
	return graph.vertices
}

func (graph *GraphImpl) GetEdgesMatrix() map[string]map[string]Edge {
	return graph.edgesMatrix
}
