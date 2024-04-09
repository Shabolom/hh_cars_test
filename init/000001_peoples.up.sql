CREATE TABLE IF NOT EXISTS public.peoples
(
    id uuid NOT NULL,
    crated_at TIMESTAMP WITH time zone,
    update_at TIMESTAMP WITH time zone,
    deleted_at TIMESTAMP WITH time zone,
    name text NOT NULL COLLATE  pg_catalog."default",
    patronymic text NOT NULL COLLATE pg_catalog."default",
    surname text NOT NULL COLLATE pg_catalog."default",
    CONSTRAINT peoples_unique UNIQUE (name, patronymic, surname),
    CONSTRAINT peoples_pk PRIMARY KEY (id)
    );

CREATE INDEX idx_peoples
    ON public.peoples (id);

ALTER TABLE IF EXISTS public.peoples
    OWNER to postgres;