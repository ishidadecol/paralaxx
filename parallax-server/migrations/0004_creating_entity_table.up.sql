
-- CRREATE TABLE FOR ENTITY

CREATE TABLE entity (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    type TEXT NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- ADD entity_id TO person
ALTER TABLE person
ADD COLUMN entity_id UUID;

-- CREATE ENTITIES FOR EXISTING PERSONS
INSERT INTO entity (
    id,
    type,
    created_at,
    updated_at
)
SELECT
    id,
    'person',
    created_at,
    updated_at
FROM person;

-- LINK THOSE PERSONS TO THEIR NEW ENTITY
UPDATE person
SET entity_id = id;

-- ADD FOREIGN KEY CONSTRAINT
ALTER TABLE person
ADD CONSTRAINT fk_person_entity
FOREIGN KEY (entity_id)
REFERENCES entity(id);

-- MAKE entity_id required
ALTER TABLE person
ALTER COLUMN entity_id SET NOT NULL;


