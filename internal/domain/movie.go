package domain

type Movie struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Year     int    `json:"year"`
	Country  string `json:"country"`
}

type EditMovieInfo struct {
	Title    *string `json:"title"`
	Director *string `json:"director"`
	Year     *int    `json:"year"`
	Country  *string `json:"country"`
}
