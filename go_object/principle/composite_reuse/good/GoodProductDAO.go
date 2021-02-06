package good

import (
	"tortoise/go_object/principle/composite_reuse/entity"
)

type GoodProductDAO struct {
	dbConn IGoodDBConnection
}

func NewProductDAO() *GoodProductDAO {
	return &GoodProductDAO{}
}

func (p *GoodProductDAO) SetDBConnection(dbConn IGoodDBConnection) {
	p.dbConn = dbConn
}

func (p *GoodProductDAO) Insert(product *entity.Product) (error, int) {
	return p.dbConn.Execute("insert into product(id,name,price) values(?, ?, ?)", product.ID, product.Name, product.Price)
}

func (p *GoodProductDAO) Update(it *entity.Product) (error, int) {
	return p.dbConn.Execute("update product set name=? price=? where id=?", it.Name, it.Price, it.ID)
}

func (p *GoodProductDAO) Delete(id int) (error, int) {
	return p.dbConn.Execute("delete from product where id=?", id)
}