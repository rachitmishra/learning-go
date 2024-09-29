package types

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	SKU         string    `json:"sku"`
	CreatedOn   time.Time `json:"-"`
	UpdatedOn   time.Time `json:"-"`
	DeletedOn   time.Time `json:"-"`
	Roast       string    `json:"roast"`
	Origin      string    `json:"origin"`
	Size        string    `json:"size"`
	ImageURL    string    `json:"image_url"`
	IsAvailable bool      `json:"is_available"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProductList() Products {
	return generateCoffeeList()
}

func generateCoffeeList() Products {
	now := time.Now()

	return []*Product{
		{
			ID:          1,
			Name:        "Espresso",
			Description: "A concentrated shot of coffee with a bold flavor.",
			Price:       3.50,
			SKU:         "COF-001",
			CreatedOn:   now,
			UpdatedOn:   now,
			DeletedOn:   time.Time{},
			Roast:       "Dark",
			Origin:      "Italy",
			Size:        "Small",
			ImageURL:    "https://example.com/espresso.jpg",
			IsAvailable: true,
		},
		{
			ID:          2,
			Name:        "Latte",
			Description: "Espresso with steamed milk and foam.",
			Price:       4.00,
			SKU:         "COF-002",
			CreatedOn:   now,
			UpdatedOn:   now,
			DeletedOn:   time.Time{},
			Roast:       "Medium",
			Origin:      "Brazil",
			Size:        "Medium",
			ImageURL:    "https://example.com/latte.jpg",
			IsAvailable: true,
		},
		{
			ID:          3,
			Name:        "Cappuccino",
			Description: "Espresso with steamed milk and foam, topped with cocoa powder.",
			Price:       4.50,
			SKU:         "COF-003",
			CreatedOn:   now,
			UpdatedOn:   now,
			DeletedOn:   time.Time{},
			Roast:       "Light",
			Origin:      "Ethiopia",
			Size:        "Large",
			ImageURL:    "https://example.com/cappuccino.jpg",
			IsAvailable: true,
		},
		&Product{
			ID:          4,
			Name:        "Americano",
			Description: "Espresso diluted with hot water.",
			Price:       3.00,
			SKU:         "COF-004",
			CreatedOn:   now,
			UpdatedOn:   now,
			DeletedOn:   time.Time{},
			Roast:       "Dark",
			Origin:      "Colombia",
			Size:        "Large",
			ImageURL:    "https://example.com/americano.jpg",
			IsAvailable: true,
		},
		&Product{
			ID:          5,
			Name:        "Mocha",
			Description: "Espresso with hot chocolate and milk.",
			Price:       4.50,
			SKU:         "COF-005",
			CreatedOn:   now,
			UpdatedOn:   now,
			DeletedOn:   time.Time{},
			Roast:       "Medium",
			Origin:      "Peru",
			Size:        "Medium",
			ImageURL:    "https://example.com/mocha.jpg",
			IsAvailable: true,
		},
		&Product{
			ID:          6,
			Name:        "Macchiato",
			Description: "Espresso with a small amount of foam.",
			Price:       3.00,
			SKU:         "COF-006",
			CreatedOn:   now,
			UpdatedOn:   now,
			DeletedOn:   time.Time{},
			Roast:       "Dark",
			Origin:      "Vietnam",
			Size:        "Small",
			ImageURL:    "https://example.com/macchiato.jpg",
			IsAvailable: true,
		},
		&Product{
			ID:          7,
			Name:        "Flat White",
			Description: "Espresso with steamed milk and a thin layer of foam.",
			Price:       3.50,
			SKU:         "COF-007",
			CreatedOn:   now,
			UpdatedOn:   now,
			DeletedOn:   time.Time{},
			Roast:       "Light",
			Origin:      "New Guinea",
			Size:        "Medium",
			ImageURL:    "https://example.com/flat_white.jpg",
			IsAvailable: true,
		},
		&Product{
			ID:          8,
			Name:        "Irish Coffee",
			Description: "Espresso with Irish whiskey, sugar, and whipped cream.",
			Price:       5.00,
			SKU:         "COF-008",
			CreatedOn:   now,
			UpdatedOn:   now,
			DeletedOn:   time.Time{},
			Roast:       "Dark",
			Origin:      "Guatemala",
			Size:        "Large",
			ImageURL:    "https://example.com/irish_coffee.jpg",
			IsAvailable: true,
		},
		&Product{
			ID:          9,
			Name:        "Cold Brew",
			Description: "Cold-brewed coffee with a smooth flavor.",
			Price:       4.00,
			SKU:         "COF-009",
			CreatedOn:   now,
			UpdatedOn:   now,
			DeletedOn:   time.Time{},
			Roast:       "Medium",
			Origin:      "Costa Rica",
			Size:        "Large",
			ImageURL:    "https://example.com/cold_brew.jpg",
			IsAvailable: true,
		},
		&Product{
			ID:          10,
			Name:        "Affogato",
			Description: "Espresso poured over a scoop of ice cream.",
			Price:       4.50,
			SKU:         "COF-010",
			CreatedOn:   now,
			UpdatedOn:   now,
			DeletedOn:   time.Time{},
			Roast:       "Dark",
			Origin:      "Indonesia",
			Size:        "Small",
			ImageURL:    "https://example.com/affogato.jpg",
			IsAvailable: true,
		},
	}
}
