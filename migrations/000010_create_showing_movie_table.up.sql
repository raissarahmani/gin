-- public.showing_movie definition

CREATE TABLE public.showing_movie (
	id serial4 NOT NULL,
	movies_id int4 NOT NULL,
	cinema_id int4 NOT NULL,
	CONSTRAINT showing_movie_pkey PRIMARY KEY (movies_id, cinema_id)
);


-- public.showing_movie foreign keys

ALTER TABLE public.showing_movie ADD CONSTRAINT fk_cinema FOREIGN KEY (cinema_id) REFERENCES public.cinema(id);