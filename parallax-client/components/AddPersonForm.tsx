import React, { useState } from "react";
import { CreatePersonRequest } from "../types/person";
import { Card } from "./ui/card";
import { Button } from "./ui/button";

interface AddPersonFormProps {
  onAddPerson: (person: CreatePersonRequest) => Promise<void>;
  onSuccess: () => void;
}

const AddPersonForm: React.FC<AddPersonFormProps> = ({ onAddPerson, onSuccess }) => {
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [gender, setGender] = useState("");
  const [birthDate, setBirthDate] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  // Handle form submission
  const handleSubmit = async (e: React.SubmitEvent) => {
    e.preventDefault();
    setLoading(true);
    setError(null);

    const personData: CreatePersonRequest = {
      first_name: firstName,
      last_name: lastName || undefined,
      gender: gender || undefined,
      birth_date: birthDate || undefined,
    };

    try {
      await onAddPerson(personData);
      setFirstName("");
      setLastName("");
      setGender("");
      setBirthDate("");
      onSuccess();
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to add person");
    } finally {
      setLoading(false);
    }
  };

  return (
    <Card className="cn-card group/card flex flex-col">
    <form onSubmit={handleSubmit} className="p-4 rounded-lg shadow-md">
      <h2 className="text-xl font-semibold mb-4">Add New Person</h2>
      {error && <p className="text-red-500 mb-4">{error}</p>}
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
        <div>
          <label htmlFor="firstName" className="block text-sm font-medium ">
            First Name:
          </label>
          <input
            type="text"
            id="firstName"
            value={firstName}
            onChange={(e) => setFirstName(e.target.value)}
            required
            className="mt-1 block w-full p-2 border rounded-md shadow-sm"
          />
        </div>
        <div>
          <label htmlFor="lastName" className="block text-sm font-medium">
            Last Name:
          </label>
          <input
            type="text"
            id="lastName"
            value={lastName}
            onChange={(e) => setLastName(e.target.value)}
            className="mt-1 block w-full p-2 border shadow-sm"
          />
        </div>
        <div>
          <label htmlFor="gender" className="block text-sm font-medium">
            Gender:
          </label>
          <input
            type="text"
            id="gender"
            value={gender}
            onChange={(e) => setGender(e.target.value)}
            className="mt-1 block w-full p-2 border  shadow-sm "
          />
        </div>
        <div>
          <label htmlFor="birthDate" className="block text-sm font-medium">
            Birth Date:
          </label>
          <input
            type="date"
            id="birthDate"
            value={birthDate}
            onChange={(e) => setBirthDate(e.target.value)}
            className="mt-1 block w-full p-2 border shadow-sm"
          />
        </div>
      </div>
      <Button className="w-full py-2 px-4">
        {loading ? "Adding..." : "Add Person"}
      </Button>
    </form>
    </Card>
  );
};

export default AddPersonForm;
