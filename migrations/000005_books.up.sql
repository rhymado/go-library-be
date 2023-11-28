CREATE TABLE "go-lib".books (
	id serial4 NOT NULL,
	book_name varchar(255) NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	price int4 NOT NULL,
	authors_id int4 NOT NULL,
	publishers_id int4 NOT NULL DEFAULT 0,
	promo_id int4 NULL,
	CONSTRAINT pk PRIMARY KEY (id)
);

ALTER TABLE "go-lib".books ADD CONSTRAINT fk_books_authors FOREIGN KEY (authors_id) REFERENCES "go-lib".authors(id);
ALTER TABLE "go-lib".books ADD CONSTRAINT fk_books_publishers FOREIGN KEY (publishers_id) REFERENCES "go-lib".publishers(id);