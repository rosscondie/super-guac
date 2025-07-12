# 📸 Photo Blog (Go + Fiber Backend) 🥑

This is a personal project to **learn Go (Golang)** by building a photo blog with a backend API and later a React frontend.

---

## 🧠 Why This Project?

I'm using this project to:

- Learn backend development with Go
- Understand how APIs are structured and built
- Work with Markdown and serve blog content dynamically
- Eventually display and manage my photography on a modern web stack
- Practice clean code structure, version control, and deployment habits

---

## ✅ What’s Been Done So Far

- [X] Initialized a Go project using [Fiber](https://github.com/gofiber/fiber)
- [X] Set up basic routes for a blog API
- [X] Created `GET /api/posts` to list all blog posts (from Markdown files)
- [X] Created `GET /api/posts/:slug` to return the full content of a post
- [X] Parsed Markdown into HTML using [`goldmark`](https://github.com/yuin/goldmark)
- [X] Installed [`air`](https://github.com/air-verse/air) for live-reloading during development

---

## 📁 Project Structure

```txt
.
├── content
│   ├── first-post.md
│   └── second-post.md
├── data
│   ├── photos.go
│   ├── post_reader.go
│   └── posts.go
├── go.mod
├── go.sum
├── handlers
│   ├── photos.go
│   └── posts.go
├── main.go
├── public
│   ├── images
│   └── index.html
├── README.md
└── routes
    ├── photos.go
    └── posts.go
```

## 🔜 TODO

- [X] Add `GET /api/photos` to return image metadata for a photo gallery
- [X] Serve images via `GET /images/:filename`
- [X] Use frontmatter in Markdown for post titles, dates, and tags
- [X] Start building a React frontend that fetches from this API
- [ ] Implement a basic admin dashboard to upload new posts/photos (eventually)
- [ ] Deploy the site publicly for viewing and learning showcase

---
