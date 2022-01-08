package handler

import (
	"context"
	"github.com/tv2169145/cart/domain/model"
	"github.com/tv2169145/cart/domain/service"
	cart "github.com/tv2169145/cart/proto/cart"
	"github.com/tv2169145/common"
)

type Cart struct {
	CartDateService service.ICartDataService
}

// 添加購物車
func (cartHandler *Cart) AddCart(ctx context.Context, request *cart.CartInfo, response *cart.ResponseAdd) error {
	newCart := &model.Cart{}
	err := common.SwapTo(request, newCart)
	if err != nil {
		return err
	}
	cartId, err := cartHandler.CartDateService.AddCart(newCart)
	if err != nil {
		return err
	}
	response.CartId = cartId
	response.Msg = "success"

	return nil
}

// 清空購物車
func (cartHandler *Cart) CleanCart(ctx context.Context, request *cart.Clean, response *cart.Response) error {
	err := cartHandler.CartDateService.CleanCart(request.UserId)
	if err != nil {
		return err
	}
	response.Msg = "success"

	return nil
}

// 增加商品數量
func (cartHandler *Cart) Incr(ctx context.Context, request *cart.Item, response *cart.Response) error {
	err := cartHandler.CartDateService.IncrNum(request.Id, request.ChangeNum)
	if err != nil {
		return err
	}

	response.Msg = "success"
	return nil
}

// 減少商品數量
func (cartHandler *Cart) Decr(ctx context.Context, request *cart.Item, response *cart.Response) error {
	err := cartHandler.CartDateService.DecrNum(request.Id, request.ChangeNum)
	if err != nil {
		return err
	}
	response.Msg = "success"

	return nil
}

// 刪除指定item
func (cartHandler *Cart) DeleteItemByID(ctx context.Context, request *cart.CartId, response *cart.Response) error {
	err := cartHandler.CartDateService.DeleteCart(request.Id)
	if err != nil {
		return err
	}
	response.Msg = "success"

	return nil
}

// 取得某user的所有購物車項目
func (cartHandler *Cart) GetAll(ctx context.Context, request *cart.CartFindAll, response *cart.CartAll) error {
	carts, err := cartHandler.CartDateService.FindAllCart(request.UserId)
	if err != nil {
		return err
	}

	for _, cartModel := range carts {
		cartInfo := &cart.CartInfo{}
		if err := common.SwapTo(cartModel, cartInfo); err != nil {
			return err
		}
		response.CartInfo = append(response.CartInfo, cartInfo)
	}

	return nil
}
