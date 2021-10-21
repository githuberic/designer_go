package good

import (
	"designer_go/principle/composite_reuse/entity"
)

/**
更好的设计, 通过Setter方法注入数据库方言实例(遵循了合成复用原则), 实现产品的CRUD
*/
type ProductDAO struct {
	dbConn IDBConnection
}

func NewProductDAO() *ProductDAO {
	return &ProductDAO{}
}

func (p *ProductDAO) SetDBConnection(dbConn IDBConnection) {
	p.dbConn = dbConn
}

func (p *ProductDAO) Insert(product *entity.Product) (error, int) {
	return p.dbConn.Execute("insert into product(id,name,price) values(?, ?, ?)", product.ID, product.Name, product.Price)
}

func (p *ProductDAO) Update(it *entity.Product) (error, int) {
	return p.dbConn.Execute("update product set name=? price=? where id=?", it.Name, it.Price, it.ID)
}

func (p *ProductDAO) Delete(id int) (error, int) {
	return p.dbConn.Execute("delete from product where id=?", id)
}
