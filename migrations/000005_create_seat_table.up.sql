-- public.seat definition

CREATE TABLE public.seat (
	id serial4 NOT NULL,
	seat_number varchar NULL,
	CONSTRAINT seat_pkey PRIMARY KEY (id)
);