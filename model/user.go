package model

// User entity of user Entry
type User struct {
	UserId      uint64 `json:"userId" form:"userId" uri:"userId"`
	Username    string `json:"username" form:"username" uri:"username"`
	RealName    string `json:"realName" form:"realName" uri:"realName"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber" uri:"phoneNumber"`
	Email       string `json:"email" form:"email" uri:"email"`
	Password    string `json:"password" form:"password" uri:"password"`
	Birthday    string `json:"birthday" form:"birthday" uri:"birthday"`
	Gender      int8   `json:"gender" form:"gender" uri:"gender"`
}
