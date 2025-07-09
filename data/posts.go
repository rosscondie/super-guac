package data

type Post struct {
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Summary string `json:"summary"`
}

func GetAllPosts() []Post {
	return []Post{
		{Title: "First Post", Slug: "first-post", Summary: "This is the first post"},
		{Title: "Second Post", Slug: "second-post", Summary: "This is the second post"},
	}
}
