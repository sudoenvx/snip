-- migrations/000001_create_urls_table.up.sql
CREATE TABLE IF NOT EXISTS urls (
    id BIGSERIAL PRIMARY KEY,
    code VARCHAR(50) UNIQUE NOT NULL,
    original_url TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    clicks BIGINT DEFAULT 0,
    expires_at TIMESTAMPTZ,
    last_accessed TIMESTAMPTZ
);

CREATE INDEX idx_urls_code ON urls(code);
CREATE INDEX idx_urls_expires_at ON urls(expires_at);