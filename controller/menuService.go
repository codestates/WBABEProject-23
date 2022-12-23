package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (p *Controller) MenuList(c *gin.Context) {
	businessName := c.Query("name")
	sortBy := c.Query("sort")
	sortOrder := c.Query("order")
	var order int
	var err error
	if sortOrder == "" {
		order = 1
	} else {
		order, err = strconv.Atoi(sortOrder)
		if err != nil {
			panic(err)
		}
	}
	menu := p.md.ListMenu(businessName, sortBy, order)
	c.JSON(200, gin.H{"msg": "ok", "list": menu})
}

func (p *Controller) MenuReadReview(c *gin.Context) {

}
