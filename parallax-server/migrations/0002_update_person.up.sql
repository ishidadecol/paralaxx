ALTER TABLE people
RENAME TO person;

ALTER TABLE person
RENAME COLUMN name TO first_name;

ALTER TABLE person
ADD COLUMN last_name TEXT;

ALTER TABLE person
ADD COLUMN birth_date DATE;

ALTER TABLE person
ADD COLUMN gender TEXT;