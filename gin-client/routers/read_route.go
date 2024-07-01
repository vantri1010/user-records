package routers

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"user-records/api"
	"user-records/gin-client/biz"
	"user-records/gin-client/model"
	"user-records/gin-client/repository"
)

func ReadRecord(conn *api.Api, opts *bind.CallOpts) gin.HandlerFunc {
	return func(c *gin.Context) {
		addrString := c.Params.ByName("address")
		hexAddr := common.HexToAddress(addrString)

		repo := repository.NewRecordRepo(conn)
		recordBiz := biz.NewCreateBiz(repo)

		result, err := recordBiz.ReadRecord(c.Request.Context(), opts, hexAddr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			panic(err)
		}

		data, _ := model.MapSolDataToEmployee(result)

		c.JSON(http.StatusFound, data)
	}
}

func GetAddress(conn *api.Api, opts *bind.CallOpts) gin.HandlerFunc {
	return func(c *gin.Context) {
		indexString := c.Params.ByName("index")
		//bigIndex, _ := strconv.ParseInt(indexString, 10, 64)
		bigIndex, ok := new(big.Int).SetString(indexString, 10)
		if !ok {
			c.JSON(http.StatusBadRequest, model.ErrBadParamInput)
			return
		}

		repo := repository.NewRecordRepo(conn)
		recordBiz := biz.NewCreateBiz(repo)

		result, err := recordBiz.GetAddress(c.Request.Context(), opts, bigIndex)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			panic(err)
		}

		c.JSON(http.StatusFound, result)
	}
}
