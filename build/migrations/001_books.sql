-- +goose Up
CREATE TABLE books (
  id BIGSERIAL PRIMARY KEY,
  name text NOT NULL,
  author text NOT NULL,
  publication_date timestamp without time zone NOT NULL,

  constraint u_book unique (name, author, publication_date)
);

-- +goose Down
DROP TABLE books;