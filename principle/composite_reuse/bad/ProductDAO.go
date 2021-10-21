package bad

import (
	"designer_go/principle/composite_reuse/entity"
)

/**
直接从BadDBConnection继承, 以获取访问数据库的能力. 继承仅仅是为了代码复用, 而不是概念复用, 因此违反了合成复用原则
*/
type ProductDAO struct {
	DBConnection
}

func NewProductDAO(url string, uid string, pwd string) *ProductDAO {
	return &ProductDAO{*NewDBConnection(url, uid, pwd)}
}

func (p *ProductDAO) Insert(it *entity.Product) (error, int) {
	return p.Execute("insert into product(id,name,price) values(?, ?, ?)", it.ID, it.Name, it.Price)
}

func (p *ProductDAO) Update(it *entity.Product) (error, int) {
	return p.Execute("update product set name=? price=? where id=?", it.Name, it.Price, it.ID)
}

func (p *ProductDAO) Delete(id int) (error, int) {
	return p.Execute("delete from product where id=?", id)
}
