CREATE TABLE "go-lib".authors (
	id serial4 NOT NULL,
	author_name varchar(100) NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	CONSTRAINT pk_authors PRIMARY KEY (id)
);