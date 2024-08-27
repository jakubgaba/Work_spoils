package backend

type Product struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type Result struct {
	Id int `json:"id"`
}
