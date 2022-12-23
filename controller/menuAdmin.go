package controller

import (
	"lecture/WBABEProject-23/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Controller) NewMenu(c *gin.Context) {
	var menu model.Menu
	if err := c.ShouldBind(&menu); err != nil {
		c.String(http.StatusBadRequest, "Bad request: %v", err)
		return
	}
	business := c.GetHeader("Business-Id")
	p.md.CreateNewMenu(menu, business)
	c.JSON(200, gin.H{"msg": "ok"})
	return
}
