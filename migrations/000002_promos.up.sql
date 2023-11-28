CREATE TABLE "go-lib".promos (
	id serial4 NOT NULL,
	promo_name varchar(100) NOT NULL,
	discount_type text NOT NULL,
	flat_amount int4 NOT NULL,
	percent_amount float8 NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	CONSTRAINT promos_discount_type_check CHECK ((discount_type = ANY (ARRAY['flat'::text, 'percent'::text]))),
	CONSTRAINT promos_percent_amount_check CHECK ((percent_amount < (1)::double precision)),
	CONSTRAINT promos_promo_name_key UNIQUE (promo_name)
);