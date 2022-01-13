package models

type Driver struct {
	ID        string `json:"id"`
	Fullname string `json:"fullname"`
	Phone  string `json:"phone"`
}

type CreateDriverBody struct {
	Fullname string `json:"fullname"`
	Phone  string `json:"phone"`
}


type Client struct {
	ID        string `json:"id"`
	Fullname string `json:"fullname"`
	Phone  string `json:"phone"`
}

type CreateClientBody struct {
	Fullname string `json:"fullname"`
	Phone  string `json:"phone"`
}

type FullOrder struct {
	Id                   string   `json:"id"`
	Status               string   `json:"status"`
	Driver               Driver  `json:"driver"`
	Client               Client  `json:"client"`
}

type Order struct {
	Id                   string   `json:"id"`
	Status               string   `json:"status"`
	DriverId             int64    `json:"driverId"`
	ClientId             int64    `json:"clientId"`
}

type ListOrders struct {
	Orders []Order  `json:"orders"`
	Count  int64     `json:"count"`
}

type UpdateOrder struct {
	Status               string   `json:"status"`
}

type CreateOrder struct {
	Status               string   `json:"status"`
	DriverId             string    `json:"driverId"`
	ClientId             string    `json:"clientId"`
}
