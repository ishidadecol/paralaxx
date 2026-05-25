"use client";

import { useEffect, useState } from "react";

import { getGraph } from "@/service/graph.service";
import { GraphResponse } from "@/types/graph";

export function useGraph() {
  const [graph, setGraph] =
    useState<GraphResponse | null>(null);

  const [loading, setLoading] = useState(true);

  useEffect(() => {
    getGraph()
      .then(setGraph)
      .finally(() => setLoading(false));
  }, []);

  return {
    graph,
    loading,
  };
}
