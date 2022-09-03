package crud

var movies []Movie

func InitData() {
	movies = append(movies, Movie{ID: "1", Isbn: "43", Title: "Scary",
		Director: &Director{Firstname: "John", Lastname: "Awesome"}})
	movies = append(movies, Movie{ID: "2", Isbn: "777", Title: "Funny",
		Director: &Director{Firstname: "Joe", Lastname: "Haha"}})
}
