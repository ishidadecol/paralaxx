import { createCompany, getCompanies } from "@/service/company";
import { Company, CreateCompanyRequest } from "@/types/company";
import { useCallback, useEffect, useState } from "react"


export const useCompany = () =>{
    const [companies, setCompanies] = useState<Company[]>([]);
    const [company, setCompany] = useState<Company | null>(null);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() =>{
        const loadCompanies = async () =>{
            setLoading(true);
            setError(null);
            try{
                const data = await getCompanies();
                setCompanies(data);
            } catch (err) {
                setError(
                    err instanceof Error ? err.message : "An unknown error occurred"
                );
            } finally {
                setLoading(false);
            }
        };

        loadCompanies();
    }, []);

    //MARK: CREATE COMPANY HOOK
    const addCompany = useCallback(async (companyData: CreateCompanyRequest) => {
        try{
            await createCompany(companyData);
        } catch (err) {
            setError(
                err instanceof Error ? err.message : "An unknown error occurred"
            );
            throw err;
        }
    }, []);

    return {
        company,
        companies,
        loading,
        error,
        addCompany
    }

}