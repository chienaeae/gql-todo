package users

type UserProfile struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func NewProfile(user User) *UserProfile {
	return &UserProfile{
		ID:       user.ID,
		Username: user.Username,
	}
}
