-- public.profile definition

CREATE TABLE public.profile (
	users_id int4 NULL,
	loyalty_id int4 NULL,
	profile_image varchar NOT NULL,
	first_name varchar NOT NULL,
	last_name varchar NOT NULL,
	phone varchar NOT NULL,
	email varchar NOT NULL,
	created_at timestamp DEFAULT now() NULL,
	updated_at timestamp DEFAULT now() NULL
);


-- public.profile foreign keys

ALTER TABLE public.profile ADD CONSTRAINT fk_loyalty FOREIGN KEY (loyalty_id) REFERENCES public.loyalty(id);
ALTER TABLE public.profile ADD CONSTRAINT fk_users FOREIGN KEY (users_id) REFERENCES public.users(id);