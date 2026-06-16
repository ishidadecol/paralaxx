import { Entity, EntityLookUp } from "@/types/entity";

//TODO: REMOVE ALL HARDCODED URLS AND REPLACE WITH ENV VARIABLES
const API_BASE_URL = "http://localhost:8080";

export const getEntities = async (): Promise<Entity[]> => {
    const response = await fetch(`${API_BASE_URL}/entity`);
    if (!response.ok) {
        throw new Error(`Error fetching entities: ${response.statusText}`);
    }
    return response.json();
}

export const getEntitiesDisplayName = async (filter: string[] = []): Promise<EntityLookUp[]> => {
    const params = new URLSearchParams();
    filter.forEach(f => params.append('filter', f));
    const queryString = params.toString();
    const url = `${API_BASE_URL}/entity/display-names${queryString ? `?${queryString}` : ''}`;

    const response = await fetch(url);
    if (!response.ok) {
        throw new Error(`Error fetching entity display names: ${response.statusText}`);
    }
    return response.json();
}