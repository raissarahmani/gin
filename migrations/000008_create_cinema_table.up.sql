-- public.cinema definition

CREATE TABLE public.cinema (
	id serial4 NOT NULL,
	cinema_name varchar NOT NULL,
	price int4 NOT NULL,
	created_at timestamp DEFAULT now() NULL,
	updated_at timestamp DEFAULT now() NULL,
	CONSTRAINT cinema_pkey PRIMARY KEY (id)
);