package models

import (
	"time"
)

type BookModel struct {
	Id          int        `db:"id" json:"id,omitempty"`
	BookName    string     `db:"book_name" form:"bookName" json:"bookName"`
	Price       int        `db:"price" form:"price" json:"price"`
	AuthorId    int        `db:"authors_id" form:"authorId" json:"authorId,omitempty"`
	PublisherId int        `db:"publishers_id" form:"publisherId" json:"publisherId,omitempty"`
	PromoId     int        `db:"promo_id" form:"promoId" json:"promoId,omitempty"`
	CreatedAt   *time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updatedAt"`
	BookResponseModel
}

type BookResponseModel struct {
	AuthorName    string  `db:"author_name" json:"authorName"`
	PublisherName string  `db:"publisher_name" json:"publisherName"`
	PromoName     *string `db:"promo_name" json:"promoName"`
}
