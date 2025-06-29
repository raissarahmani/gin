-- public.movies_image definition

CREATE TABLE public.movies_image (
	id serial4 NOT NULL,
	poster varchar NULL,
	CONSTRAINT movies_image_pkey PRIMARY KEY (id)
);