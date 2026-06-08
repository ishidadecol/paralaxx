export type EntityConnection = {
    id : string,
    sourceEntityId: string,
    targetEntityId: string,
    relationshipType: string,
    startDate?: string, // api will return as date though
    endDate?: string, // api will return as date though
    createdAt: string,
    updatedAt: string,
}

export type CreateEntityConnectionRequest = {
    sourceEntityId: string,
    targetEntityId: string,
    relationshipType: string,
    startDate?: string | null, 
    endDate?: string | null, 
}