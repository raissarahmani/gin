-- public.loyalty definition

CREATE TABLE public.loyalty (
	id serial4 NOT NULL,
	loyalty_title varchar NOT NULL,
	points int4 NULL,
	max_point int4 NOT NULL,
	created_at timestamp DEFAULT now() NULL,
	updated_at timestamp DEFAULT now() NULL,
	CONSTRAINT loyalty_pkey PRIMARY KEY (id)
);