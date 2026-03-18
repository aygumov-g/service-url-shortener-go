CREATE TABLE links (
    id BIGSERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    custom_code TEXT UNIQUE,
    click_count INTEGER NOT NULl,
    created_at TIMESTAMPTZ NOT NULL,
    last_accessed_at TIMESTAMPTZ
);
