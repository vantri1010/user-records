package routers

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"net/http"
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
		createRcBiz := biz.NewCreateBiz(repo)

		err = createRcBiz.CreateRecord(c.Request.Context(), txOpts, solData)
		if err != nil {
			panic(err)
		}

	}
}

func DeleteRecord(conn *api.Api, txOpts *bind.TransactOpts) gin.HandlerFunc {
	return func(c *gin.Context) {
		addrString := c.Params.ByName("address")
		hexAddr := common.HexToAddress(addrString)
		repo := repository.NewRecordRepo(conn)
		deleteBiz := biz.NewCreateBiz(repo)

		rowDeleted, err := deleteBiz.DeleteUser(c.Request.Context(), txOpts, hexAddr)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, rowDeleted)
	}
}
