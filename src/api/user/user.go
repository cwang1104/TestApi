package apiUser

import (
	"TestApi/common/config"
	"TestApi/common/grpc-etcdv3/getcdv3"
	pbUser "TestApi/proto/user"
	"context"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RequestParams struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func SayHello(c *gin.Context) {
	var req RequestParams
	if err := c.ShouldBindJSON(&req); err != nil {
		logs.Debug("get params err", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "参数错误",
		})
	}

	pbParams := pbUser.HelloRequest{
		Name: req.Name,
		Age:  int32(req.Age),
	}

	conn, err := getcdv3.GetGrpcConn(config.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	}

	client := pbUser.NewSayHelloClient(conn)
	reply, err := client.SayHelloToWho(context.Background(), &pbParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"reply": reply,
	})
}
