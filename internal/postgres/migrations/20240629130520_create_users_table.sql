-- +goose Up
CREATE TYPE user_role AS ENUM ('admin', 'customer');
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    role user_role NOT NULL DEFAULT 'customer',
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()   
);
INSERT INTO users (
    role,
    email,
    password,
    full_name,
    phone
) VALUES (
    'admin',
    'thomas.n@compfest.id',
    '$2a$10$4tS9MQtS6l/9PWY.MiR8O.3.yFKHvaB34kpQVGgSVnGrla6ztOaam',
    'Thomas N',
    '08123456789'
);

-- +goose Down
DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS user_role;