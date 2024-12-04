package data

// Record представляет модель данных для элемента
type Record struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Genre  string  `json:"genre"`
	Price  float64 `json:"price"`
}

// DB Records
var Records = []Record{
	{ID: 1, Title: "Stayin Alive", Artist: "Bee Gees", Genre: "Rock", Price: 56.99},
	{ID: 2, Title: "My Heart Will Go On", Artist: "Celein Dilon", Genre: "Pop", Price: 60.79},
	{ID: 3, Title: "Lose Yourself", Artist: "Eminem", Genre: "Rap", Price: 20.99},
}
