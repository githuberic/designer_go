package product

import "time"

/**
拍品
*/
type IProduct interface {
	ID() int
	Name() string

	String() string
}

/**
品牌馆-拍品
*/
type IBrandProduct interface {
	IProduct
	Brand() string
}

/**
名家-拍品
*/
type IFamousAuthorProduct interface {
	IProduct
	Author() string
}

/**
拍卖行-拍品
*/
type IAuctionHouseProduct interface {
	IProduct
	GetInfo() *AuctionHouse
}

/**
拍卖行-拍品
*/
type AuctionHouse struct {
	// 拍卖行
	auctionHouse string
	// 拍卖截止日期
	endTime time.Time
}
