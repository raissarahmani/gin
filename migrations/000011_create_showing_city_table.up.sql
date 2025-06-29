-- public.showing_city definition

CREATE TABLE public.showing_city (
	id serial4 NOT NULL,
	cinema_id int4 NOT NULL,
	city_id int4 NOT NULL,
	CONSTRAINT showing_city_pkey PRIMARY KEY (city_id, cinema_id)
);


-- public.showing_city foreign keys

ALTER TABLE public.showing_city ADD CONSTRAINT fk_cinema FOREIGN KEY (cinema_id) REFERENCES public.cinema(id);
ALTER TABLE public.showing_city ADD CONSTRAINT fk_city FOREIGN KEY (city_id) REFERENCES public.city(id);