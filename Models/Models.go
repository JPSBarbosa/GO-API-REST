package Models

type Products struct {
	ID         uint    `gorm:"primary_key" json:"id"`
	Nome       string  `gorm:"type:varchar(255);not null" json:"nome"`
	Descricao  string  `gorm:"type:text" json:"descricao"`
	Preco      float64 `gorm:"type:decimal(10,2);not null" json:"preco"`
	Quantidade int     `gorm:"not null" json:"quantidade"`
}
