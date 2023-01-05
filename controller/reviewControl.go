package controller

import (
	"lecture/WBABEProject-23/model"
	"lecture/WBABEProject-23/protocol"

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
	if err := c.ShouldBind(&input); err != nil {
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
