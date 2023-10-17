package dto

type T struct {
	Table struct {
		Type  string     `json:"type"`
		Name  string     `json:"name"`
		Week  int        `json:"week"`
		Group string     `json:"group"`
		Table [][]string `json:"table"`
		Link  string     `json:"link"`
	} `json:"table"`
	Weeks []int `json:"weeks"`
}

type YourModel struct {
	ID   string `gorm:"primaryKey" json:"group_name"`
	Week int    `gorm:"primaryKey" json:"week"`
	DATA string `gorm:"type:json"`
}
