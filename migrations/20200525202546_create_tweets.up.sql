CREATE TABLE IF NOT EXISTS tweets (
    id serial PRIMARY KEY,
    uuid UUID NOT NULL,
    tweet_id VARCHAR (100) UNIQUE NOT NULL,
    created_at DATE,
    deleted_at DATE,
    updated_at DATE
)
