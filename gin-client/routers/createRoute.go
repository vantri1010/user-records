package routers

import (
	"github.com/gin-gonic/gin"
	"user-records/api"
	"user-records/gin-client/biz"
	"user-records/gin-client/model"
	"user-records/gin-client/repository"
)

func CreateRecord(conn *api.Api) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.Employee

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		solData, err := model.MapToSolData(data)
		if err != nil {
			panic(err)
		}

		repo := repository.NewRecordRepo()
		record := biz.NewRecord()

	}
}
