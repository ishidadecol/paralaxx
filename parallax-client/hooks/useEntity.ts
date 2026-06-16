import { useState, useEffect } from "react";
import { EntityLookUp } from "@/types/entity";
import { getEntitiesDisplayName } from "@/service/entity";

export const useEntity = (filter: string[] = []) => {
    const [entityDisplayNames, setEntityDisplayNames] = useState<EntityLookUp[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const loadEntityDisplayNames = async () => {
            setLoading(true);
            setError(null);
            console.log("filter", filter);
            try {
                const data = await getEntitiesDisplayName(filter);
                setEntityDisplayNames(data);
            } catch (err) {
                setError(
                    err instanceof Error
                        ? err.message
                        : "An unknown error occurred"
                );
            } finally {
                setLoading(false);
            }
        };

        loadEntityDisplayNames();
    }, [filter]);

    return {
        entityDisplayNames,
        loading,
        error,
    };
};