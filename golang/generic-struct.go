package main

import (
    "generic-struct/blog"
    "generic-struct/cache"
)

func main() {
    category := blog.Category{
        ID: 1,
        Name: "Go Generics",
        Slug: "go-generics",
    }

    cc := cache.New[blog.Category]()
    cc.Set(category.Slug, category)

    post := blog.Post{
        ID: 1,
        Categories: []blog.Category{
            {ID: 1, Name: "Go Generics", Slug: "go-generics"},
        },
        Title: "Generics in Golang structs",
        Text: "Here go's the text",
        Slug: "generics-in-golang-structs",
    }

    cp := cache.New[blog.Post]()
    cp.Set(post.Slug, post)
}

type Category struct {
    ID int32
    Name string
    Slug string
}

type Post struct {
    ID int32
    Categories []Category
    Title string
    Text string
    Slug string
}

type cacheable interface {
    blog.Category | blog.Post
}

type cache[T cacheable] struct {
    data map[string]T
}

func New[T cacheable]() *cache[T] {
    c := cache[T]{}
    c.data = make(map[string]T)
    return &c
}

func (c *cache[T]) Set(key string, value T) {
    c.data[key] = value
}

func (c *cache[T]) Get(key string) (v T) {
    if v, ok := c.data[key]; ok {
        return v
    }
    return
}
