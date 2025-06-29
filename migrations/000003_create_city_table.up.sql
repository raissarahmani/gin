-- public.city definition

CREATE TABLE public.city (
	id serial4 NOT NULL,
	city varchar NOT NULL,
	CONSTRAINT city_pkey PRIMARY KEY (id)
);