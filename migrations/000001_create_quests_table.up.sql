CREATE TABLE IF NOT EXISTS quests (
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    description text NOT NULL,
    reward integer NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone,
    version integer NOT NULL DEFAULT 1
);