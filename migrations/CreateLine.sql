CREATE TABLE IF NOT EXISTS lines (
    id SERIAL PRIMARY KEY,
    draft_id INT NOT NULL,
    coordX FLOAT NOT NULL,
    coordY FLOAT NOT NULL,
    type VARCHAR(50) NOT NULL,
    width FLOAT NOT NULL,
    height FLOAT NOT NULL,
    path JSONB NOT NULL,
    strokeWidth FLOAT NOT NULL,
    color VARCHAR(50) NOT NULL,
    layout INT NOT NULL,
    CONSTRAINT fk_draft FOREIGN KEY (draft_id) REFERENCES drafts(draft_id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_line_id ON lines USING HASH (id);
CREATE INDEX IF NOT EXISTS idx_draft_id ON lines USING HASH (draft_id);
