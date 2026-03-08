-- SQL script to create the database schema for the inventory management system

-- Create items table
CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    category VARCHAR(255) NOT NULL
);

-- Create accounts table
CREATE TABLE IF NOT EXISTS accounts (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);

-- Create congregations table
CREATE TABLE IF NOT EXISTS congregations (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL
);

-- Optional: Insert some sample data (uncomment if needed)
-- INSERT INTO items (name, description, price, quantity, category) VALUES
-- ('Laptop', 'High-performance laptop for professionals', 1500, 10, 'Electronics'),
-- ('Wireless Mouse', 'Ergonomic wireless mouse with precision tracking', 50, 100, 'Electronics'),
-- ('Office Chair', 'Comfortable ergonomic office chair', 300, 25, 'Furniture');
--
-- INSERT INTO accounts (id, name, email) VALUES
-- ('550e8400-e29b-41d4-a716-446655440000', 'Alice Smith', 'alice@example.com'),
-- ('550e8400-e29b-41d4-a716-446655440001', 'Bob Johnson', 'bob@example.com');
--
-- INSERT INTO congregations (id, name, location) VALUES
-- ('550e8400-e29b-41d4-a716-446655440002', 'First Congregation', 'City A'),
-- ('550e8400-e29b-41d4-a716-446655440003', 'Second Congregation', 'City B');