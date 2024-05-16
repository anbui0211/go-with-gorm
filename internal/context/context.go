package icontext

import (
	"context"

	"gorm.io/gorm"
)

const (
	DataBaseContext = "databaseKey"
)

func Connection(c context.Context) *gorm.DB {
	return c.Value(DataBaseContext).(*gorm.DB)
}
