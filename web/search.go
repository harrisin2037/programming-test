package web

type DFS struct {
	start, end string
	marked     map[string]bool
	dict       map[string][]string
	paths      []string
}

func NewDFS(graph Graph, start, end string) *DFS {
	// fmt.Println("graph.GetAdjacencyVerticesList()", graph.GetAdjacencyVerticesList())
	return &DFS{
		start, end,
		make(map[string]bool),
		graph.GetAdjacencyVerticesList(),
		[]string{},
	}
}

func (dfs *DFS) Search() [][]string {

	var (
		current = []string{dfs.start}
		result  = [][]string{}
	)

	return *dfs.search(current, &result)
}

func (dfs *DFS) SearchShortest() []string {

	var (
		current = []string{dfs.start}
		result  = []string{}
	)

	return *dfs.searchShortest(current, &result)
}

func (dfs *DFS) search(current []string, result *[][]string) *[][]string {

	if current[len(current)-1] == dfs.end {
		clone := make([]string, len(current))
		copy(clone, current)
		*result = append(*result, clone)
		return result
	}

	last := current[len(current)-1]

	for i := 0; i < len(dfs.dict[last]); i++ {

		if exist, ok := dfs.marked[dfs.dict[last][i]]; ok && exist {
			continue
		}

		current = append(current, dfs.dict[last][i])

		dfs.marked[dfs.dict[last][i]] = true
		dfs.search(current, result)
		dfs.marked[dfs.dict[last][i]] = false

		current = current[:len(current)-1]
	}

	return result
}

func (dfs *DFS) searchShortest(current []string, result *[]string) *[]string {

	short := []string{}

	if current[len(current)-1] == dfs.end {
		clone := make([]string, len(current))
		copy(clone, current)
		*result = clone
		return result
	}

	last := current[len(current)-1]

	for i := 0; i < len(dfs.dict[last]); i++ {

		if exist, ok := dfs.marked[dfs.dict[last][i]]; ok && exist {
			continue
		}

		current = append(current, dfs.dict[last][i])

		dfs.marked[dfs.dict[last][i]] = true
		new := dfs.searchShortest(current, result)
		if len(short) == 0 || len(*new) < len(short) {
			short = *new
		}
		dfs.marked[dfs.dict[last][i]] = false

		current = current[:len(current)-1]
	}

	return &short
}
