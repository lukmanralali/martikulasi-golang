package v2

import (
	"../../constants"
	"../../helpers"
	"../../objects"
	"../../services/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type V2UserController struct {
	userService v2.V2UserService
	errorHelper helpers.ErrorHelper
}

func V2UserControllerHandler(router *gin.Engine) {

	handler := &V2UserController{
		userService: v2.V2UserServiceHandler(),
		errorHelper: helpers.ErrorHelperHandler(),
	}

	group := router.Group("v2/users")
	{
		group.POST(":id", handler.UpdateById)
	}

}

func (handler *V2UserController) UpdateById(context *gin.Context) {

	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.RequestParameterInvalid)
	}

	requestObject := objects.V2UserObjectRequest{}
	context.ShouldBind(&requestObject)

	result, err := handler.userService.UpdateById(id, requestObject)
	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.InternalServerError)
	}

	context.JSON(http.StatusOK, result)

}
