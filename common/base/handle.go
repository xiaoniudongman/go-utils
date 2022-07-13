package base

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/xiaoniudongman/go-utils/tools/errs"

	"github.com/gin-gonic/gin"
)

func SendResponse(c *gin.Context, retCode ResponseCode, rsp interface{}) {
	rspData := Response{}
	rspData.Code = retCode
	rspData.Msg = responseCodeToMsg[retCode]
	if retCode == ResponseCode_Succ {
		rspData.Data = rsp
	}

	c.JSON(http.StatusOK, rspData)
}

func NowFunc() string {
	pc, _, _, _ := runtime.Caller(1)
	return "NowFunc:" + runtime.FuncForPC(pc).Name() + " "
}

func NowFuncError() string {
	pc, _, _, _ := runtime.Caller(1)
	return "NowFunc:" + runtime.FuncForPC(pc).Name() + " Error "
}

func RecoverFunc(c *gin.Context) {
	if rec := recover(); rec != nil {
		c.Header("Content-Type", "text/json; charset=utf-8")
		c.String(http.StatusInternalServerError, "[]")
		buf := make([]byte, 4096)
		n := runtime.Stack(buf, false)
		errs.CheckErrSendEmail(fmt.Errorf("recovery:%s\nstack:%s", rec, string(buf[:n])))
	}
}
