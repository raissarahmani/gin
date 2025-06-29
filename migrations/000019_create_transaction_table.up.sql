-- public."transaction" definition

CREATE TABLE public."transaction" (
	id serial4 NOT NULL,
	users_id int4 NULL,
	fullname varchar NOT NULL,
	email varchar NOT NULL,
	phone varchar NOT NULL,
	total_price int4 NOT NULL,
	payment_method_id int4 NOT NULL,
	payment_done bool NULL,
	created_at timestamp DEFAULT now() NULL,
	updated_at timestamp DEFAULT now() NULL,
	CONSTRAINT transaction_pkey PRIMARY KEY (id)
);


-- public."transaction" foreign keys

ALTER TABLE public."transaction" ADD CONSTRAINT fk_payment_method FOREIGN KEY (payment_method_id) REFERENCES public.payment_method(id);
ALTER TABLE public."transaction" ADD CONSTRAINT fk_users FOREIGN KEY (users_id) REFERENCES public.users(id);