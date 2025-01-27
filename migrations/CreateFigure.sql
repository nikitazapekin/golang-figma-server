CREATE TABLE IF NOT EXISTS figures (
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

-- Создаем хеш-индексы для быстрого доступа
CREATE INDEX idx_figure_id ON figures USING HASH (figure_id);
CREATE INDEX idx_draft_id ON figures USING HASH (draft_id);
