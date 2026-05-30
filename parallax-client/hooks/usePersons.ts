import { useState, useEffect, useCallback } from "react";
import { Person, CreatePersonRequest } from "../types/person";
import { getPeople, createPerson } from "../service/person";

export const usePersons = () => {
  const [persons, setPersons] = useState<Person[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  // Initial data fetch on mount
  useEffect(() => {
    const loadPersons = async () => {
      setLoading(true);
      setError(null);
      try {
        const data = await getPeople();
        setPersons(data);
      } catch (err) {
        setError(err instanceof Error ? err.message : "An unknown error occurred");
      } finally {
        setLoading(false);
      }
    };
    loadPersons();
  }, []); // Empty dependency array means it runs once on mount

  const addPerson = useCallback(
    async (personData: CreatePersonRequest) => {
      try {
        await createPerson(personData);
        // No need to update persons state here, fetchPersons will do it
      } catch (err) {
        setError(err instanceof Error ? err.message : "An unknown error occurred");
        throw err; // Re-throw to allow component to handle
      }
    },
    []
  );

  // This function is now for manual re-fetching, not for initial load
  const fetchPersons = useCallback(async () => {
    setLoading(true);
    setError(null);
    try {
      const data = await getPeople();
      setPersons(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : "An unknown error occurred");
    } finally {
      setLoading(false);
    }
  }, []);

  return { persons, loading, error, addPerson, fetchPersons };
};
