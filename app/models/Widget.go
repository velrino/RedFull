package models

type Widget struct  {
	ID        uint   `json:"id"`
	Color string `json:"color"`	
	Inventory  int `json:"inventory"`
	Name  string `json:"name "`
	Price  string `json:"price"`
}