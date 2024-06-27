package routers

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gin-gonic/gin"
	"user-records/api"
	"user-records/gin-client/biz"
	"user-records/gin-client/model"
	"user-records/gin-client/repository"
)

func CreateRecord(conn *api.Api, txOpts *bind.TransactOpts) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.Employee

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		solData, err := model.MapToSolData(data)
		if err != nil {
			panic(err)
		}

		repo := repository.NewRecordRepo(conn)
		record := biz.NewRecord(repo)

		err = record.CreateRecord(c.Request.Context(), txOpts, solData)
		if err != nil {
			panic(err)
		}

	}
}
