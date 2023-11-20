package product

const (
	CREATE_QUERY        = `INSERT INTO tb_products (id, name, brand, description, price, stock, created_at) VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP)`
	CREATE_IMAGES_QUERY = `INSERT INTO tb_images (product_id, url) VALUES ($1, $2)`
	// FIND_ALL_QUERY                   = `SELECT * FROM tb_products`
	FIND_ALL_QUERY                   = `SELECT tb_products.*, array_agg(tb_images.url) AS images FROM tb_products JOIN tb_images ON tb_products.id = tb_images.product_id GROUP BY tb_products.id`
	FIND_ALL_IMAGES_OF_PRODUCT_QUERY = `SELECT url FROM tb_images WHERE product_id = $1`
)
