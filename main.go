package main

import (
	"log"
	"net/http"
)

func NewInMemoryPostStore() *InMemoryPostStore {
	return &InMemoryPostStore{
		map[int]Post{
			1: Post{1, "title", "text"},
		},
	}
}

type InMemoryPostStore struct {
	store map[int]Post
}

func (i *InMemoryPostStore) GetAllPosts() []Post {
	posts := make([]Post, 0, len(i.store))
	for _, post := range i.store {
		posts = append(posts, post)
	}
	return posts
}

func (i *InMemoryPostStore) GetPostByID(id int) (Post, error) {
	return i.store[id], nil
}

func main() {
	server := &PostServer{NewInMemoryPostStore()}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}