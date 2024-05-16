package service

import (
	"errors"
	iauth "gwg/internal/auth"
	"gwg/src/dao"
	"gwg/src/entity"
	responsemodel "gwg/src/models/response"

	"gorm.io/gorm"
)

type User struct {
	itemsRes []responsemodel.User
	user     *iauth.User
}

func NewUser(user *iauth.User) User {
	return User{
		user:     user,
		itemsRes: make([]responsemodel.User, 0),
	}
}
func (u *User) List(conn *gorm.DB) (result responsemodel.UserAll) {
	err := dao.User{}.GetListWith(conn, u)
	if err != nil {
		return
	}

	result.Users = u.itemsRes
	result.Total = len(u.itemsRes)
	return
}
func (u *User) OnUserListGetterWith(_ *gorm.DB, _ int32, user entity.User, location entity.Location) error {
	u.itemsRes = append(u.itemsRes, responsemodel.User{
		UserID:   user.UserID,
		UserName: user.UserName,
		Age:      user.Age,
		Gender:   user.Gender,
		UserType: user.UserType,
		Status:   user.Status,
		Address: &responsemodel.Address{
			FullAddress: location.FullAddress,
			Province:    location.Province,
			District:    location.District,
			Ward:        location.Ward,
		},
	})

	return nil
}

func (u *User) GetInfo(conn *gorm.DB) (result responsemodel.User, err error) {
	user, err := dao.User{}.Get(conn, u.user.ID)
	if err != nil {
		err = errors.New("lỗi get data user")
		return
	}

	result = responsemodel.User{
		UserID:   user.UserID,
		UserName: user.UserName,
		Age:      user.Age,
		Gender:   user.Gender,
		UserType: user.UserType,
		Status:   user.Status,
	}

	return
}
