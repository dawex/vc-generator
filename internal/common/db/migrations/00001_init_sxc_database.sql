-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS event (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    contract_id VARCHAR(256) NOT NULL,
    execution_id VARCHAR(256) NOT NULL,
    source VARCHAR(256) NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL,
    metric VARCHAR(256) NOT NULL,
    value VARCHAR(1024) NOT NULL,
    log VARCHAR(1024) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS verifiable_credential (
    id VARCHAR(256) PRIMARY KEY,
    context JSONB NOT NULL,
    type JSONB NOT NULL,
    credential_subject JSONB NOT NULL,
    issuer_id VARCHAR(256) NOT NULL,
    issuer_name VARCHAR(256) NOT NULL,
    issuance_date TIMESTAMPTZ NOT NULL,
    proof_created TIMESTAMPTZ NOT NULL,
    proof_jws VARCHAR(1024) NOT NULL,
    proof_purpose VARCHAR(256) NOT NULL,
    proof_type VARCHAR(256) NOT NULL,
    proof_verification_method VARCHAR(256) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS event;
DROP TABLE IF EXISTS verifiable_credential;