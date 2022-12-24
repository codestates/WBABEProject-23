package controller

import (
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

// MenuList godoc
// @Summary call MenuList, return ok by json.
// @메뉴 조회
// @name MenuList
// @Accept  json
// @Produce  json
// @Param id query string true "User input name"
// @Param id query string true "User input sort할 컬럼이름"
// @Param id query string true "User input order= 1은 오름찬순 그 외 내림차순 "
// @Router /menu/list [GET]
// @Success 200 {object} Controller
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

// MenuReadReview godoc
// @Summary call MenuReadReview, return ok by json.
// @메뉴 리뷰 조회 서비스
// @name MenuReadReview
// @Accept  json
// @Produce  json
// @Param id query string true "가게 이름"
// @Param id query string true "메뉴 이름"
// @Router /menu/list/review [GET]
// @Success 200 {object} Controller
func (p *Controller) MenuReadReview(c *gin.Context) {
	businessName := c.Query("id")
	menuName := c.Query("name")
	objId, err := primitive.ObjectIDFromHex(businessName)
	if err != nil {
		panic(err)
	}
	result := p.md.ReadMenuReview(objId, menuName)
	c.JSON(200, gin.H{"msg": "ok", "list": result})
}
