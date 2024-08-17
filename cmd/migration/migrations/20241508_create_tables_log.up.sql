CREATE TABLE IF NOT EXISTS logs
(
    token_id   uuid PRIMARY KEY NOT NULL,
    token      VARCHAR(255)     NOT NULL,
    user_agent VARCHAR(255)     NOT NULL,
    url        VARCHAR(255)     NOT NULL,
    count      INT              NOT NULL DEFAULT 0
)