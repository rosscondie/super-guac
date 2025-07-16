package data

type Post struct {
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Summary string `json:"summary"`
	Date    string `json:"date"`
}

type frontmatter struct {
	Title   string `yaml:"title"`
	Summary string `yaml:"summary"`
	Date    string `yaml:"date"`
}

type PostContent struct {
	Title string `json:"title"`
	Date  string `json:"date"`
	HTML  string `json:"html"`
}

type Photo struct {
	Filename string `json:"filename"`
	URL      string `json:"url"`
	Size     int64  `json:"size"` // bytes
}

type Album struct {
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Cover string `json:"cover,omitempty"`
}
