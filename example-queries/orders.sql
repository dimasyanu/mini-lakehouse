CREATE TABLE IF NOT EXISTS iceberg.tenant_a.orders (
    id UUID NOT NULL,
    description VARCHAR(512),
    amount INT NOT NULL,
    created_at TIMESTAMP
);
