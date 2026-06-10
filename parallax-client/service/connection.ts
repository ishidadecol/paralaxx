import { CreateEntityConnectionRequest, EntityConnection, EntityConnectionDetail } from "@/types/entityConnection";

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

//MARK: GET CONNECTIONS FOR AN ENTITY
export const getConnectionsForEntity = async (entityId: string, typeFilter?: string[]): Promise<EntityConnectionDetail[]> => {
  const params = new URLSearchParams();

  typeFilter?.forEach(type =>
    params.append("type", type)
  );

  const response = await fetch(`${API_BASE_URL}/connection/entity/${entityId}?${params.toString()}`);
  if (!response.ok) {
    throw new Error(`Error fetching connections for entity ${entityId}: ${response.statusText}`);
  }
  return response.json();
}