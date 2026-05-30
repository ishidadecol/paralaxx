"use client"

import { ColumnDef } from "@tanstack/react-table"
import { Person } from "../types/person"

export const personColumns: ColumnDef<Person>[] = [
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
      const birthDateValue = row.getValue("birth_date");
      if (!birthDateValue) {
        return "-/-";
      }
      const date = new Date(birthDateValue as string);
      const formattedDate = date.toLocaleDateString("en-GB", {
        timeZone: "UTC",
        day: "2-digit",
        month: "2-digit",
        year: "numeric",
      });
      return formattedDate;
    },
  },
  {
    accessorKey: "gender",
    header: "Gender",
  },
]
