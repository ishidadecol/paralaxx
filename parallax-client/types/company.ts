export type Company = {
    entity_id: string;
    name : string;
    legal_name?: string;
    cnpj?: string;
    industry?: string;
    website?: string;
    description?: string;
}

export type CreateCompanyRequest = {
    entity_id: string;
    name : string;
    legal_name?: string;
    cnpj?: string;
    industry?: string;
    website?: string;
    description?: string;
}