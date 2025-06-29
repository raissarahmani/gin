-- public.showing_seat_schedule definition

CREATE TABLE public.showing_seat_schedule (
	id serial4 NOT NULL,
	seat_id int4 NOT NULL,
	schedule_id int4 NOT NULL,
	city_id int4 NOT NULL,
	cinema_id int4 NOT NULL,
	movies_id int4 NOT NULL,
	is_available bool DEFAULT true NULL,
	CONSTRAINT showing_seat_schedule_pkey PRIMARY KEY (schedule_id, city_id, cinema_id, movies_id, seat_id)
);


-- public.showing_seat_schedule foreign keys

ALTER TABLE public.showing_seat_schedule ADD CONSTRAINT showing_seat_schedule_schedule_id_city_id_cinema_id_movies_fkey FOREIGN KEY (schedule_id,city_id,cinema_id,movies_id) REFERENCES public.showing_schedule(schedule_id,city_id,cinema_id,movies_id);
ALTER TABLE public.showing_seat_schedule ADD CONSTRAINT showing_seat_schedule_seat_id_fkey FOREIGN KEY (seat_id) REFERENCES public.seat(id);