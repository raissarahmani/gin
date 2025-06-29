-- public.showing definition

CREATE TABLE public.showing (
	id serial4 NOT NULL,
	movies_id int4 NOT NULL,
	cinema_id int4 NOT NULL,
	city_id int4 NOT NULL,
	CONSTRAINT showing_pkey PRIMARY KEY (cinema_id, movies_id, city_id)
);


-- public.showing foreign keys

ALTER TABLE public.showing ADD CONSTRAINT showing_city_id_cinema_id_fkey FOREIGN KEY (city_id,cinema_id) REFERENCES public.showing_city(city_id,cinema_id);
ALTER TABLE public.showing ADD CONSTRAINT showing_movies_id_cinema_id_fkey FOREIGN KEY (movies_id,cinema_id) REFERENCES public.showing_movie(movies_id,cinema_id);