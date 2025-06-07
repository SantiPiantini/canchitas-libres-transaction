-- CREATE DATABASE IF NOT EXISTS peya

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE transactions (
                        transaction_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        payment_id VARCHAR(50) NOT NULL,
                        user_id VARCHAR(100) NOT NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

