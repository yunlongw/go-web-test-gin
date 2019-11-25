package main

import (
	"fmt"
	"go-web-test-gin/base/retriever/mock"
	"go-web-test-gin/base/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

const url = "http://www.imooc.com"

func post(poster Poster)  {
	poster.Post(url, map[string]string{
		"name": "ccmouse",
		"course" : "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents" : "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	r = &mock.Retriever{
		Contents: "this is a Retriever",
	}
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.5",
		TimeOut:   time.Minute,
	}
	inspect(r)
	//fmt.Println(download(r))

	// Type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	// Type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	mockRetriever := mock.Retriever{Contents: "this is a fake imooc.com"}
	r = &mockRetriever
	fmt.Println("Try a session with mockRetriever")
	fmt.Println(session(&mockRetriever))

}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > Type:%T Value:%v\n", r, r)
	fmt.Print(" > Type switch: ")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
