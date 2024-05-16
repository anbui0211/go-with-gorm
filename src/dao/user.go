package dao

import (
	ilogger "gwg/internal/logger"
	"gwg/src/entity"

	"gorm.io/gorm"
)

type UserListGetter interface {
	OnUserListGetter(conn *gorm.DB, count int32, rec entity.User) error
}

type User struct{}

func (User) GetList(conn *gorm.DB, svcFunc UserListGetter) (err error) {
	rows, err := conn.Table("users as u ").
		Select("users.*, locations.location_id, locations.full_address, locations.province, locations.district, locations.ward").
		Joins("LEFT JOIN locations ON users.location_id = locations.location_id").
		Rows()
	if err != nil {
		ilogger.Error("UserDao - GetList", ilogger.LogData{
			"err": err,
		})
		return
	}
	defer rows.Close()

	var count int32
	for rows.Next() {
		count++
		var rec entity.User

		err = conn.ScanRows(rows, &rec)
		if err != nil {
			ilogger.Error("UserDao - ScanRows", ilogger.LogData{
				"err": err,
			})
		}

		err = svcFunc.OnUserListGetter(conn, count, rec)
		if err != nil {
			ilogger.Error("UserDao - OnUserListGetter", ilogger.LogData{
				"err": err,
			})
		}

	}

	return
}
