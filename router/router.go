package router

import (
	"main/infra"

	"github.com/gin-gonic/gin"
)

func NewRoutes(ctx *infra.IntegrationContext) *gin.Engine {
	r := gin.Default()

	PhoneRoutes(r, ctx)
	return r
}

func PhoneRoutes(r *gin.Engine, ctx *infra.IntegrationContext) {
	phoneRoute := r.Group("/api")
	{
		phoneRoute.GET("/number/cek", ctx.Ctl.CekNum.CekNumber)
		phoneRoute.GET("/number", ctx.Ctl.CekNum.GetAllNumbers)
		phoneRoute.GET("/number/find", ctx.Ctl.CekNum.FindNumber)
		phoneRoute.PUT("/number/update", ctx.Ctl.CekNum.UpdateNumber)
		phoneRoute.DELETE("/number/delete", ctx.Ctl.CekNum.DeleteNumber)
	}
}
