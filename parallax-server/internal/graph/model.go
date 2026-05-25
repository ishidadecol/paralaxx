package graph

type Node struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Edge struct {
	ID     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
	Type   string `json:"type"`
}

type Response struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}
