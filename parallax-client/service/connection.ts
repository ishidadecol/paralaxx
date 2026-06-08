import { CreateEntityConnectionRequest, EntityConnection } from "@/types/entityConnection";

const API_BASE_URL = "http://localhost:8080"; 

//MARK: GET ALL CONNECTIONS
export const getConnections = async (): Promise<EntityConnection[]> => {
  const response = await fetch(`${API_BASE_URL}/connection`);
  if (!response.ok) {
    throw new Error(`Error fetching connections: ${response.statusText}`);
  }
  return response.json();
};

//MARK: CREATE A NEW CONNECTION
export const createConnection = async (
  connectionData: CreateEntityConnectionRequest
): Promise<EntityConnection> => {
  const response = await fetch(`${API_BASE_URL}/connection`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(connectionData),
  });

  if (!response.ok) {
    const errorBody = await response.text();
    throw new Error(
      `Error creating connection: ${response.statusText} - ${errorBody}`
    );
  }
  return response.json();
};