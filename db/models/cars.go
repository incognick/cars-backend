package models

type Car struct {
	ID         string `json:"id"`
	Make       string `validate:"required" json:"make"`
	Model      string `validate:"required" json:"model"`
	Package    string `validate:"required" json:"package"`
	Color      string `validate:"required" json:"color"`
	Year       int    `validate:"gte=1900,lte=2100" json:"year"`
	Category   string `validate:"oneof=Truck Sedan SUV" json:"category"`
	Mileage    int    `validate:"required" json:"mileage"`
	PriceCents int    `validate:"required" json:"price_cents"`
}
