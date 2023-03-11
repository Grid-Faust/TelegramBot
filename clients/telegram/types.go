package telegram

type UpdatesResponse struct {
	Ok 		bool 
	Result 	[]Update
}

type Update struct {
	ID 		int			`json:"update_id"`
	Message string 		`json:"message"`
}