package product

const (
	CREATE_QUERY = `INSERT INTO tb_products (id, name, brand, description, price, stock, created_at) VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP)`
)
