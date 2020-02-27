package account

type User struct {
	Id 			int	`json:"uid"`
	Nickname 	string `json:"nickname"`
	Avatar		string `json:"avatar"`
}