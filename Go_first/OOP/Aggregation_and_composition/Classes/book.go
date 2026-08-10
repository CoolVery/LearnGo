package classes

func NewBook(title, author string, year int, genre string) *Book {
	newBook := Book{
		Title: title,
		Author: author,
		Year: year,
		Genre: genre,
	}
	return &newBook
}
type Book struct {
	Title string
	Author string
	Year int
	Genre string
}

