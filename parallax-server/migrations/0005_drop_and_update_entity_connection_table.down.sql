DROP INDEX IF EXISTS idx_connection_source;
DROP INDEX IF EXISTS idx_connection_target;
DROP INDEX IF EXISTS idx_connection_relationship;

DROP TABLE IF EXISTS entity_connection;

CREATE TABLE entity_connection (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    source_type TEXT NOT NULL,
    source_id UUID NOT NULL,

    target_type TEXT NOT NULL,
    target_id UUID NOT NULL,

    relationship_type TEXT NOT NULL,

    start_date DATE,
    end_date DATE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_connection_source
ON entity_connection(source_type, source_id);

CREATE INDEX idx_connection_target
ON entity_connection(target_type, target_id);

CREATE INDEX idx_connection_relationship
ON entity_connection(relationship_type);