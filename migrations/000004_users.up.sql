CREATE TABLE "go-lib".users (
	id serial4 NOT NULL,
	username varchar(100) NOT NULL,
	email text NOT NULL,
	userpass varchar(100) NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	user_role varchar(10) NOT NULL DEFAULT 'user'::character varying,
	image varchar NULL,
	CONSTRAINT users_email_key UNIQUE (email),
	CONSTRAINT users_pk PRIMARY KEY (id),
	CONSTRAINT users_user_role_check CHECK (((user_role)::text = ANY ((ARRAY['user'::character varying, 'admin'::character varying])::text[]))),
	CONSTRAINT users_username_key UNIQUE (username)
);