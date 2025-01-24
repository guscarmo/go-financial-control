package models

type Transacao struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Amount      float64 `json:"amount"`
	Typ         string  `json:"typ"`
	Payment     string  `json:"payment"`
	Obs         string  `json:"obs"`
	Date        string  `json:"date"`
}
