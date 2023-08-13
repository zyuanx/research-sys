package controller

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zyuanx/research-sys/internal/form"
	"github.com/zyuanx/research-sys/internal/model"
	"github.com/zyuanx/research-sys/internal/pkg/config"
	"github.com/zyuanx/research-sys/internal/pkg/constant"
	"github.com/zyuanx/research-sys/internal/pkg/errors"
	"github.com/zyuanx/research-sys/internal/pkg/errors/ecode"
	"github.com/zyuanx/research-sys/internal/pkg/jwt"
	"github.com/zyuanx/research-sys/internal/pkg/response"
	"golang.org/x/crypto/bcrypt"
)

func (c *Controller) UserLogin(ctx *gin.Context) {
	payload := model.UserLoginReq{}
	var err error
	if err = ctx.ShouldBindJSON(&payload); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "参数错误")
		response.JSON(ctx, err, nil)
		return
	}
	user := model.User{}
	if err = c.service.UserFindByUsername(&user, payload.Username); err != nil {
		err = errors.Wrap(err, ecode.RecordRetrieveErr, "未找到记录")
		response.JSON(ctx, err, nil)
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		err = errors.Wrap(err, ecode.UserLoginErr, "密码错误")
		response.JSON(ctx, err, nil)
		return
	}
	expireAt := time.Now().Add(24 * 7 * time.Hour)
	claims := jwt.BuildClaims(expireAt, int64(user.ID))
	token, err := jwt.GenToken(claims, config.GlobalConfig.JwtSecret)
	if err != nil {
		response.JSON(ctx, errors.Wrap(err, ecode.UserLoginErr, "生成用户授权token失败"), nil)
		return
	}
	response.JSON(ctx, nil, struct {
		Token    string    `json:"token"`
		ExpireAt time.Time `json:"expire_at"`
	}{
		Token:    token,
		ExpireAt: expireAt,
	})

}

func (c *Controller) UserGetInfo(ctx *gin.Context) {
	id, exist := ctx.Get(constant.UserID)
	if !exist {
		var err error
		response.JSON(ctx, errors.Wrap(err, ecode.AuthTokenErr, "未登录"), nil)
		return
	}
	user := model.User{}
	if err := c.service.UserRetrieve(&user, id.(int)); err != nil {
		response.JSON(ctx, errors.Wrap(err, ecode.NotFoundErr, "未找到记录"), nil)
		return
	}
	response.JSON(ctx, nil, user)
}

func (c *Controller) UserChangePassword(ctx *gin.Context) {
	passwordForm := form.UserChangePasswordForm{}
	var err error
	if err = ctx.ShouldBindJSON(&passwordForm); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "参数错误")
		response.JSON(ctx, err, nil)
		return
	}
	if passwordForm.Password1 != passwordForm.Password2 {
		err = errors.Wrap(err, ecode.ValidateErr, "两次密码不一致")
		response.JSON(ctx, err, nil)
		return
	}
	user := model.User{}
	userId, exist := ctx.Get(constant.UserID)
	if !exist {
		err = errors.Wrap(err, ecode.AuthTokenErr, "未登录")
		response.JSON(ctx, err, nil)
		return
	}
	if err = c.service.UserRetrieve(&user, userId.(int)); err != nil {
		err = errors.Wrap(err, ecode.NotFoundErr, "未找到记录")
		response.JSON(ctx, err, nil)
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordForm.Password)); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "原密码错误")
		response.JSON(ctx, err, nil)
		return
	}
	var hashedPassword []byte
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(passwordForm.Password1), bcrypt.DefaultCost); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "生成密码错误")
		response.JSON(ctx, err, nil)
		return
	}
	payload := map[string]interface{}{
		"password": string(hashedPassword),
	}
	user.Password = string(hashedPassword)
	if err = c.service.UserUpdate(&user, payload); err != nil {
		err = errors.Wrap(err, ecode.RecordUpdateErr, "更新用户失败")
		response.JSON(ctx, err, nil)
		return
	}
	response.JSON(ctx, nil, nil)
}

