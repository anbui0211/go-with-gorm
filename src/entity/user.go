package entity

type User struct {
	UserID     string `gorm:"column:user_id"`
	UserName   string `gorm:"column:user_name"`
	Age        int    `gorm:"column:age"`
	Gender     int    `gorm:"column:gender"`
	UserType   int    `gorm:"column:user_type"`
	Status     string `gorm:"column:status"`
	LocationID string `gorm:"column:location_id"`
}

type Location struct {
	LocationID  string `gorm:"column:location_id"`
	FullAddress string `gorm:"column:full_address"`
	Province    string `gorm:"column:province"`
	District    string `gorm:"column:district"`
	Ward        string `gorm:"column:ward"`
}
