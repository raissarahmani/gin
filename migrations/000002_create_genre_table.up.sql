-- public.genre definition

CREATE TABLE public.genre (
	id serial4 NOT NULL,
	genre_name varchar NOT NULL,
	CONSTRAINT genre_pkey PRIMARY KEY (id)
);