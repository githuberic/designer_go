package entity

// 产品实体类
type Product struct {
	ID int
	Name string
	Price float64
}

func NewProduct(id int,name string,price float64) *Product {
	return &Product{
		ID: id,
		Name: name,
		Price: price,
	}
}
