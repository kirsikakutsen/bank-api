package response


type AuthDto struct {
	Token 	string		`json:"token"`
	Profile AccountDto	`json:"profile"`
}
