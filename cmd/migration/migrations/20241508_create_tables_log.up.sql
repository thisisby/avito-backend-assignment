CREATE TABLE IF NOT EXISTS log (
    id uuid PRIMARY KEY,
    user_agent VARCHAR(255) NOT NULL,
    request_id VARCHAR(255) NULL,
    random_value VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    count INT NOT NULL
)