package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/tv2169145/cart/domain/model"
)

type ICartRepository interface {
	InitTable() error
	FindCartByID(int64) (*model.Cart, error)
	CreateCart(*model.Cart) (int64, error)
	DeleteCartByID(int64) error
	UpdateCart(*model.Cart) error
	FindAll(int64) ([]model.Cart, error)

	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

type CartRepository struct {
	mysqlDb *gorm.DB
}

func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{mysqlDb: db}
}

func (repo *CartRepository) InitTable() error {
	return repo.mysqlDb.CreateTable(&model.Cart{}).Error
}

func (repo *CartRepository) FindCartByID(cartID int64) (cart *model.Cart, err error) {
	cart = &model.Cart{}
	return cart, repo.mysqlDb.First(cart, cartID).Error
}

func (repo *CartRepository) CreateCart(cart *model.Cart) (int64, error) {
	db := repo.mysqlDb.FirstOrCreate(cart, model.Cart{
		ProductId: cart.ProductId,
		SizeId:    cart.SizeId,
		UserId:    cart.UserId,
	})

	if err := db.Error; err != nil {
		return 0, err
	}

	if db.RowsAffected == 0 {
		return 0, errors.New("新增購物車失敗")
	}

	return cart.ID, nil
}

func (repo *CartRepository) DeleteCartByID(cartId int64) error {
	return repo.mysqlDb.Where("id = ?", cartId).Delete(&model.Cart{}).Error
}

func (repo *CartRepository) UpdateCart(cart *model.Cart) error {
	return repo.mysqlDb.Model(cart).Update(cart).Error
}

func (repo *CartRepository) FindAll(userId int64) (cartAll []model.Cart, err error) {
	return cartAll, repo.mysqlDb.Where("user_id=?", userId).Find(cartAll).Error
}

func (repo *CartRepository) CleanCart(userId int64) error {
	return repo.mysqlDb.Where("user_id = ?", userId).Delete(&model.Cart{}).Error
}

// 增加
func (repo *CartRepository) IncrNum(cartId int64, num int64) error {
	cart := &model.Cart{ID: cartId}
	return repo.mysqlDb.Model(cart).
		UpdateColumn("num", gorm.Expr("num + ?", num)).
		Error
}

// 減少
func (repo *CartRepository) DecrNum(cartId int64, num int64) error {
	cart := &model.Cart{ID: cartId}
	db := repo.mysqlDb.Model(cart).
		Where("num >= ?", num).
		UpdateColumn("num", gorm.Expr("num - ?", num))

	if err := db.Error; err != nil {
		return err
	}

	if db.RowsAffected == 0 {
		return errors.New("減少失敗")
	}

	return nil
}
