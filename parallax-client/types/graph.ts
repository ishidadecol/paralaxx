export type GraphNode = {
  id: string;
  name: string;
  type: string;
};

export type GraphEdge = {
  id: string;
  source: string;
  target: string;
  type: string;
}

export type GraphResponse = {
  nodes: GraphNode[];
  edges: GraphEdge[];
}
