CREATE TABLE IF NOT EXISTS comment (
  id UUID,
  body VARCHAR NOT NULL,
  created_at TIMESTAMP NOT NULL,
  article_id UUID NOT NULL,
  author_id UUID NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS category (
  id UUID,
  name VARCHAR NOT NULL,
  author_id UUID,
  slug VARCHAR NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS article (
    id UUID,
    slug VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    body VARCHAR NOT NULL,
    status VARCHAR,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    author_id UUID,
    comments_count INT,
    image_url VARCHAR,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS article_category (
    id UUID,
    category_id UUID,
    article_id UUID,
    PRIMARY KEY (id)
);
