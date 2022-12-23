package controller

import (
	"lecture/WBABEProject-23/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (p *Controller) MakeOrder(c *gin.Context) {
	order := model.Order{}
	if err := c.ShouldBind(&order); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	order.Id = primitive.NewObjectID()
	order.CreatedAt = time.Now()
	order.Status = model.Receipting
	p.md.MakeOrder(order)
	c.JSON(200, gin.H{"msg": "ok"})
}
