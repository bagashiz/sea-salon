-- +goose Up
CREATE TYPE account_role AS ENUM ('admin', 'customer');
CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    role account_role NOT NULL DEFAULT 'customer',
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()   
);
INSERT INTO accounts (
    role,
    email,
    password,
    full_name,
    phone_number
) VALUES (
    'admin',
    'thomas.n@compfest.id',
    '$2a$10$4tS9MQtS6l/9PWY.MiR8O.3.yFKHvaB34kpQVGgSVnGrla6ztOaam',
    'Thomas N',
    '08123456789'
);

-- +goose Down
DROP TABLE IF EXISTS accounts;
DROP TYPE IF EXISTS accounts_role;
