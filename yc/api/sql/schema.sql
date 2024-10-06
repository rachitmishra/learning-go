CREATE TABLE users (
    id SERIAL PRIMARY KEY,            -- Unique ID for each user
    username VARCHAR(50) NOT NULL UNIQUE,  -- Username, must be unique
    email VARCHAR(100) NOT NULL UNIQUE,    -- Email, must be unique
    password_hash TEXT NOT NULL,           -- Password stored as hash
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Record creation timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp when updated
    is_active BOOLEAN DEFAULT TRUE         -- Status of the account
);

CREATE TABLE links (
    id SERIAL PRIMARY KEY,              -- Unique ID for each link
    user_id INTEGER REFERENCES users(user_id), -- Foreign key referencing users table
    original_url TEXT NOT NULL,              -- The original full URL
    short_url VARCHAR(50) NOT NULL UNIQUE,   -- The shortened URL
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp when the link is created
    clicks INTEGER DEFAULT 0,                -- Count of how many times the link was clicked
    is_active BOOLEAN DEFAULT TRUE           -- Status of the link (active/inactive)
);