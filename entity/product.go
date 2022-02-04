package entity

type Product struct {
	Id           int    `json:"id" gorm:"primary_key"`
	ProductName  string `json:"productname"`
	ProductPrice int    `json:"productprice"`
	ShopId       int    `json:"shopid"`
}
