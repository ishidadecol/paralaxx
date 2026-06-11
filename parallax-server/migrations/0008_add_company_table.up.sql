CREATE TABLE company (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    entity_id UUID NOT NULL UNIQUE,

    name TEXT NOT NULL,
    legal_name TEXT,
    cnpj TEXT,

    industry TEXT,
    website TEXT,
    description TEXT,

    start_date DATE,
    end_date DATE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_company_entity
        FOREIGN KEY (entity_id)
        REFERENCES entity(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_company_entity_id
ON company(entity_id);

CREATE INDEX idx_company_name
ON company(name);

CREATE INDEX idx_company_cnpj
ON company(cnpj);