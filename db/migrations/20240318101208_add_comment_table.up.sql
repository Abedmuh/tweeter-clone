CREATE TABLE IF NOT EXISTS comments (
  id VARCHAR(255) PRIMARY KEY,
  creator VARCHAR(255) NOT NULL,
  post_id VARCHAR(255) NOT NULL,
  comment_in_html TEXT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
  FOREIGN KEY (creator) REFERENCES users(id) ON DELETE CASCADE
);
