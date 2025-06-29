-- public.showing_schedule definition

CREATE TABLE public.showing_schedule (
	schedule_id int4 NOT NULL,
	movies_id int4 NOT NULL,
	cinema_id int4 NOT NULL,
	city_id int4 NOT NULL,
	id serial4 NOT NULL,
	CONSTRAINT showing_schedule_pkey PRIMARY KEY (schedule_id, city_id, cinema_id, movies_id)
);


-- public.showing_schedule foreign keys

ALTER TABLE public.showing_schedule ADD CONSTRAINT showing_schedule_city_id_cinema_id_movies_id_fkey FOREIGN KEY (city_id,cinema_id,movies_id) REFERENCES <?>();
ALTER TABLE public.showing_schedule ADD CONSTRAINT showing_schedule_schedule_id_fkey FOREIGN KEY (schedule_id) REFERENCES public.schedule(id);