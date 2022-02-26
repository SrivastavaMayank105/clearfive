package controller

import (
	"github.com/gin-gonic/gin"
	"ClearFive/mapper"
	"ClearFive/service"
	"net/http"
	"fmt"
	"strconv"
)
func GetUserDetails(ctx *gin.Context){
	var request mapper.DataRequest

	err:=ctx.ShouldBindJSON(&request);if err !=nil{
		bindErr :=fmt.Errorf("Bad requedt")
		ctx.AbortWithStatusJSON(http.StatusBadRequest,bindErr.Error())
		return
	}

	if request.Age <0 {
		err:=fmt.Errorf("invalid age")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,err.Error())
	}
	
	mob:=strconv.Itoa(request.MobileNuber)
	if len(mob) !=10 || string(mob[0])=="0"{
		err:=fmt.Errorf("invalid mobile number")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,err.Error())
	}

	userProfileRes:=service.NewUserProfileService()
	response,responseErr:=userProfileRes.ValidateUserData(ctx,request)
	if responseErr!=nil{
		err:=fmt.Errorf("unable to count the data")
		ctx.AbortWithStatusJSON(http.StatusServiceUnavailable,err.Error())
		return
	}

	ctx.JSON(http.StatusOK,response)
}