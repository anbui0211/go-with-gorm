package responsemodel

type User struct {
	UserID   string  `json:"user_id"`
	UserName string  `json:"user_name"`
	Age      int     `json:"age"`
	Gender   int     `json:"gender"`
	UserType int     `json:"user_type"`
	Status   string  `json:"status"`
	Address  Address `json:"address"`
}

type Address struct {
	FullAddress string `json:"full_address"`
	Province    string `json:"province"`
	District    string `json:"district"`
	Ward        string `json:"ward"`
}
