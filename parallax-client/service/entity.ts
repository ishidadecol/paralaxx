import { Entity } from "@/types/entity";

//TODO: REMOVE ALL HARDCODED URLS AND REPLACE WITH ENV VARIABLES
const API_BASE_URL = "http://localhost:8080";

export const getEntities = async (): Promise<Entity[]> => {
    const response = await fetch(`${API_BASE_URL}/entity`);
    if (!response.ok) {
        throw new Error(`Error fetching entities: ${response.statusText}`);
    }
    return response.json();
}