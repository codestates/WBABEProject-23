package controller

import (
	"lecture/WBABEProject-23/protocol"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AdminListOrderController godoc
// @Summary call AdminListOrderController, return ok by json.
// @가게에서 주문 목록 조회
// @name AdminListOrderController
// @Accept  json
// @Produce  json
// @Param id query string true "사업체 id"
// @Router /order/admin [GET]
// @Success 200 {object} Controller
func (p *Controller) AdminListOrderController(c *gin.Context) {
	id := c.Query("id")
	BID, res := p.adminListOrderInputValidate(id)
	if res != nil {
		res.Response(c)
	}
	result := p.md.AdminListOrder(BID)
	result.Response(c)
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
// @Router /order/admin [PATCH]
// @Success 200 {object} Controller
func (p *Controller) UpdateState(c *gin.Context) {
	var input UpdateStateInput
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

type UpdateStateInput struct {
	OrderId string `bson:"orderid" binding:"required"`
	State   int    `bson:"state" binding:"required,gte=1,lte=10"`
}
