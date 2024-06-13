package routers

import (
	"github.com/gin-gonic/gin"
	"user-records/api"
	"user-records/gin-client/model"
)

func CreateRecord(conn *api.Api) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.Employee

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

	}
}
