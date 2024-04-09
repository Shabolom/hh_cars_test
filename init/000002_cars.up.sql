CREATE TABLE IF NOT EXISTS public.cars
(
    id  uuid NOT NULL,
    created_at TIMESTAMP WITH time zone,
    updated_at TIMESTAMP WITH time zone,
    deleted_at TIMESTAMP WITH time zone,
    owner_id uuid NOT NULL UNIQUE REFERENCES public.peoples(id),
    mark text NOT NULL COLLATE pg_catalog."default",
    model text NOT NULL COLLATE pg_catalog."default",
    reg_num text NOT NULL COLLATE pg_catalog."default",
    year integer NOT NULL,
    CONSTRAINT cars_unique UNIQUE (mark, model, owner_id, reg_num),
    CONSTRAINT cars_pk PRIMARY KEY (id)
    );

CREATE INDEX idx_cars_id
    ON public.cars (id);

ALTER TABLE IF EXISTS public.cars
    OWNER to postgres;