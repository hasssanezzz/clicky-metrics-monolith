CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(1024) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE "url" (
    id SERIAL PRIMARY KEY,
    user_username VARCHAR(255) REFERENCES "user"(username) ON DELETE SET NULL,
    short VARCHAR(255) NOT NULL UNIQUE,
    long VARCHAR(1024) NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE "url_analytics" (
    id SERIAL PRIMARY KEY,
    url_id INT REFERENCES "url"(id) ON DELETE CASCADE,
    accessed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ip_address VARCHAR(45),
    user_agent TEXT,
    browser VARCHAR(100),
    device VARCHAR(100),
    location VARCHAR(255)
);

CREATE INDEX IF NOT EXISTS idx_user_username ON "user"(username);
CREATE INDEX IF NOT EXISTS idx_url_username ON "url"(user_username);
CREATE INDEX IF NOT EXISTS idx_url_short ON "url"(short);
CREATE INDEX IF NOT EXISTS idx_url_id ON "url_analytics"(url_id);