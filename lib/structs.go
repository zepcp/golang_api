package lib

type Book struct {
    ID     int     `json:"id"`
    Isbn   string  `json:"isbn"`
    Title  string  `json:"title"`
    Author *Author `json:"author"`
}

type Author struct {
    FirstName string `json:"firstname"`
    LastName  string `json:"lastname"`
}
