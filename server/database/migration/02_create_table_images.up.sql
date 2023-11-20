CREATE TABLE IF NOT EXISTS tb_images (
  id SERIAL PRIMARY KEY,
  product_id VARCHAR(36) NOT NULL,
  url TEXT UNIQUE NOT NULL,
  CONSTRAINT fk_product_id 
    FOREIGN KEY (product_id)
      REFERENCES tb_products (id)
);