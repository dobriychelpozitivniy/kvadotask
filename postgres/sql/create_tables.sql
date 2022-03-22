CREATE TABLE IF NOT EXISTS books (
  book_id SERIAL PRIMARY KEY NOT NULL,
  title TEXT NOT NULL,
  genre TEXT NOT NULL,
  year INT NOT NULL
);

CREATE TABLE IF NOT EXISTS authors (
  author_id SERIAL PRIMARY KEY NOT NULL,
  name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS books_authors (
	book_id INT NOT NULL REFERENCES books(book_id) ON DELETE CASCADE,
	author_id INT NOT NULL REFERENCES authors(author_id) ON DELETE CASCADE,
	UNIQUE (book_id, author_id)
);

CREATE INDEX books_authors_author_id_idx ON books_authors(author_id);
CREATE INDEX books_authors_book_id_idx ON books_authors(book_id);