CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE issues (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    recommendation TEXT,
    deleted_at TIMESTAMP NULL
);