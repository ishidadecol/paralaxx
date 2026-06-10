-- Add display_name to entity
ALTER TABLE entity
ADD COLUMN display_name TEXT;

-- Backfill existing person entities
UPDATE entity e
SET display_name = CONCAT(p.first_name, ' ', p.last_name)
FROM person p
WHERE p.entity_id = e.id;

-- Make display_name required
ALTER TABLE entity
ALTER COLUMN display_name SET NOT NULL;