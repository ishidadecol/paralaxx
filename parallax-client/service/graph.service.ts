import { GraphResponse } from "@/types/graph";

export async function getGraph(): Promise<GraphResponse> {
  const response = await fetch(
    "http://localhost:8080/graph"
  );

  if (!response.ok) {
    throw new Error("Failed to fetch graph")
  }

  return response.json();
}
