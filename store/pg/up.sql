CREATE TABLE posts (
  id text NOT NULL PRIMARY KEY,
  createtime text NOT NULL,
  link text NOT NULL
);

CREATE TABLE tags (
  post_id text REFERENCES posts(id),
  tag text NOT NULL
);
