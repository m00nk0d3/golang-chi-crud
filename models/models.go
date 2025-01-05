package models

type Book struct {
	ID               string `json:"id"`
	Title            string `json:"name"`
	Author           string `json:"author"`
	PublishedDate    string `json:"published_date"`
	OriginalLanguage string `json:"original_language"`
}

var Books = []*Book{
	{
		ID:               "1",
		Title:            "The Lord of the Rings",
		Author:           "J.R.R Tolkien",
		PublishedDate:    "1954-07-29",
		OriginalLanguage: "English",
	},
	{
		ID:               "2",
		Author:           "J.R.R Tolkien",
		Title:            "The Hobbit",
		PublishedDate:    "1937-09-21",
		OriginalLanguage: "English",
	},
	{
		ID:               "3",
		Title:            "Harry Potter and the Sorcerer's Stone",
		Author:           "J.K. Rowling",
		PublishedDate:    "1997-06-26",
		OriginalLanguage: "English",
	},
	{
		ID:               "4",
		Title:            "To Kill a Mockingbird",
		Author:           "Harper Lee",
		PublishedDate:    "1960-07-11",
		OriginalLanguage: "English",
	},
	{
		ID:               "5",
		Title:            "1984",
		Author:           "George Orwell",
		PublishedDate:    "1949-06-08",
		OriginalLanguage: "English",
	},
	{
		ID:               "6",
		Title:            "Pride and Prejudice",
		Author:           "Jane Austen",
		PublishedDate:    "1813-01-28",
		OriginalLanguage: "English",
	},
	{
		ID:               "7",
		Title:            "The Great Gatsby",
		Author:           "F. Scott Fitzgerald",
		PublishedDate:    "1925-04-10",
		OriginalLanguage: "English",
	},
	{
		ID:               "8",
		Title:            "Moby Dick",
		Author:           "Herman Melville",
		PublishedDate:    "1851-10-18",
		OriginalLanguage: "English",
	},
}

func ListBooks() []*Book {
	return Books
}

func GetBook(id string) *Book {
	for _, b := range Books {
		if b.ID == id {
			return b
		}
	}
	return nil
}

func StoreBook(book Book) {
	Books = append(Books, &book)
}

func DeleteBook(id string) *Book {
	for i, b := range Books {
		if b.ID == id {
			Books = append(Books[:i], Books[i+1:]...)
			return &Book{}
		}
	}
	return nil
}

func UpdateBook(id string, bookUpdate Book) *Book {
	for i, b := range Books {
		if b.ID == id {
			Books[i] = &bookUpdate
			return b
		}
	}
	return nil
}
