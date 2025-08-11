-- public.categories definition

-- Drop table

-- DROP TABLE public.categories;

CREATE TABLE public.categories (
	id uuid NOT NULL,
	category_name varchar(255) NULL,
	created_by varchar(255) DEFAULT 'SYSTEM'::character varying NULL,
	updated_by varchar(255) DEFAULT 'SYSTEM'::character varying NULL,
	deleted_by varchar(255) NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	CONSTRAINT categories_category_name_key UNIQUE (category_name),
	CONSTRAINT categories_pkey PRIMARY KEY (id)
);

-- public.inventory_movement definition

-- Drop table

-- DROP TABLE public.inventory_movement;

CREATE TABLE public.inventory_movement (
	id uuid NOT NULL,
	movement_date timestamptz NULL,
	quantity_change int4 NULL,
	movement_type varchar(255) NULL,
	current_stock int4 NULL,
	created_by varchar(255) DEFAULT 'SYSTEM'::character varying NULL,
	updated_by varchar(255) DEFAULT 'SYSTEM'::character varying NULL,
	deleted_by varchar(255) NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	product_id uuid NULL,
	sales_id uuid NULL,
	CONSTRAINT inventory_movement_pkey PRIMARY KEY (id)
);

-- public.locations definition

-- Drop table

-- DROP TABLE public.locations;

CREATE TABLE public.locations (
	id uuid NOT NULL,
	city varchar(255) NULL,
	state varchar(255) NULL,
	postal_code varchar(255) NULL,
	region varchar(255) NULL,
	country varchar(255) NULL,
	created_by varchar(255) DEFAULT 'SYSTEM'::character varying NULL,
	updated_by varchar(255) DEFAULT 'SYSTEM'::character varying NULL,
	deleted_by varchar(255) NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	CONSTRAINT locations_pkey PRIMARY KEY (id),
	CONSTRAINT locations_postal_code_key UNIQUE (postal_code)
);

-- public.products definition

-- Drop table

-- DROP TABLE public.products;

CREATE TABLE public.products (
	id uuid NOT NULL,
	product_name varchar(255) NULL,
	manufacturer varchar(255) NULL,
	base_price float8 NULL,
	created_by varchar(255) DEFAULT 'SYSTEM'::character varying NULL,
	updated_by varchar(255) DEFAULT 'SYSTEM'::character varying NULL,
	deleted_by varchar(255) NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	category_id uuid NULL,
	sub_category_id uuid NULL,
	CONSTRAINT products_pkey PRIMARY KEY (id),
	CONSTRAINT products_product_name_key UNIQUE (product_name)
);


-- public.sales definition

-- Drop table

-- DROP TABLE public.sales;

CREATE TABLE public.sales (
	id uuid NOT NULL,
	ship_date timestamptz NULL,
	ship_mode varchar(255) NULL,
	customer_name varchar(255) NULL,
	quantity int2 NULL,
	sales_amount float8 NULL,
	discount float8 NULL,
	profit float8 NULL,
	profit_ratio float8 NULL,
	number_of_record int2 NULL,
	order_id varchar(255) NULL,
	order_date timestamptz NULL,
	created_by varchar(255) DEFAULT 'SYSTEM'::character varying NULL,
	updated_by varchar(255) DEFAULT 'SYSTEM'::character varying NULL,
	deleted_by varchar(255) NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	location_id uuid NULL,
	product_id uuid NULL,
	segment_id uuid NULL,
	CONSTRAINT sales_pkey PRIMARY KEY (id)
);


-- public.segments definition

-- Drop table

-- DROP TABLE public.segments;

CREATE TABLE public.segments (
	id uuid NOT NULL,
	segment_name varchar(255) NULL,
	created_by varchar(255) DEFAULT 'SYSTEM'::character varying NULL,
	updated_by varchar(255) DEFAULT 'SYSTEM'::character varying NULL,
	deleted_by varchar(255) NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	CONSTRAINT segments_pkey PRIMARY KEY (id),
	CONSTRAINT segments_segment_name_key UNIQUE (segment_name)
);

-- public.sub_categories definition

-- Drop table

-- DROP TABLE public.sub_categories;

CREATE TABLE public.sub_categories (
	id uuid NOT NULL,
	sub_category_name varchar(255) NULL,
	created_by varchar(255) DEFAULT 'SYSTEM'::character varying NULL,
	updated_by varchar(255) DEFAULT 'SYSTEM'::character varying NULL,
	deleted_by varchar(255) NULL,
	created_at timestamptz NOT NULL,
	updated_at timestamptz NOT NULL,
	category_id uuid NULL,
	CONSTRAINT sub_categories_pkey PRIMARY KEY (id),
	CONSTRAINT sub_categories_sub_category_name_key UNIQUE (sub_category_name)
);

-- public.products foreign keys

ALTER TABLE public.products ADD CONSTRAINT products_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id) ON DELETE SET NULL ON UPDATE CASCADE;
ALTER TABLE public.products ADD CONSTRAINT products_category_id_fkey1 FOREIGN KEY (category_id) REFERENCES public.categories(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE public.products ADD CONSTRAINT products_sub_category_id_fkey FOREIGN KEY (sub_category_id) REFERENCES public.sub_categories(id) ON DELETE SET NULL ON UPDATE CASCADE;

-- public.sales foreign keys

ALTER TABLE public.sales ADD CONSTRAINT sales_location_id_fkey FOREIGN KEY (location_id) REFERENCES public.locations(id) ON DELETE SET NULL ON UPDATE CASCADE;
ALTER TABLE public.sales ADD CONSTRAINT sales_location_id_fkey1 FOREIGN KEY (location_id) REFERENCES public.locations(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE public.sales ADD CONSTRAINT sales_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE SET NULL ON UPDATE CASCADE;
ALTER TABLE public.sales ADD CONSTRAINT sales_product_id_fkey1 FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE public.sales ADD CONSTRAINT sales_segment_id_fkey FOREIGN KEY (segment_id) REFERENCES public.segments(id) ON DELETE SET NULL ON UPDATE CASCADE;

-- public.inventory_movement foreign keys

ALTER TABLE public.inventory_movement ADD CONSTRAINT inventory_movement_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE SET NULL ON UPDATE CASCADE;
ALTER TABLE public.inventory_movement ADD CONSTRAINT inventory_movement_sales_id_fkey FOREIGN KEY (sales_id) REFERENCES public.sales(id) ON DELETE SET NULL ON UPDATE CASCADE;

-- public.sub_categories foreign keys

ALTER TABLE public.sub_categories ADD CONSTRAINT sub_categories_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id) ON DELETE SET NULL ON UPDATE CASCADE;
ALTER TABLE public.sub_categories ADD CONSTRAINT sub_categories_category_id_fkey1 FOREIGN KEY (category_id) REFERENCES public.categories(id) ON DELETE CASCADE ON UPDATE CASCADE;