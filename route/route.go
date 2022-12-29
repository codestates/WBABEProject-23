package route

import (
	"fmt"
	ctl "lecture/WBABEProject-23/controller"
	"lecture/WBABEProject-23/docs"

	"github.com/gin-gonic/gin"
	swgFiles "github.com/swaggo/files"
	ginSwg "github.com/swaggo/gin-swagger"

	"lecture/WBABEProject-23/logger"
)

type Router struct {
	ct *ctl.Controller
}

func NewRouter(ctl *ctl.Controller) (*Router, error) {
	r := &Router{ct: ctl}

	return r, nil
}

// cross domain을 위해 사용
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
		CORS 허용을 위해서 모든 도메인을 허용한다면 보안에 이슈가 생깁니다. 
		보통 운영되는 시스템의 경우는 특정한 도메인만을 허용하고 그 이외의 요청은 거부하도록 설정합니다.
		*/
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//허용할 header 타입에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Forwarded-For, Authorization, accept, origin, Cache-Control, X-Requested-With")
		//허용할 method에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// 임의 인증을 위한 함수
func liteAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c == nil {
			c.Abort() // 미들웨어에서 사용, 이후 요청 중지
			return
		}
		//http 헤더내 "Authorization" 폼의 데이터를 조회
		auth := c.GetHeader("Authorization")
		//실제 인증기능이 올수있다. 단순히 출력기능만 처리 현재는 출력예시
		fmt.Println("Authorization-word ", auth)

		c.Next() // 다음 요청 진행
	}
}

func (p *Router) Index() *gin.Engine {

	// gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.DebugMode)

	e := gin.New()
	e.Use(logger.GinLogger())
	e.Use(logger.GinRecovery(true))
	e.Use(CORS())

	logger.Info("start server")

	e.GET("/swagger/:any", ginSwg.WrapHandler(swgFiles.Handler))
	docs.SwaggerInfo.Host = "localhost:8080"

	/*
	엔드포인트 구성에 대해서 전반적인 코멘트 드립니다.

	1. REST API의 성숙도 모델에 대해서 공부해보시면 좋을 것 같습니다.

	2. 일반적으로 HTTP URI에 new, modify 와 같은 행위는 들어가지 않습니다. 
		복수형의 단어로 구성을 하고, 동일한 URI 내에서 http method만 변경하여 행위를 표현하는 것이 일반적인 REST API의 구성 방식입니다.

		e.g.
		GET v1/menus -> 메뉴 목록을 조회.
		GET v1/menus/1 -> 1번 메뉴를 조회.
		POST v1/menus -> 메뉴를 생성.
		PATCH v1/menus/1 -> 1번 메뉴에 대해서 업데이트
		DELETE v1/menus/1 -> 1번 메뉴에 대해서 삭제
	*/

	/*
	Group을 통해서 나누어주니 가독성이 좋아 보이네요. 좋은 코드입니다.
	*/
	menuAdmin := e.Group("/menu/admin", liteAuth())
	{
		menuAdmin.POST("/new", p.ct.NewMenu)
		menuAdmin.PATCH("modify", p.ct.ModifyMenu)

	}
	menuService := e.Group("/menu", liteAuth())
	{
		menuService.GET("/list", p.ct.MenuList)
		menuService.GET("/list/review", p.ct.MenuReadReview)
	}
	order := e.Group("/order", liteAuth())
	{
		order.POST("/make", p.ct.MakeOrder) //주문

		order.GET("/list", p.ct.ListOrder)                      //주문 조회
		order.GET("/admin/list", p.ct.AdminListOrderController) //주문 상태 조회
		order.PATCH("/modify", p.ct.ModifyOrder)                //주문 변경
		order.PATCH("/admin/update", p.ct.UpdateState)          //주문 상태 변경
	}
	review := e.Group("review", liteAuth())
	{
		review.POST("", p.ct.MakeReview) //리뷰 작성
	}

	return e
}
