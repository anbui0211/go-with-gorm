package dao

import (
	"errors"
	ilogger "gwg/internal/logger"
	"gwg/src/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserListGetterWith interface {
	OnUserListGetterWith(conn *gorm.DB, count int32, user entity.User, location entity.Location) error
}

type User struct{}

func (User) GetListWith(conn *gorm.DB, svcFunc UserListGetterWith) (err error) {
	rows, err := conn.Table("users as u ").
		Select("u.*, " +
			"l.location_id AS l_location_id, l.full_address AS l_full_address, " +
			"l.province AS l_province, l.district AS l_district, l.ward AS l_ward").
		Joins("LEFT JOIN locations AS l " +
			"ON u.location_id = l.location_id").
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
		var rec struct {
			User     entity.User     `gorm:"embedded"`
			Location entity.Location `gorm:"embedded;embeddedPrefix:l_"`
		}

		err = conn.ScanRows(rows, &rec)
		if err != nil {
			ilogger.Error("UserDao - ScanRows", ilogger.LogData{
				"err": err,
			})
		}

		err = svcFunc.OnUserListGetterWith(conn, count, rec.User, rec.Location)
		if err != nil {
			ilogger.Error("UserDao - OnUserListGetter", ilogger.LogData{
				"err": err,
			})
		}

	}

	return
}

func (User) Get(conn *gorm.DB, userID string) (rec entity.User, err error) {
	result := conn.
		Where(&entity.User{UserID: userID}, "UserID").
		Take(&rec)
	if result.Error != nil {
		ilogger.Error("UserDao - Get", ilogger.LogData{
			"userID": userID,
			"err":    err,
		})
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return rec, result.Error
		}
	}

	return rec, nil
}

func (User) Lock(conn *gorm.DB, userID string) (locked entity.User, err error) {
	result := conn.Clauses(clause.Locking{Strength: "UPDATE", Options: "NOWAIT"}).
		Where(&entity.User{UserID: userID}, "UserID").
		Take(&locked)
	if result.Error != nil {
		return locked, result.Error
	}

	return locked, nil
}
