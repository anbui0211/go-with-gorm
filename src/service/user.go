package service

import (
	"gwg/src/dao"
	"gwg/src/entity"
	responsemodel "gwg/src/models/response"

	"gorm.io/gorm"
)

type User struct {
	itemsRes []responsemodel.User
}

func (u *User) List(conn *gorm.DB) (result []responsemodel.User) {
	err := dao.User{}.GetList(conn, u)
	if err != nil {
		return
	}

	return
}
func (s *User) OnUserListGetter(_ *gorm.DB, count int32, rec entity.User) error {
	s.itemsRes = append(s.itemsRes, responsemodel.User{
		UserID:   rec.UserID,
		UserName: rec.UserName,
		Age:      rec.Age,
		Gender:   rec.Gender,
		UserType: rec.UserType,
		Status:   rec.Status,
		Address:  responsemodel.Address{},
	})

	return nil
}
