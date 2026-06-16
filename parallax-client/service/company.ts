import { Company, CreateCompanyRequest } from "@/types/company"

const API_BASE_URL = "http://localhost:8080"

//MARK: GET ALL COMPANIES
export const getCompanies = async (): Promise<Company[]> => {
    const response = await fetch(`${API_BASE_URL}/company`)
    if(!response.ok){
        throw new Error(`Error fetching people: ${response.statusText}`);
    }
    return response.json()
}


//MARK: CREATE COMPANY
export const createCompany = async (companyData: CreateCompanyRequest): Promise<Company> => {
    const response = await fetch(`${API_BASE_URL}/company`,{
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(companyData),
    });

    if(!response.ok){
        throw new Error(`Error fetching people: ${response.statusText}`);
    }

    return response.json()
}