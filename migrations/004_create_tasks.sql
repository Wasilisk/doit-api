-- +goose Up
CREATE TABLE tasks (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id      UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name         VARCHAR(255) NOT NULL,
    description  TEXT,
    date         TIMESTAMP,
    time_start   TIMESTAMP,
    time_end     TIMESTAMP,
    is_completed BOOLEAN NOT NULL DEFAULT FALSE,
    is_favourite BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at   TIMESTAMP,
    created_at   TIMESTAMP DEFAULT NOW(),
    updated_at   TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE tasks;