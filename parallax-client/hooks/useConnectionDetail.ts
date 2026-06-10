import { getConnectionsForEntity } from "@/service/connection";
import { EntityConnectionDetail } from "@/types/entityConnection";
import { useEffect, useState } from "react";

export const useEntityConnectionDetail = (entityId: string, typeFilter?: string[]) => {
    const [connectionDetail, setConnectionDetail] = useState<EntityConnectionDetail[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchConnectionDetail = async () => {
            try {
                setLoading(true);

                const data = await getConnectionsForEntity(
                    entityId,
                    typeFilter
                );

                setConnectionDetail(data);
            } catch (err) {
                setError(err instanceof Error ? err.message : 'Unknown error');
            } finally {
                setLoading(false);
            }
        };

        fetchConnectionDetail();
    }, [entityId, typeFilter]);

    return { connectionDetail, loading, error };
}