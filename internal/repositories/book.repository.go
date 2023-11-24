package repositories

import (
	"lib/internal/models"

	"github.com/jmoiron/sqlx"
)

type BookRepository struct {
	*sqlx.DB
}

func InitBookRepository(db *sqlx.DB) *BookRepository {
	return &BookRepository{db}
}

func (b *BookRepository) ReadAllBooks(page int, limit int) ([]models.BookModel, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 5
	}
	query := `select b.id, b.book_name, 
	a.author_name, pb.publisher_name, b.price, pr.promo_name, b.created_at, b.updated_at
	from books b
	join authors a on b.authors_id = a.id
	join publishers pb on b.publishers_id = pb.id
	left join promos pr on b.promo_id = pr.id
	limit $1 offset $2`
	offset := (page - 1) * limit
	values := []any{limit, offset}

	result := []models.BookModel{}
	err := b.Select(&result, query, values...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *BookRepository) CreateNewBook(body *models.BookModel) error {
	query := "insert into books (book_name, authors_id, publishers_id, price, promo_id) values ($1, $2, $3, $4, $5)"

	values := []any{body.BookName, body.AuthorId, body.PublisherId, body.Price, body.PromoId}

	if _, err := b.Exec(query, values...); err != nil {
		return err
	}
	return nil
}
