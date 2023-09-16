ALTER TABLE IF EXISTS brewery ADD COLUMN IF NOT EXISTS creator_id uuid;
ALTER TABLE brewery ADD FOREIGN KEY (creator_id) REFERENCES users (id);