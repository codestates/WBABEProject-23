package controller

import (
	"fmt"
	"lecture/WBABEProject-23/model"
	"lecture/WBABEProject-23/protocol"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateOrder godoc
// @Summary call CreateOrder, return ok by json.
// @주문.
// @name CreateOrder
// @Accept  json
// @Produce  json
// @Param input body CreateOrderInput true "주문자 이름,  메뉴 배열형태로 메뉴ID, 주문 수량 입력"
// @Router /order/make [POST]
// @Success 200 {object} Controller
func (p *Controller) CreateOrder(c *gin.Context) {
	var input = new(CreateOrderInput)
	if err := c.ShouldBind(&input); err != nil {
		protocol.Fail(err, protocol.BadRequest).Response(c)
		return
	}
	order, res := p.createOrderInputValidate(input)
	if res != nil {
		res.Response(c)
		return
	}
	result := p.md.CreateOrder(order)
	result.Response(c)
}

func (p *Controller) createOrderInputValidate(body *CreateOrderInput) (*model.Order, *protocol.ApiResponse[any]) {
	var order = new(model.Order)

	for i, menu := range body.Menu {
		t, err := primitive.ObjectIDFromHex(menu.MenuID)
		order.Menu = append(order.Menu, model.MenuNum{t, menu.Number, false})
		if err != nil {
			return nil, protocol.Fail(err, protocol.BadRequest)
		}
		if r, e := p.md.CheckMenuID(order.Menu[i].MenuID); !r {
			msg := fmt.Sprintf("No menu id with %v\n", menu.MenuID)
			return nil, protocol.FailCustomMessage(e, msg, protocol.BadRequest)
		}
	}
	var err error
	order.BID, err = primitive.ObjectIDFromHex(body.BID)
	if err != nil {
		return nil, protocol.Fail(err, protocol.BadRequest)
	}
	if r, e := p.md.CheckBusinessID(order.BID); !r {
		msg := fmt.Sprintf("No menu id with %v\n", body.BID)
		return nil, protocol.FailCustomMessage(e, msg, protocol.BadRequest)
	}
	order.Orderer = body.Orderer
	order.ID = primitive.NewObjectID()
	order.CreatedAt = time.Now()
	order.State = model.Receipting
	return order, nil
}

type CreateOrderInput struct {
	Orderer string `bson:"orderer"`
	BID     string `bson:"business_id"`
	Menu    []struct {
		MenuID string `bson:"menu_id"`
		Number int    `bson:"number"`
	} `bson:"menu"`
}

// ListOrder godoc
// @Summary call ListOrder, return ok by json.
// @주문자 - 주문조회 서비스
// @name ListOrder
// @Accept  json
// @Produce  json
// @Param name query string true "유저이름"
// @Param cur query string true "1은 현재 주문, 그 외 과거 주문"
// @Router /order [GET]
// @Success 200 {object} Controller
func (p *Controller) ListOrder(c *gin.Context) {
	userName := c.Query("name")
	cur := c.Query("cur")

	result := p.md.ListOrder(userName, cur == "1")

	c.JSON(200, gin.H{"msg": "ok", "list": result})
}

// UpdateOrder godoc
// @Summary call UpdateOrder, return ok by json.
// @주문자 - 주문 변경 서비스
// @name UpdateOrder
// @Accept  json
// @Produce  json
// @Param input body UpdateOrderInput true "수정할 주문 번호, 변경한 주문 메뉴 [{메뉴이름, 수량}]"
// @Router /order [PATCH]
// @Success 200 {object} Controller
func (p *Controller) UpdateOrder(c *gin.Context) {
	var input UpdateOrderInput
	if err := c.ShouldBind(&input); err != nil {
		protocol.Fail(err, protocol.BadRequest)
		return
	}
	objID, menu, res := p.updateOrderInputValidate(input)
	if res != nil {
		res.Response(c)
		return
	}
	result := p.md.UpdateOrder(objID, menu)
	result.Response(c)
}

func (p *Controller) updateOrderInputValidate(input UpdateOrderInput) (orderID primitive.ObjectID, menu []model.MenuNum, res *protocol.ApiResponse[any]) {
	orderID, err := primitive.ObjectIDFromHex(input.OrderID)
	if err != nil {
		return primitive.NilObjectID, nil, protocol.Fail(err, protocol.BadRequest)
	}
	for _, m := range input.Menu {
		id, err := primitive.ObjectIDFromHex(m.MenuID)
		if err != nil {
			return primitive.NilObjectID, nil, protocol.Fail(err, protocol.BadRequest)
		}
		menu = append(menu, model.MenuNum{
			MenuID:     id,
			Number:     m.Number,
			IsReviewed: false,
		})
	}
	return orderID, menu, nil
}

type UpdateOrderInput struct {
	OrderID string `bson:"orderid"`
	Menu    []struct {
		MenuID string `bson:"menu_id"`
		Number int    `bson:"number"`
	} `bson:"menu"`
}
