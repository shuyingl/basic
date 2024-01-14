/*
    Purpose: Create basic schema for the project
*/

DO $$
BEGIN
    -- Check if the table 'users' exists
    IF NOT EXISTS (
        SELECT 1
        FROM information_schema.tables
        WHERE table_name = 'users'
    ) THEN
        -- Create the 'users' table
        CREATE TABLE users (
            id UUID PRIMARY KEY,
            email TEXT UNIQUE NOT NULL
        );
    END IF;
END
$$;
