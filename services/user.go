package services

import (
	"context"
	"gin-research-sys/models"
	"gin-research-sys/pkg/global"
)

type IUserService interface {
	FindUserByUsername(user *models.User) error

	List(page int, size int, users *[]models.User, total *int64) error
	Retrieve(user *models.User, id int) error
	Create(user *models.User) error
	Update(user *models.User) error
	Destroy(id int) error

	ListRole2(user *models.User, roles *[]int) error
	UpdateRole(user *models.User, ids []int) error
}
type UserService struct{}

func NewUserService() IUserService {
	return UserService{}
}

var ctx = context.Background()

func (u UserService) FindUserByUsername(user *models.User) error {
	result := global.Mysql.
		Where("username = ?", user.Username).
		First(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u UserService) List(page int, size int, users *[]models.User, total *int64) error {
	if err := global.Mysql.Model(&models.User{}).Count(total).
		Scopes(global.Paginate(page, size)).
		Find(&users).Error; err != nil {
		return err
	}
	return nil
}

func (u UserService) Retrieve(user *models.User, id int) error {
	if err := global.Mysql.Model(&models.User{}).
		Preload("Roles").
		First(&user, id).Error; err != nil {
		return err
	}
	return nil

	//val, err := global.Redis.Get(ctx, strconv.Itoa(id)).Result()
	//if err != redis.Nil {
	//	return json.Unmarshal([]byte(val), user)
	//}
	//if err = global.Mysql.Model(&models.User{}).
	//	Preload("Roles").
	//	First(&user, id).Error; err != nil {
	//	return err
	//}
	//marshal, err := json.Marshal(user)
	//if err != nil {
	//	return err
	//}
	//_, err = global.Redis.SetEX(ctx, strconv.Itoa(id), marshal, time.Hour).Result()
	//if err != nil {
	//	return err
	//}
	//return nil
}

func (u UserService) Create(user *models.User) error {
	if err := global.Mysql.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u UserService) Update(user *models.User) error {
	if err := global.Mysql.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u UserService) Destroy(id int) error {
	if err := global.Mysql.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (u UserService) ListRole2(user *models.User, roles *[]int) error {
	var t []int
	for _, value := range user.Roles {
		t = append(t, int(value.ID))
	}
	*roles = t
	return nil
}

func (u UserService) UpdateRole(user *models.User, ids []int) error {
	var roles []models.Role
	if err := global.Mysql.Model(&models.Role{}).Find(&roles, "id IN ?", ids).Error; err != nil {
		return err
	}
	if err := global.Mysql.Model(&user).Association("Roles").Replace(roles); err != nil {
		return err
	}
	return nil
}
