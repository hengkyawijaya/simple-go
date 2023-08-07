CREATE TABLE public.users (
    user_id SERIAL PRIMARY KEY,
    name varchar(255),
    email varchar(255),
    password varchar(255),
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp NULL DEFAULT NULL
);