import { getConnections, createConnection} from "@/service/connection";
import { CreateEntityConnectionRequest, EntityConnection } from "@/types/entityConnection";
import { useCallback, useEffect, useState } from "react";

export const useEntityConnection = () => {
  const [connections, setConnections] = useState<EntityConnection[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  //Initial fetch of connections on component mount
  useEffect(() => {
    const loadConnections = async () => {
      setLoading(true);
      setError(null);
      try {
        const data = await getConnections();
        setConnections(data);
      } catch (err) {
        setError(
          err instanceof Error ? err.message : "An unknown error occurred"
        );
      } finally {
        setLoading(false);
      }
    };
    loadConnections();
  }, []); // Empty dependency array means it runs once on mount

  const fetchEntityConnections = useCallback(async () => {
    setLoading(true);
    setError(null);

    try {
      const data = await getConnections();
      setConnections(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Unknown error");
      return [];
    } finally {
      setLoading(false);
    }
  }, []);

  const addEntityConnection = useCallback(
    async (connection: CreateEntityConnectionRequest) => {
        try{
            await createConnection(connection)
        } catch (err){
            setError(err instanceof Error ? err.message : "Unknown error");
            throw err; // Re-throw to allow component to handle
        }
    },
    []
  );

  return {
    connections,
    loading,
    error,
    fetchEntityConnections,
    addEntityConnection,
  };
};
