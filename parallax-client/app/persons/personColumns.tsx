"use client"

import { ColumnDef } from "@tanstack/react-table"
import { format, parseISO } from "date-fns"

export type PersonTableColumn = {
  id: string,
  first_name: string,
  last_name: string,
  birth_date: string,
  gender: string
}

export const columns: ColumnDef<PersonTableColumn>[] = [
  {
    accessorKey: "first_name",
    header: "First Name",
  },
  {
    accessorKey: "last_name",
    header: "Last Name",
  },
  {
    accessorKey: "birth_date",
    header: "Birth Date",
    cell: ({ row }) => {
      const date = parseISO(row.getValue("birth_date"))
      return format(date, "PPP")
    },
  },
  {
    accessorKey: "gender",
    header: "Gender",
  }
]