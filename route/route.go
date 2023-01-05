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

	menu := e.Group("/menu", liteAuth())
	{
		menu.POST("", p.ct.CreateMenu)
		menu.PATCH("", p.ct.UpdateMenu)
		menu.GET("", p.ct.ListMenu)
	}
	order := e.Group("/order", liteAuth())
	{
		order.POST("", p.ct.CreateOrder)                   //주문
		order.GET("", p.ct.ListOrder)                      //주문 조회
		order.GET("/admin", p.ct.AdminListOrderController) //주문 상태 조회
		order.PATCH("", p.ct.UpdateOrder)                  //주문 변경
		order.PATCH("/admin", p.ct.UpdateState)            //주문 상태 변경
	}
	review := e.Group("/review", liteAuth())
	{
		review.POST("", p.ct.CreateReview) //리뷰 작성
		review.GET("", p.ct.ReadReviewControl)
	}

	return e
}

// cross domain을 위해 사용
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
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
