package database

var ProductList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red",
		Price:       300,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQsa3K1PkfEgVzc6JeayE-uGwejpsBDBbVBUw&s",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is green",
		Price:       200,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRJJfODaTyBw4581VyPy5wQHvq4yfAIzGRHVA&s",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana contain potasium",
		Price:       20,
		ImgURL:      "https://bazarstore.az/cdn/shop/products/30015539_767x1000.jpg?v=1693895477",
	}

	ProductList = append(ProductList, prd1)
	ProductList = append(ProductList, prd2)
	ProductList = append(ProductList, prd3)

}
