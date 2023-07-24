package todo

type Foodstuff struct {
	Id      int `json:"id"`
	Protein int `json:"protein"`
	Fat     int `json:"fat"`
	Carb    int `json:"carb"`
	Price   int `json:"price"`
}

type UpdateFoodstuffInput struct {
	Protein int `json:"protein"`
	Fat     int `json:"fat"`
	Carb    int `json:"carb"`
	Price   int `json:"price"`
}
