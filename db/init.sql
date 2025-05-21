-- CREATE DATABASE IF NOT EXISTS peya
SELECT 'CREATE DATABASE transactionDB'
    WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'transactionDB')\gexec