CREATE TABLE IF NOT EXISTS  messages(
    message_id text PRIMARY KEY,
    school int NOT NULL,
    x         int NOT NULL,
    y         int NOT NULL,
    message text NOT NULL,
    created_at text NOT NULL,
    float_time real NOT NULL,
    likes int NOT NULL DEFAULT 0,
    parent_id text
);
