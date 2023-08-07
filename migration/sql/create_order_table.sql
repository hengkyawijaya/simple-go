CREATE TABLE public.orders (
    order_id SERIAL PRIMARY KEY,
    user_id int,
    product_id int,
    price float,
    status varchar(255),
    rc varchar(255),
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp NULL DEFAULT NULL
);