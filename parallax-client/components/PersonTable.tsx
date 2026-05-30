"use client"

import React from "react"
import { Person } from "../types/person"
import { DataTable } from "./DataTable" // Corrected import path
import { personColumns } from "./personColumns"

interface PersonTableProps {
  persons: Person[]
  loading: boolean
  error: string | null
}

const PersonTable: React.FC<PersonTableProps> = ({ persons, loading, error }) => {
  if (loading) {
    return <p>Loading persons...</p>
  }

  if (error) {
    return <p className="text-red-500">Error: {error}</p>
  }

  if (!persons || persons.length === 0) {
    return <p>No persons found. Add a new person to get started!</p>
  }

  return (
    <div className="overflow-x-auto">
      <DataTable columns={personColumns} data={persons} />
    </div>
  )
}

export default PersonTable
