package user

import (
	"github.com/1024casts/snake/internal/ecode"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/1024casts/snake/app/api"
	"github.com/1024casts/snake/internal/service"
	"github.com/1024casts/snake/pkg/errno"
	"github.com/1024casts/snake/pkg/log"
)

// Get 获取用户信息
// @Summary 通过用户id获取用户信息
// @Description Get an user by user id
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param id path string true "用户id"
// @Success 200 {object} model.UserInfo "用户信息"
// @Router /users/:id [get]
func Get(c *gin.Context) {
	log.Info("Get function called.")

	userID := cast.ToUint64(c.Param("id"))
	if userID == 0 {
		api.SendResponse(c, errno.ErrParam, nil)
		return
	}

	c.Request.Context()

	// Get the user by the `user_id` from the database.
	u, err := service.Svc.UserSvc().GetUserByID(c.Request.Context(), userID)
	if err != nil {
		log.Warnf("get user info err: %v", err)
		api.SendResponse(c, ecode.ErrUserNotFound, nil)
		return
	}

	api.SendResponse(c, nil, u)
}
