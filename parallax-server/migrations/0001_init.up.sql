CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- =========================================================
-- PEOPLE
-- =========================================================

CREATE TABLE people (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    name TEXT NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
