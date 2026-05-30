import { Person, CreatePersonRequest } from "../types/person";

const API_BASE_URL = "http://localhost:8080"; 

//MARK: GET ALL PERSONS
export const getPeople = async (): Promise<Person[]> => {
  const response = await fetch(`${API_BASE_URL}/person`);
  if (!response.ok) {
    throw new Error(`Error fetching people: ${response.statusText}`);
  }
  return response.json();
};

//MARK: GET A PERSON BY ID
export const getPersonById = async (id: string): Promise<Person> => {
  const response = await fetch(`${API_BASE_URL}/person/${id}`);
  if (!response.ok) {
    throw new Error(`Error fetching person with ID ${id}: ${response.statusText}`);
  }

  return response.json();
}

//MARK: CREATE A NEW PERSON
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

//MARK: UPDATE A PERSON
export const updatePerson = async (id: string, personData: Partial<CreatePersonRequest>): Promise<Person> => {
  const response = await fetch(`${API_BASE_URL}/person/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(personData),
  });

  if (!response.ok) {
    const errorBody = await response.text();
    throw new Error(
      `Error updating person with ID ${id}: ${response.statusText} - ${errorBody}`
    );
  }
  return response.json();
}
