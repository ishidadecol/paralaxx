-- Drop old indexes
DROP INDEX IF EXISTS idx_connection_source;
DROP INDEX IF EXISTS idx_connection_target;
DROP INDEX IF EXISTS idx_connection_relationship;

-- Drop old table
DROP TABLE IF EXISTS entity_connection;

-- Create new table
CREATE TABLE entity_connection (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    source_entity_id UUID NOT NULL,
    target_entity_id UUID NOT NULL,

    relationship_type TEXT NOT NULL,

    start_date DATE,
    end_date DATE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_source_entity
        FOREIGN KEY (source_entity_id)
        REFERENCES entity(id),

    CONSTRAINT fk_target_entity
        FOREIGN KEY (target_entity_id)
        REFERENCES entity(id)
);

CREATE INDEX idx_connection_source
ON entity_connection(source_entity_id);

CREATE INDEX idx_connection_target
ON entity_connection(target_entity_id);

CREATE INDEX idx_connection_relationship
ON entity_connection(relationship_type);