-- public.movies_genre definition

CREATE TABLE public.movies_genre (
	movies_id int4 NOT NULL,
	genre_id int4 NOT NULL,
	CONSTRAINT movies_genre_pkey PRIMARY KEY (movies_id, genre_id)
);


-- public.movies_genre foreign keys

ALTER TABLE public.movies_genre ADD CONSTRAINT genre FOREIGN KEY (genre_id) REFERENCES public.genre(id);
ALTER TABLE public.movies_genre ADD CONSTRAINT movies FOREIGN KEY (movies_id) REFERENCES public.movies(id);