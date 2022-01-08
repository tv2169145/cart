package model

type Cart struct {
	ID        int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	UserId    int64 `gorm:"not_null" json:"user_id"`
	ProductId int64 `gorm:"not_null" json:"product_id"`
	SizeId    int64 `gorm:"not_null" json:"size_id"`
	Num       int64 `gorm:"not_null" json:"num"`
}
