CREATE TABLE IF NOT EXISTS realtime_order (
    `id` String,
    `sale_id` String,
    `quantity` UInt8,
    `sale_amount` Float32,
    `discount` Float32,
    `profit` Float32,
    `profit_ratio` Float32,
    `order_id` String,
    `order_date` DateTime,
    `location_id` String,
    `product_id` String,
    `segment_id` String,
    `product_name` String,
    `segment_name` String,
    `created_at` DateTime
)
ENGINE = MergeTree
PARTITION BY product_id
ORDER BY id;