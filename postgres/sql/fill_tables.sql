INSERT INTO authors(name) VALUES ('Dmitry');
INSERT INTO authors(name) VALUES ('Evgeniy');
INSERT INTO authors(name) VALUES ('Aleksey');
INSERT INTO authors(name) VALUES ('Elizaveta');
INSERT INTO authors(name) VALUES ('Ekaterina');

INSERT INTO books(title,genre,year) VALUES ('Book', 'Humor', '2004');
INSERT INTO books(title,genre,year) VALUES ('Book2', 'Mystery', '2003');
INSERT INTO books(title,genre,year) VALUES ('Book3', 'Classic', '2002');
INSERT INTO books(title,genre,year) VALUES ('Book4', 'Horror', '2001');
INSERT INTO books(title,genre,year) VALUES ('Book5', 'Western', '2005');

INSERT INTO books_authors(book_id, author_id) VALUES (1,1);
INSERT INTO books_authors(book_id, author_id) VALUES (1,2);
INSERT INTO books_authors(book_id, author_id) VALUES (1,3);
INSERT INTO books_authors(book_id, author_id) VALUES (2,1);
INSERT INTO books_authors(book_id, author_id) VALUES (2,2);
INSERT INTO books_authors(book_id, author_id) VALUES (2,3);
INSERT INTO books_authors(book_id, author_id) VALUES (4,5);
INSERT INTO books_authors(book_id, author_id) VALUES (5,4);