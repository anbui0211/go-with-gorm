package icontext

import (
	"context"

	"gorm.io/gorm"
)

const (
	DataBaseContext = "databaseKey"
	UserKeyContext  = "userKey"
)

func Connection(c context.Context) *gorm.DB {
	return c.Value(DataBaseContext).(*gorm.DB)
}

func GetCurrenUser(c context.Context) interface{} {
	return c.Value(UserKeyContext)
}
