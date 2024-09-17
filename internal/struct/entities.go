package entities

// Структура для получения данных из запроса и для записи в неё данных из БД
type Building struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	City   string `json:"city"`
	Year   int    `json:"year"`
	Floors int    `json:"floors"`
}
