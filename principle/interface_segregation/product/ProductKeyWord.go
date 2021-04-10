package product

import "time"

type PKeyWorld struct {
	*Product
}

func NewPKeyWorld(product *Product) IProduct {
	return &PKeyWorld{
		Product: product,
	}
}

func (p *PKeyWorld) Keyword() string {
	var keyword = p.sName

	keyword +=",文玩"

	// 调用搜索...
	// keyword=""

	// 520大促
	now := time.Now()
	var timestart = "2021-05-20 00:00:00"
	t1, err := time.Parse("2021-05-20 00:00:00", timestart)
	t2, err := time.Parse("2021-05-20 23:59:59", timestart)
	if err == nil && now.After(t1) && now.Before(t2) {
		keyword += "520促销"
	}

	// 11.9大促
	t3, err := time.Parse("2021-11-09 00:00:00", timestart)
	t4, err := time.Parse("2021-11-09 23:59:59", timestart)
	if err == nil && now.After(t3) && now.Before(t4) {
		keyword += "11大促促销"
	}

	// 直播处理...


	return keyword
}
