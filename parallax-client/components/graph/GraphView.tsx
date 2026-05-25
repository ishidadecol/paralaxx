"use client";

import React, { useEffect, useRef } from "react";
import * as d3 from "d3";
import { useGraph } from "@/hooks/useGraph";
import { GraphNode, GraphEdge } from "@/types/graph"; // Import existing types

// Define types for D3 nodes and links
interface D3GraphNode extends GraphNode, d3.SimulationNodeDatum {}
interface D3GraphEdge extends GraphEdge, d3.SimulationLinkDatum<D3GraphNode> {}

export default function GraphView() {
  const svgRef = useRef<SVGSVGElement | null>(null);
  const { graph, loading } = useGraph();

  useEffect(() => {
    if (!svgRef.current || !graph) return;

    const width = svgRef.current.clientWidth;
    const height = svgRef.current.clientHeight;

    d3.select(svgRef.current).selectAll("*").remove();

    const svg = d3
      .select(svgRef.current)
      .attr("width", width)
      .attr("height", height);

    const nodes: D3GraphNode[] = graph.nodes.map((d) => ({ ...d }));
    const links: D3GraphEdge[] = graph.edges.map((d) => ({
      ...d,
      source: d.source, // D3 expects source/target to be node objects or IDs
      target: d.target,
    }));

    const simulation = d3
      .forceSimulation<D3GraphNode, D3GraphEdge>(nodes)
      .force(
        "link",
        d3
          .forceLink<D3GraphNode, D3GraphEdge>(links)
          .id((d) => d.id)
          .distance(100)
      )
      .force("charge", d3.forceManyBody().strength(-300))
      .force("center", d3.forceCenter(width / 2, height / 2));

    const link = svg
      .append("g")
      .attr("stroke", "rgba(148, 163, 184, 0.35)")
      .attr("stroke-width", 1)
      .selectAll("line")
      .data(links)
      .join("line");

    const node = svg
      .append("g")
      .attr("stroke", "#fff")
      .attr("stroke-width", 1.5)
      .selectAll("circle")
      .data(nodes)
      .join("circle")
      .attr("r", 8)
      .attr("fill", "#f5f5f5")
      .call(drag(simulation)); // Apply drag behavior

    node.append("title").text((d) => d.name);

    const labels = svg
      .append("g")
      .attr("class", "labels")
      .selectAll("text")
      .data(nodes)
      .join("text")
      .attr("font-size", 10)
      .attr("fill", "#cbd5e1")
      .attr("text-anchor", "middle")
      .attr("dy", -10) // Position label above the node
      .text((d) => d.name);

    simulation.on("tick", () => {
      link
        .attr("x1", (d) => (d.source as D3GraphNode).x!)
        .attr("y1", (d) => (d.source as D3GraphNode).y!)
        .attr("x2", (d) => (d.target as D3GraphNode).x!)
        .attr("y2", (d) => (d.target as D3GraphNode).y!);

      node.attr("cx", (d) => d.x!).attr("cy", (d) => d.y!);

      labels.attr("x", (d) => d.x!).attr("y", (d) => d.y!);
    });

    const zoom = d3
      .zoom<SVGSVGElement, unknown>()
      .scaleExtent([0.1, 4])
      .on("zoom", (event) => {
        svg.attr("transform", event.transform);
      });

    d3.select(svgRef.current).call(zoom);

    function drag(simulation: d3.Simulation<D3GraphNode, D3GraphEdge>) {
      function dragstarted(event: d3.D3DragEvent<SVGCircleElement, D3GraphNode, D3GraphNode>) {
        if (!event.active) simulation.alphaTarget(0.3).restart();
        event.subject.fx = event.subject.x;
        event.subject.fy = event.subject.y;
      }

      function dragged(event: d3.D3DragEvent<SVGCircleElement, D3GraphNode, D3GraphNode>) {
        event.subject.fx = event.x;
        event.subject.fy = event.y;
      }

      function dragended(event: d3.D3DragEvent<SVGCircleElement, D3GraphNode, D3GraphNode>) {
        if (!event.active) simulation.alphaTarget(0);
        event.subject.fx = null;
        event.subject.fy = null;
      }

      return d3
        .drag<SVGCircleElement, D3GraphNode>()
        .on("start", dragstarted)
        .on("drag", dragged)
        .on("end", dragended);
    }
  }, [graph]);

  if (loading) {
    return (
      <div className="w-screen h-screen bg-[#0b0b0b] text-white flex items-center justify-center">
        Loading graph...
      </div>
    );
  }

  if (!graph) {
    return (
      <div className="w-screen h-screen bg-[#0b0b0b] text-white flex items-center justify-center">
        No graph data
      </div>
    );
  }

  return (
    <div className="w-screen h-screen bg-[#0b0b0b]">
      <svg ref={svgRef} className="w-full h-full"></svg>
    </div>
  );
}
