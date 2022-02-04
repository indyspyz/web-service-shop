package entity

type Shop struct {
	Id       int    `json:"id" gorm:"primary_key"`
	ShopName string `json:"shopname"`
}

type ShopList struct {
	Id          int       `json:"id" gorm:"primary_key"`
	ShopName    string    `json:"shopname"`
	ProductList []Product `gorm:"foreignKey:ShopId;references:Id"`
}
