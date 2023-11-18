package product_test

import (
	"testing"

	"github.com/ogabrielrodrigues/go-shop/server/internal/model/product"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	new_product := product.NewProduct(
		"Razer DeathAdder Elite Gaming Mouse",
		"Razer",
		"Equipped with the new eSports-grade optical sensor that has true 16,000 DPI and true tracking at 450 Inches Per Second (IPS), the Razer DeathAdder Elite gives you the absolute advantage of having the fastest sensor in the world. Engineered to redefine the standards of accuracy and speed, this incredible mouse sensor crushes the competition with a Resolution Accuracy of 99.4 percent, so you can land more killing blows with pinpoint precision. 1000 Herz Ultrapolling.",
		41.98,
		98,
	)

	assert.NotNil(t, new_product, "the new product must not be null")
}
