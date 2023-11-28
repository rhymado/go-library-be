CREATE TABLE "go-lib".publishers (
	id serial4 NOT NULL,
	publisher_name varchar(100) NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	CONSTRAINT pk_publishers PRIMARY KEY (id)
);