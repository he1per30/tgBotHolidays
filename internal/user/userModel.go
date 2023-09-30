package user

type Users struct {
	Lastname    string `json:"last_name"`
	FirstName   string `json:"first_name"`
	MiddleName  string `json:"middle_name"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	City        string `json:"city"`
	Role        string `json:"role"`
	TgUserName  string `json:"telegram_id"`
	Id          int    `json:"id"`
}
