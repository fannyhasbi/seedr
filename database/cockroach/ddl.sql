CREATE TABLE IF NOT EXISTS comment (
  id UUID,
  body VARCHAR NOT NULL,
  created_at TIMESTAMP NOT NULL,
  article_id UUID NOT NULL,
  author_id UUID NOT NULL,
  PRIMARY KEY (id)
)