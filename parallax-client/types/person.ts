export type Person = {
  id: string;
  first_name: string;
  last_name: string;
  birth_date?: string; // Using string for date for simplicity, can be Date object
  gender: string;
  created_at: string;
  updated_at: string;
};

export type CreatePersonRequest = {
  first_name: string;
  last_name?: string;
  gender?: string;
  birth_date?: string; // Using string for date for simplicity
};
