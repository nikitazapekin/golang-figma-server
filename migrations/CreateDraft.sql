CREATE TABLE IF NOT EXISTS drafts (
    draft_id SERIAL PRIMARY KEY,
    draft_name VARCHAR(30),
    draft_description VARCHAR(256),
    likes INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    draft_author INT NOT NULL,
    CONSTRAINT fk_author FOREIGN KEY (draft_author) REFERENCES users(id) ON DELETE CASCADE
);
