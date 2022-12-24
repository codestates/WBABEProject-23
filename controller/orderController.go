package controller

import (
	"lecture/WBABEProject-23/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ModifyOrderInput struct {
	BusinessID string          `bson:"businessid"`
	Menu       []model.MenuNum `bson:"menu"`
}

// MakeOrder godoc
// @Summary call MakeOrder, return ok by json.
// @주문.
// @name MakeOrder
// @Accept  json
// @Produce  json
// @Param id body model.Order true "User input 주문자 이름, 주문 가게 이름, 메뉴 배열형태만 입력 ]"
// @Router /order/make [POST]
// @Success 200 {object} Controller
func (p *Controller) MakeOrder(c *gin.Context) {
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		panic(err)
	}
	order := model.Order{}
	if err := c.ShouldBind(&order); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	order.Id = primitive.NewObjectID()
	order.CreatedAt = time.Now().In(loc)
	order.Status = model.Receipting
	p.md.MakeOrder(order)
	c.JSON(200, gin.H{"msg": "ok"})
}

func (p *Controller) ListOrder(c *gin.Context) {
	userName := c.Query("name")
	result := p.md.ListOrder(userName)

	c.JSON(200, gin.H{"msg": "ok", "list": result})
}

func (p *Controller) ModifyOrder(c *gin.Context) {
	var input ModifyOrderInput
	if err := c.ShouldBind(&input); err != nil {
		panic(err)
	}
	objID, err := primitive.ObjectIDFromHex(input.BusinessID)
	if err != nil {
		panic(err)
	}
	result := p.md.ModifyOrder(objID, input.Menu)
	if result {
		c.JSON(200, gin.H{"msg": "update request success"})
	} else {
		c.JSON(200, gin.H{"msg": "update request failed"})
	}

}
