package controller

import (
	"lecture/WBABEProject-23/protocol"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AdminReadOrder godoc
// @Summary call AdminReadOrder, return ok by json.
// @가게에서 주문 목록 조회
// @name AdminReadOrder
// @Accept  json
// @Produce  json
// @Param id query string true "사업체 id"
// @Router /order/admin [GET]
// @Success 200 {object} Controller
func (p *Controller) AdminReadOrder(c *gin.Context) {
	id := c.Query("id")
	BID, res := p.adminReadOrderInputValidate(id)
	if res != nil {
		res.Response(c)
	}
	result := p.md.AdminListOrder(BID)
	result.Response(c)
}

func (p *Controller) adminReadOrderInputValidate(id string) (primitive.ObjectID, *protocol.ApiResponse[any]) {
	BID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, protocol.Fail(err, protocol.BadRequest)
	}
	return BID, nil
}

// UpdateOrderState godoc
// @Summary call UpdateOrderState, return ok by json.
// @가게에서 주문 상태 변경
// @name UpdateOrderState
// @Accept  json
// @Produce  json
// @Param input body UpdateOrderStateInput true "주문 번호, 상태 "
// @Router /order/admin [PATCH]
// @Success 200 {object} Controller
func (p *Controller) UpdateOrderState(c *gin.Context) {
	var input UpdateOrderStateInput
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		protocol.Fail(err, protocol.BadRequest).Response(c)
		return
	}
	orderId, err := primitive.ObjectIDFromHex(input.OrderId)
	if err != nil {
		protocol.Fail(err, protocol.BadRequest).Response(c)
		return
	}
	result := p.md.UpdateOrderState(orderId, input.State)
	result.Response(c)
}

type UpdateOrderStateInput struct {
	OrderId string `bson:"orderid" binding:"required"`
	State   int    `bson:"state" binding:"required,gte=1,lte=10"`
}
