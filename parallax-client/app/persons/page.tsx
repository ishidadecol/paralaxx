"use client";

import React from "react";
import { usePersons } from "../../hooks/usePersons";
import PersonTable from "../../components/PersonTable";
import AddPersonForm from "../../components/AddPersonForm";

const PersonsPage: React.FC = () => {
  const { persons, loading, error, addPerson, fetchPersons } = usePersons();

  const handleAddPersonSuccess = () => {
    fetchPersons(); // Re-fetch persons to update the table
  };

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-3xl font-bold mb-6">Persons Management</h1>

      <div className="mb-8">
        <AddPersonForm onAddPerson={addPerson} onSuccess={handleAddPersonSuccess} />
      </div>

      <div>
        <h2 className="text-2xl font-semibold mb-4">All Persons</h2>
        <PersonTable persons={persons} loading={loading} error={error} />
      </div>
    </div>
  );
};

export default PersonsPage;
