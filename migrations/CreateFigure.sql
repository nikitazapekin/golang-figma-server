CREATE TABLE IF NOT EXISTS figures (
    id SERIAL PRIMARY KEY,
    draft_id INT NOT NULL,
    coordX FLOAT NOT NULL,
    coordY FLOAT NOT NULL,
    type VARCHAR(50) NOT NULL,
    width FLOAT NOT NULL,
    height FLOAT NOT NULL,
    opacity FLOAT NOT NULL,
    border FLOAT NOT NULL,
    background VARCHAR(50) NOT NULL,
    stroke FLOAT NOT NULL,
    strokeColor VARCHAR(50) NOT NULL,
    shadowColor VARCHAR(50) NOT NULL,
    shadowX FLOAT NOT NULL,
    shadowY FLOAT NOT NULL,
    layout INT NOT NULL,
    CONSTRAINT fk_draft FOREIGN KEY (draft_id) REFERENCES drafts(draft_id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_figure_id ON figures USING HASH (id);
CREATE INDEX IF NOT EXISTS idx_draft_id ON figures USING HASH (draft_id);

/* CREATE TABLE IF NOT EXISTS figures (
    figure_id SERIAL PRIMARY KEY,
    draft_id INT NOT NULL,
    figure_coordX FLOAT NOT NULL,
    figure_coordY FLOAT NOT NULL,
    figure_color VARCHAR(50) NOT NULL,
    figure_radius FLOAT NOT NULL,
    figure_opacity FLOAT NOT NULL,
    figure_layout VARCHAR(50) NOT NULL,
    figure_width FLOAT NOT NULL,
    figure_height FLOAT NOT NULL,
    figure_text TEXT DEFAULT NULL,
    figure_text_color VARCHAR(50) DEFAULT NULL,
    figure_text_family VARCHAR(100) DEFAULT NULL,
    figure_text_font VARCHAR(100) DEFAULT NULL,
    CONSTRAINT fk_draft FOREIGN KEY (draft_id) REFERENCES drafts(draft_id) ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS idx_figure_id ON figures USING HASH (figure_id);
CREATE INDEX IF NOT EXISTS idx_draft_id ON figures USING HASH (draft_id);
 */