-- public.movies definition

CREATE TABLE public.movies (
	id serial4 NOT NULL,
	title varchar NOT NULL,
	movies_image_id int4 NULL,
	release_date date NOT NULL,
	duration int4 NOT NULL,
	director varchar NOT NULL,
	casts _varchar NOT NULL,
	synopsis text NOT NULL,
	created_at timestamp DEFAULT now() NULL,
	updated_at timestamp DEFAULT now() NULL,
	CONSTRAINT movies_pkey PRIMARY KEY (id)
);


-- public.movies foreign keys

ALTER TABLE public.movies ADD CONSTRAINT fk_movies_image FOREIGN KEY (movies_image_id) REFERENCES public.movies_image(id);