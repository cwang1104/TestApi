package apiLogin

import (
	"TestApi/common/config"
	"TestApi/common/grpc-etcdv3/getcdv3"
	pbLogin "TestApi/proto/login"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type paramsAccountRegister struct {
	Phone    string `json:"phone" binding:"required,min=1,max=32"`
	Code     int32  `json:"code" binding:"required"`
	Platform int32  `json:"platform" binding:"required"`
}

func AccountRegister(c *gin.Context) {
	paramsReq := paramsAccountRegister{}
	if err := c.ShouldBindJSON(&paramsReq); err != nil {
		fmt.Println("get prams failed , err : ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	//todo:从etcd服务发现获取调用

	// grpc调用
	pbdata := pbLogin.AccountRegisterReq{
		Phone:    paramsReq.Phone,
		Platform: paramsReq.Platform,
	}

	//conn, err := grpc.Dial("127.0.0.1:52201", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	fmt.Println("conn failed ; err = ", err)
	//	return
	//}

	//conn := getcdv3.GetConn(config.EtcdAddress, config.LoginName)

	conn, err := getcdv3.GetGrpcConn(config.LoginName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errCode": 500,
			"errMsg":  err.Error(),
		})
	}
	client := pbLogin.NewUserLoginClient(conn)
	reply, err := client.AccountRegister(context.Background(), &pbdata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": err.Error()})
		return
	}

	log.Println(reply)

	c.JSON(
		http.StatusOK,
		gin.H{
			"errCode": reply.ErrCode,
			"errMsg":  reply.ErrMsg,
			"data": gin.H{
				"data": reply.Data,
			},
		},
	)

}
