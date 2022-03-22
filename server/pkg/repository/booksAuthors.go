package repository

import (
	"context"
	"fmt"
	pb "taskserver/pkg/grpc/proto"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type BooksAuthorsRepository struct {
	db *pgxpool.Pool
}

func NewBooksAuthors(conn *pgxpool.Pool) *BooksAuthorsRepository {
	return &BooksAuthorsRepository{db: conn}
}

// Return author's book by book ID from DB.
func (r *BooksAuthorsRepository) FindAuthorsByBookID(bookID int) ([]*pb.FindAuthorsResponse_Author, error) {
	sql := "select authors.author_id,authors.name from books_authors JOIN authors ON books_authors.author_id = authors.author_id WHERE books_authors.book_id = $1"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	rows, err := r.db.Query(ctx, sql, bookID)
	if err != nil {
		return nil, fmt.Errorf("Error query to db: %s", err)
	}

	defer rows.Close()

	var authors []*pb.FindAuthorsResponse_Author
	for rows.Next() {
		var author pb.FindAuthorsResponse_Author

		err = rows.Scan(&author.AuthorId, &author.Name)
		if err != nil {
			return nil, fmt.Errorf("Error scan row from db: %s", err)
		}

		authors = append(authors, &author)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("Row error: %s", rows.Err())
	}

	if len(authors) == 0 {
		return nil, fmt.Errorf("Empty result from query")
	}

	return authors, nil
}

// Return book's author by author ID from DB.
func (r *BooksAuthorsRepository) FindBooksByAuthorID(authorID int) ([]*pb.FindBooksResponse_Book, error) {
	sql := "select books.book_id,books.title ,books.genre,books.year from books_authors INNER JOIN books ON books_authors.book_id = books.book_id WHERE books_authors.author_id = $1"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	rows, err := r.db.Query(ctx, sql, authorID)
	if err != nil {
		return nil, fmt.Errorf("Error query to db: %s", err)
	}

	defer rows.Close()

	var books []*pb.FindBooksResponse_Book
	for rows.Next() {
		var book pb.FindBooksResponse_Book

		err = rows.Scan(&book.BookId, &book.Name, &book.Genre, &book.Year)
		if err != nil {
			return nil, fmt.Errorf("Error scan row from db: %s", err)
		}

		books = append(books, &book)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("Row error: %s", rows.Err())
	}

	if len(books) == 0 {
		return nil, fmt.Errorf("Empty result from query")
	}

	return books, nil
}

// Return book's author by author name from DB.
func (r *BooksAuthorsRepository) FindBooksByAuthorName(authorName string) ([]*pb.FindBooksResponse_Book, error) {
	sql := "select books.book_id, books.title, books.genre, books.year FROM authors JOIN books_authors ON authors.author_id = books_authors.author_id JOIN books ON books_authors.book_id = books.book_id WHERE authors.name = $1"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	rows, err := r.db.Query(ctx, sql, authorName)
	if err != nil {
		return nil, fmt.Errorf("Error query to db: %s", err)
	}

	defer rows.Close()

	var books []*pb.FindBooksResponse_Book
	for rows.Next() {
		var book pb.FindBooksResponse_Book

		err = rows.Scan(&book.BookId, &book.Name, &book.Genre, &book.Year)
		if err != nil {
			return nil, fmt.Errorf("Error scan row from db: %s", err)
		}

		books = append(books, &book)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("Row error: %s", rows.Err())
	}

	if len(books) == 0 {
		return nil, fmt.Errorf("Empty result from query")
	}

	return books, nil
}

// Return author's book by book name from DB.
func (r *BooksAuthorsRepository) FindAuthorByBookName(bookName string) ([]*pb.FindAuthorsResponse_Author, error) {
	sql := "select authors.author_id, authors.name from books JOIN books_authors ON books.book_id = books_authors.book_id JOIN authors ON books_authors.author_id = authors.author_id WHERE books.title = $1"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	rows, err := r.db.Query(ctx, sql, bookName)
	if err != nil {
		return nil, fmt.Errorf("Error query to db: %s", err)
	}

	defer rows.Close()

	var authors []*pb.FindAuthorsResponse_Author
	for rows.Next() {
		var author pb.FindAuthorsResponse_Author

		err = rows.Scan(&author.AuthorId, &author.Name)
		if err != nil {
			return nil, fmt.Errorf("Error scan row from db: %s", err)
		}

		authors = append(authors, &author)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("Row error: %s", rows.Err())
	}

	if len(authors) == 0 {
		return nil, fmt.Errorf("Empty result from query")
	}

	return authors, nil
}
