-- public.seat_selection definition

CREATE TABLE public.seat_selection (
	transaction_id int4 NOT NULL,
	schedule_id int4 NOT NULL,
	city_id int4 NOT NULL,
	cinema_id int4 NOT NULL,
	movies_id int4 NOT NULL,
	seat_id int4 NOT NULL,
	CONSTRAINT seat_selection_pkey PRIMARY KEY (transaction_id, schedule_id, city_id, cinema_id, movies_id, seat_id)
);


-- public.seat_selection foreign keys

ALTER TABLE public.seat_selection ADD CONSTRAINT fk_transaction FOREIGN KEY (transaction_id) REFERENCES public."transaction"(id);
ALTER TABLE public.seat_selection ADD CONSTRAINT seat_selection_schedule_id_city_id_cinema_id_movies_id_sea_fkey FOREIGN KEY (schedule_id,city_id,cinema_id,movies_id,seat_id) REFERENCES public.showing_seat_schedule(schedule_id,city_id,cinema_id,movies_id,seat_id);