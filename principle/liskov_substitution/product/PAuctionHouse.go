package product

import "time"

type PAuctionHouse struct {
	*Product

	// 拍卖行
	auctionHouse string
	// 拍卖截止日期
	endTime time.Time
}

func NewPAuctionHouse(id int, name string, price float32,
	auctionHouse string, endTime time.Time) IProduct {
	return &PAuctionHouse{
		&Product{iID: id, sName: name, price: price},
		auctionHouse,
		endTime,
	}
}

func NewPAuctionHouseV2(product *Product, auctionHouse string, endTime time.Time) IProduct {
	return &PAuctionHouse{
		Product:      product,
		auctionHouse: auctionHouse,
		endTime:      endTime,
	}
}

func (p *PAuctionHouse) GetInfo() *AuctionHouse {
	return &AuctionHouse{auctionHouse: p.auctionHouse, endTime: p.endTime}
}
