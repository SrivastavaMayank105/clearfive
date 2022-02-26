package service

import (
	"github.com/gin-gonic/gin"	
	"ClearFive/mapper" 
	"errors"
	"strings"
	"math/rand"
)

type GetUserData interface{
	ValidateUserData(ctx *gin.Context,reqest mapper.DataRequest) (mapper.DataResponse,error)
}
type userProfileData struct{}

func NewUserProfileService()GetUserData{
	return userProfileData{}
}

func (userProfile userProfileData)ValidateUserData(ctx *gin.Context,request mapper.DataRequest)(mapper.DataResponse,error){
	
	var response mapper.DataResponse
	trmiMessage:=strings.TrimSpace(request.Message)
	if len(trmiMessage)==0{
		return response,errors.New("invalid input")
	}

	randomNum:=rand.Int()
	if strings.ToUpper(trmiMessage) =="YES"{
		return mapper.DataResponse{
			ReqID:randomNum,
			ReqStatus: mapper.ReqStatus{
				Status:"Success",
				ErrorCode:"Code:00",
				ErrorMsg:"no error",
			},
		},nil
	}
	
	return mapper.DataResponse{
		ReqID:-1,
		ReqStatus: mapper.ReqStatus{
			Status:"Failed",
			ErrorCode:"Code:6998",
			ErrorMsg:"no such user info present",
		},
	},nil
}
