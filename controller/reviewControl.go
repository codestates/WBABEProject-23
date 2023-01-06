package controller

import (
	"lecture/WBABEProject-23/model/entitiy"
	"lecture/WBABEProject-23/protocol"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
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

func (p *Controller) createReviewInputValidate(body ReviewInput) (*entitiy.Review, *protocol.ApiResponse[any]) {
	review := new(entitiy.Review)
	var err error
	review.OrderID, err = primitive.ObjectIDFromHex(body.OrderID)
	if err != nil {
		return nil, protocol.Fail(err, protocol.BadRequest)
	}
	review.MenuID, err = primitive.ObjectIDFromHex(body.MenuID)
	if err != nil {
		protocol.Fail(err, protocol.BadRequest)
	}
	review.Orderer = body.Orderer
	review.Content = body.Content
	review.Score = body.Score
	if r := p.md.CheckOrderReviewable(review); r != nil {
		return nil, r
	}
	return review, nil
}

type ReviewInput struct {
	OrderID string  `bson:"order_id" json:"order_id"`
	MenuID  string  `bson:"menu_id" json:"menu_id"`
	Orderer string  `bson:"orderer"`
	Content string  `bson:"content"`
	Score   float32 `bson:"score" binding:"gte=0, lte=5"`
}

// ReadReviewControl godoc
// @Summary call ReadReviewControl, return ok by json.
// @메뉴 리뷰 조회 서비스
// @name ReadReviewControl
// @Accept  json
// @Produce  json
// @Param id query string true "메뉴 id"
// @Router /review [GET]
// @Success 200 {object} Controller
func (p *Controller) ReadReviewControl(c *gin.Context) {
	menuID := c.Query("id")
	id, res := p.readReviewInputValidate(menuID)
	if res != nil {
		res.Response(c)
		return
	}
	result := p.md.ReadReview(id)
	result.Response(c)
}

func (p *Controller) readReviewInputValidate(id string) (primitive.ObjectID, *protocol.ApiResponse[any]) {
	menuID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, protocol.Fail(err, protocol.BadRequest)
	}
	if r, e := p.md.CheckMenuID(menuID); !r {
		return primitive.NilObjectID, protocol.Fail(e, protocol.BadRequest)
	}
	return menuID, nil

}
