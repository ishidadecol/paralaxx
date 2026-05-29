import { Person, CreatePersonRequest } from "../types/person";

const API_BASE_URL = "http://localhost:8080"; 

export const getPeople = async (): Promise<Person[]> => {
  const response = await fetch(`${API_BASE_URL}/person`);
  if (!response.ok) {
    throw new Error(`Error fetching people: ${response.statusText}`);
  }
  return response.json();
};

export const createPerson = async (
  personData: CreatePersonRequest
): Promise<Person> => {
  const response = await fetch(`${API_BASE_URL}/person`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(personData),
  });

  if (!response.ok) {
    const errorBody = await response.text();
    throw new Error(
      `Error creating person: ${response.statusText} - ${errorBody}`
    );
  }
  return response.json();
};
