import React from "react";
import { Person } from "../types/person";

interface PersonTableProps {
  persons: Person[];
  loading: boolean;
  error: string | null;
}

const PersonTable: React.FC<PersonTableProps> = ({ persons, loading, error }) => {
  if (loading) {
    return <p>Loading persons...</p>;
  }

  if (error) {
    return <p className="text-red-500">Error: {error}</p>;
  }

  if (!persons || persons.length === 0) {
    return <p>No persons found. Add a new person to get started!</p>;
  }

  return (
    <div className="overflow-x-auto">
      <table className="min-w-full bg-white border border-gray-200">
        <thead>
          <tr>
            <th className="py-2 px-4 border-b">ID</th>
            <th className="py-2 px-4 border-b">First Name</th>
            <th className="py-2 px-4 border-b">Last Name</th>
            <th className="py-2 px-4 border-b">Gender</th>
            <th className="py-2 px-4 border-b">Birth Date</th>
            <th className="py-2 px-4 border-b">Created At</th>
          </tr>
        </thead>
        <tbody>
          {persons.map((person) => (
            <tr key={person.id} className="hover:bg-gray-50">
              <td className="py-2 px-4 border-b">{person.id}</td>
              <td className="py-2 px-4 border-b">{person.first_name}</td>
              <td className="py-2 px-4 border-b">{person.last_name}</td>
              <td className="py-2 px-4 border-b">{person.gender}</td>
              <td className="py-2 px-4 border-b">
                {person.birth_date ? new Date(person.birth_date).toLocaleDateString() : "N/A"}
              </td>
              <td className="py-2 px-4 border-b">
                {new Date(person.created_at).toLocaleString()}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default PersonTable;
