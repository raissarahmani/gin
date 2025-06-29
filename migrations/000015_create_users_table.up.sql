-- public.users definition

CREATE TABLE public.users (
	id serial4 NOT NULL,
	email varchar NOT NULL,
	"password" varchar NOT NULL,
	"role" varchar NULL,
	created_at timestamp DEFAULT now() NULL,
	updated_at timestamp DEFAULT now() NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);