package model

type User struct {
	userID   int
	userName string
	email    string
	mobile   string
}

func NewUser(userID int, userName, email, mobile string) User {
	return User{
		userID:   userID,
		userName: userName,
		email:    email,
		mobile:   mobile,
	}
}

func (u *User) SetUserID(userID int) {
	u.userID = userID
}

func (u User) GetUserID() int {
	return u.userID
}

func (u *User) SetUserName(userName string) {
	u.userName = userName
}

func (u User) GetUserName() string {
	return u.userName
}

func (u *User) SetUserMobile(userMobile string) {
	u.mobile = userMobile
}

func (u User) GetUserMobile() string {
	return u.mobile
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u User) GetUserEmail() string {
	return u.email
}
