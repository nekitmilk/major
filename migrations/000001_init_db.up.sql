CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE severities (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    level INT NOT NULL,
    description TEXT
);

CREATE TABLE issues (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    recommendation TEXT,
    severity_id UUID NOT NULL REFERENCES severities(id)
);

CREATE TABLE rules (
   id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
   name VARCHAR(255) NOT NULL,
   description TEXT NOT NULL,
   issue_id UUID NOT NULL REFERENCES issues(id),
   expression TEXT NOT NULL,
   enabled BOOLEAN NOT NULL
);