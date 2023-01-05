package controller

import (
	"lecture/WBABEProject-23/protocol"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

// ListMenu godoc
// @Summary call ListMenu, return ok by json.
// @메뉴 조회
// @name ListMenu
// @Accept  json
// @Produce  json
// @Param id query string true "id"
// @Param sort query string true "sort할 컬럼이름"
// @Param order query string true "order= 1은 오름찬순 그 외 내림차순 "
// @Router /menu [GET]
// @Success 200 {object} Controller
func (p *Controller) ListMenu(c *gin.Context) {
	id := c.Query("id")
	sortBy := c.Query("sort")
	sortOrder := c.Query("order")
	bID, order, res := p.listMenuValidate(id, sortBy, sortOrder)
	if res != nil {
		res.Response(c)
		return
	}
	result := p.md.ListMenuModel(bID, sortBy, order)
	result.Response(c)
}

func (p *Controller) listMenuValidate(id, sort, order string) (primitive.ObjectID, int, *protocol.ApiResponse[any]) {
	bID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, 0, protocol.Fail(err, protocol.BadRequest)
	}
	if r, e := p.md.CheckBusinessID(bID); !r {
		return primitive.NilObjectID, 0, protocol.FailCustomMessage(e, "No document was found with the business id", protocol.BadRequest)
	}
	if r, e := p.md.CheckMenuFieldExists(sort); r {
		if order == "1" {
			return bID, 1, nil
		} else {
			return bID, -1, nil
		}

	} else {
		return primitive.NilObjectID, 0, protocol.FailCustomMessage(e, "No document was found with the business id", protocol.BadRequest)
	}
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
