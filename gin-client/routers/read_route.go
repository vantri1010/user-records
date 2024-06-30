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

func ReadRecord(conn *api.Api, opts *bind.CallOpts, addr common.Address) gin.HandlerFunc {
	return func(c *gin.Context) {
		repo := repository.NewRecordRepo(conn)
		recordBiz := biz.NewCreateBiz(repo)

		result, err := recordBiz.ReadRecord(c.Request.Context(), opts, addr)
		if err != nil {
			panic(err)
		}

		data, _ := model.MapSolDataToEmployee(&result)

		c.JSON(http.StatusOK, data)

	}
}
