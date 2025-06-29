-- public.schedule definition

CREATE TABLE public.schedule (
	id serial4 NOT NULL,
	book_date date NOT NULL,
	book_time time NOT NULL,
	created_at timestamp DEFAULT now() NULL,
	updated_at timestamp DEFAULT now() NULL,
	CONSTRAINT schedule_pkey PRIMARY KEY (id)
);