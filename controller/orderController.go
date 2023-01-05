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
	// loc, err := time.LoadLocation("Asia/Seoul")
	// if err != nil {
	// 	protocol.Fail(err, protocol.BadRequest).Response(c)
	// 	return
	// }
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
// @Router /order/list [GET]
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
// @Router /order/modify [PATCH]
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

// AdminListOrderController godoc
// @Summary call AdminListOrderController, return ok by json.
// @가게에서 주문 목록 조회
// @name AdminListOrderController
// @Accept  json
// @Produce  json
// @Param id query string true "사업체 id"
// @Router /order/admin/list [GET]
// @Success 200 {object} Controller
func (p *Controller) AdminListOrderController(c *gin.Context) {
	id := c.Query("id")
	BID, res := p.adminListOrderInputValidate(id)
	if res != nil {
		res.Response(c)
	}
	result := p.md.AdminListOrder(BID)
	c.JSON(200, gin.H{"msg": "ok", "list": result})
}

func (p *Controller) adminListOrderInputValidate(id string) (primitive.ObjectID, *protocol.ApiResponse[any]) {
	BID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, protocol.Fail(err, protocol.BadRequest)
	}
	return BID, nil
}

// UpdateState godoc
// @Summary call UpdateState, return ok by json.
// @가게에서 주문 상태 변경
// @name UpdateState
// @Accept  json
// @Produce  json
// @Param input body UpdateStateInput true "주문 번호, 상태 "
// @Router /order/admin/update [PATCH]
// @Success 200 {object} Controller
func (p *Controller) UpdateState(c *gin.Context) {
	var input UpdateStateInput
	c.ShouldBind(&input)
	orderId, err := primitive.ObjectIDFromHex(input.OrderId)
	if err != nil {
		panic(err)
	}
	result := p.md.UpdateOrderState(orderId, input.State)
	if result {
		c.JSON(200, gin.H{"msg": "update request success"})
	} else {
		c.JSON(200, gin.H{"msg": "update request failed"})
	}
}

type UpdateStateInput struct {
	OrderId string `bson:"orderid:`
	State   int    `bson:"state"`
}

// CreateReview godoc
// @Summary call CreateReview, return ok by json.
// @리뷰 작성하기
// @name CreateReview
// @Accept  json
// @Produce  json
// @Param input body ReviewInput true "리뷰"
// @Router /review [POST]
// @Success 200 {object} Controller
func (p *Controller) CreateReview(c *gin.Context) {
	var input ReviewInput
	err := c.ShouldBind(&input)
	if err != nil {
		protocol.Fail(err, protocol.BadRequest).Response(c)
		return
	}
	review, res := p.createReviewInputValidate(input)
	if res != nil {
		res.Response(c)
		return
	}
	result := p.md.CreateReview(review)
	result.Response(c)
}

func (p *Controller) createReviewInputValidate(body ReviewInput) (*model.Review, *protocol.ApiResponse[any]) {
	review := new(model.Review)
	var err error
	review.OrderID, err = primitive.ObjectIDFromHex(body.OrderID)
	if err != nil {
		return nil, protocol.Fail(err, protocol.BadRequest)
	}
	if r, e := p.md.CheckOrderByID(review.OrderID); !r {
		return nil, protocol.FailCustomMessage(e, "No matching order", protocol.BadRequest)
	}
	review.MenuID, err = primitive.ObjectIDFromHex(body.MenuID)
	if err != nil {
		protocol.Fail(err, protocol.BadRequest)
	}

	review.Orderer = body.Orderer
	review.Content = body.Content
	review.Score = body.Score
	return review, nil
}

type ReviewInput struct {
	OrderID string  `bson:"order_id" json:"order_id"`
	MenuID  string  `bson:"menu_id" json:"menu_id"`
	Orderer string  `bson:"orderer"`
	Content string  `bson:"content"`
	Score   float32 `bson:"score"`
}
