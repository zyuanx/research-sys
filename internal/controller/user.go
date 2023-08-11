package controller

import (
	"gin-research-sys/internal/util"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zyuanx/research-sys/internal/form"
	"github.com/zyuanx/research-sys/internal/model"
	"github.com/zyuanx/research-sys/internal/pkg/constant"
	"github.com/zyuanx/research-sys/internal/pkg/errors"
	"github.com/zyuanx/research-sys/internal/pkg/errors/ecode"
	"github.com/zyuanx/research-sys/internal/pkg/response"
	"golang.org/x/crypto/bcrypt"
)

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
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "参数错误")
		return
	}
	if passwordForm.Password1 != passwordForm.Password2 {
		util.Fail(ctx, gin.H{}, "两次密码不一致")
		return
	}
	user := model.User{}
	instance := middleware.JWTAuthMiddleware.IdentityHandler(ctx).(*model.User)
	if err = c.service.UserRetrieve(&user, int(instance.ID)); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "未找到记录")
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordForm.Password)); err != nil {
		util.Fail(ctx, gin.H{}, "原密码错误")
		return
	}
	var hashedPassword []byte
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(passwordForm.Password1), bcrypt.DefaultCost); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "生成密码错误")
		return
	}
	payload := map[string]interface{}{
		"password": string(hashedPassword),
	}
	user.Password = string(hashedPassword)
	if err = c.service.UserUpdate(&user, payload); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "修改密码错误")
		return
	}
	util.Success(ctx, gin.H{}, "修改密码成功")
}

func (c *Controller) UserList(ctx *gin.Context) {
	userListQuery := form.UserListQuery{}
	if err := ctx.ShouldBindQuery(&userListQuery); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, nil, "参数错误")
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
		log.Println(err.Error())
		util.Fail(ctx, nil, "获取数据失败")
		return
	}
	util.Success(ctx, gin.H{
		"page":    page,
		"size":    size,
		"results": users,
		"total":   total,
	}, "")
}

func (c *Controller) UserRetrieve(ctx *gin.Context) {
	var id int
	var err error
	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "ID错误")
		return
	}
	user := model.User{}
	if err = c.service.UserRetrieve(&user, id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "未找到记录")
		return
	}

	util.Success(ctx, gin.H{"user": user}, "获取用户信息成功")
}

func (c *Controller) UserCreate(ctx *gin.Context) {
	createForm := form.UserCreateForm{}
	var err error
	if err = ctx.ShouldBindJSON(&createForm); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "参数错误")
		return
	}
	if createForm.Password1 != createForm.Password2 {
		util.Fail(ctx, gin.H{}, "两次密码不一致")
		return
	}
	var hashedPassword []byte
	if hashedPassword, err = bcrypt.GenerateFromPassword([]byte(createForm.Password1), bcrypt.DefaultCost); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "生成密码错误")
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
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "创建用户失败")
		return
	}
	// 更新为公共角色
	//if err = userServices.UpdateRole(&user, []int{1}); err != nil {
	//	log.Println(err.Error())
	//	util.Fail(ctx, gin.H{}, "更新角色失败")
	//	return
	//}
	util.Success(ctx, gin.H{}, "创建用户成功")
}

func (c *Controller) UserUpdate(ctx *gin.Context) {
	var id int
	var err error
	id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "ID错误")
		return
	}
	updateForm := form.UserUpdateForm{}
	if err = ctx.ShouldBindJSON(&updateForm); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "参数错误")
		return
	}
	user := model.User{}
	if err = c.service.UserRetrieve(&user, id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "获取信息失败")
		return
	}
	payload := map[string]interface{}{
		"nickname":  updateForm.Nickname,
		"telephone": updateForm.Telephone,
		"email":     updateForm.Email,
		//"roles":     updateForm.Roles,
	}
	if err = c.service.UserUpdate(&user, payload); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "更新用户失败")
		return
	}
	if err = c.service.UserUpdateRole(&user, updateForm.Roles); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "更新用户失败")
		return
	}
	util.Success(ctx, gin.H{}, "更新成功")
}

func (c *Controller) UserDestroy(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "参数错误")
		return
	}
	if err = c.service.UserDestroy(id); err != nil {
		log.Println(err.Error())
		util.Fail(ctx, gin.H{}, "删除用户失败")
		return
	}
	util.Success(ctx, gin.H{}, "删除用户成功")
}