func (c *Controller) UserList(ctx *gin.Context) {
	userListQuery := form.UserListQuery{}
	if err := ctx.ShouldBindQuery(&userListQuery); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "参数错误")
		response.JSON(ctx, err, nil)
		return
	}
	var users []model.User
	page, size := userListQuery.Page, userListQuery.Size
	var total int64
	query := make(map[string]interface{})
	if userListQuery.Username != "" {
		query["username"] = userListQuery.Username
	}
	if userListQuery.Name != "" {
		query["name"] = userListQuery.Name
	}
	if err := c.service.UserList(&users, page, size, &total, query); err != nil {
		err = errors.Wrap(err, ecode.RecordListErr, "获取数据失败")
		response.JSON(ctx, err, nil)
		return
	}
	response.JSON(ctx, nil, gin.H{
		"page":    page,
		"size":    size,
		"results": users,
		"total":   total,
	})
}

func (c *Controller) UserRetrieve(ctx *gin.Context) {
	var id int
	var err error
	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "ID错误")
		response.JSON(ctx, err, nil)
		return
	}
	user := model.User{}
	if err = c.service.UserRetrieve(&user, id); err != nil {
		err = errors.Wrap(err, ecode.RecordRetrieveErr, "未找到记录")
		response.JSON(ctx, err, nil)
		return
	}

	response.JSON(ctx, nil, user)
}

func (c *Controller) UserCreate(ctx *gin.Context) {
	createForm := form.UserCreateForm{}
	var err error
	if err = ctx.ShouldBindJSON(&createForm); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "参数错误")
		response.JSON(ctx, err, nil)
		return
	}
	if createForm.Password1 != createForm.Password2 {
		err = errors.Wrap(err, ecode.ValidateErr, "两次密码不一致")
		response.JSON(ctx, err, nil)
		return
	}
	var hashedPassword []byte
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(createForm.Password1), bcrypt.DefaultCost); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "生成密码错误")
		response.JSON(ctx, err, nil)
		return
	}
	user := model.User{
		Username:  createForm.Username,
		Nickname:  createForm.Nickname,
		Password:  string(hashedPassword),
		Telephone: createForm.Telephone,
		Email:     createForm.Email,
	}
	if err = c.service.UserCreate(&user); err != nil {
		err = errors.Wrap(err, ecode.RecordCreateErr, "创建用户失败")
		response.JSON(ctx, err, nil)
		return
	}
	// 更新为公共角色
	//if err = userServices.UpdateRole(&user, []int{1}); err != nil {
	//	log.Println(err.Error())
	//	util.Fail(ctx, gin.H{}, "更新角色失败")
	//	return
	//}
	response.JSON(ctx, nil, user)
}

func (c *Controller) UserUpdate(ctx *gin.Context) {
	var id int
	var err error
	id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "ID错误")
		response.JSON(ctx, err, nil)
		return
	}
	updateForm := form.UserUpdateForm{}
	if err = ctx.ShouldBindJSON(&updateForm); err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "参数错误")
		response.JSON(ctx, err, nil)
		return
	}
	user := model.User{}
	if err = c.service.UserRetrieve(&user, id); err != nil {
		err = errors.Wrap(err, ecode.RecordRetrieveErr, "获取记录失败")
		response.JSON(ctx, err, nil)
		return
	}
	payload := map[string]interface{}{
		"nickname":  updateForm.Nickname,
		"telephone": updateForm.Telephone,
		"email":     updateForm.Email,
		//"roles":     updateForm.Roles,
	}
	if err = c.service.UserUpdate(&user, payload); err != nil {
		err = errors.Wrap(err, ecode.RecordUpdateErr, "更新用户失败")
		response.JSON(ctx, err, nil)
		return
	}
	if err = c.service.UserUpdateRole(&user, updateForm.Roles); err != nil {
		err = errors.Wrap(err, ecode.RecordUpdateErr, "更新用户失败")
		response.JSON(ctx, err, nil)
		return
	}
	response.JSON(ctx, nil, user)
}

func (c *Controller) UserDelete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		err = errors.Wrap(err, ecode.ValidateErr, "ID错误")
		response.JSON(ctx, err, nil)
		return
	}
	if err = c.service.UserDestroy(id); err != nil {
		err = errors.Wrap(err, ecode.RecordDeleteErr, "删除用户失败")
		response.JSON(ctx, err, nil)
		return
	}
	response.JSON(ctx, nil, nil)
}
