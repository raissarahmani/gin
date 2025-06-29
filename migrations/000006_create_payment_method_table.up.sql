-- public.payment_method definition

CREATE TABLE public.payment_method (
	id serial4 NOT NULL,
	payment_method varchar NULL,
	CONSTRAINT payment_method_pkey PRIMARY KEY (id)
);