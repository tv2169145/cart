package service

import (
	"github.com/tv2169145/cart/domain/model"
	"github.com/tv2169145/cart/domain/repository"
)

type ICartDataService interface {
	AddCart(*model.Cart) (int64, error)
	DeleteCart(int64) error
	UpdateCart(*model.Cart) error
	FindCartByID(int64) (*model.Cart, error)
	FindAllCart(int64) ([]model.Cart, error)

	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

type CartDataService struct {
	CartRepository repository.ICartRepository
}

func NewCartDataService(cartRepository repository.ICartRepository) ICartDataService {
	return &CartDataService{CartRepository: cartRepository}
}

func (service *CartDataService) AddCart(cart *model.Cart) (int64, error) {
	return service.CartRepository.CreateCart(cart)
}

func (service *CartDataService) DeleteCart(cartId int64) error {
	return service.CartRepository.DeleteCartByID(cartId)
}

func (service *CartDataService) UpdateCart(cart *model.Cart) error {
	return service.CartRepository.UpdateCart(cart)
}

func (service *CartDataService) FindCartByID(cartId int64) (*model.Cart, error) {
	return service.CartRepository.FindCartByID(cartId)
}

func (service *CartDataService) FindAllCart(userId int64) ([]model.Cart, error) {
	return service.CartRepository.FindAll(userId)
}

func (service *CartDataService) CleanCart(userId int64) error {
	return service.CartRepository.CleanCart(userId)
}

func (service *CartDataService) IncrNum(cartId int64, num int64) error {
	return service.CartRepository.IncrNum(cartId, num)
}

func (service *CartDataService) DecrNum(cartId int64, num int64) error {
	return service.CartRepository.DecrNum(cartId, num)
}
