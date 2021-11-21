package web

type Edge interface {
	GetVertices() (vertices []string)
}

type EdgeImpl struct {
	vertices *[]string
}

func NewEdgeImpl(vertices []string) Edge {
	return &EdgeImpl{
		vertices: &vertices,
	}
}

func (edge *EdgeImpl) GetVertices() (vertices []string) {
	return *edge.vertices
}
