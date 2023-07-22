package todo

type TodoList struct {
	Id 			int    `json:"id"`
	Title 		string `json:"title"`
	Description string `json:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoList struct {
	Id 			int    `json:"id"`
	Title 		string `json:"title"`
	Description string `json:"description"`
	Done 		bool   `json:"done"`
}

type UsersList struct {
	Id     int
	UserId int
	ItemId int
}
