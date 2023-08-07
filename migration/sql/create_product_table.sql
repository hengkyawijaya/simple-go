CREATE TABLE public.products (
    product_id SERIAL PRIMARY KEY,
    name varchar(255),
    price float,
    status varchar(255),
    code varchar(255),
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp NULL DEFAULT NULL
);