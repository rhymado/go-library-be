package models

import (
	"time"
)

type BookModel struct {
	Id          int        `db:"id" json:"id,omitempty" valid:"-"`
	BookName    string     `db:"book_name" form:"bookName" json:"bookName" valid:"alphanum,required"`
	Price       int        `db:"price" form:"price" json:"price" valid:"required"`
	AuthorId    int        `db:"authors_id" form:"authorId" json:"authorId,omitempty" valid:"required"`
	PublisherId int        `db:"publishers_id" form:"publisherId" json:"publisherId,omitempty" valid:"required"`
	PromoId     int        `db:"promo_id" form:"promoId" json:"promoId,omitempty" valid:"optional"`
	CreatedAt   *time.Time `db:"created_at" json:"createdAt" valid:"-"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updatedAt" valid:"-"`
}

type BookResponseModel struct {
	BookModel
	AuthorName    string  `db:"author_name" json:"authorName" valid:"-"`
	PublisherName string  `db:"publisher_name" json:"publisherName" valid:"-"`
	PromoName     *string `db:"promo_name" json:"promoName" valid:"-"`
}

type BookEditModel struct {
	BookName string `form:"bookName" json:"bookName" valid:"alphanum,optional"`
	Price    int    `form:"price" json:"price" valid:"optional"`
	PromoId  int    `form:"promoId" json:"promoId,omitempty" valid:"optional"`
}
