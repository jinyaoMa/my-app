package main

import "my-app/backend/pkg/storage"

func main() {
	s := storage.New()
	added, _ := s.AddPaths("C:/", "D:/tmp", "D:/")
	println("Added:", added)
	println("Total avil:", s.TotalAvailable())
	println("Total size:", s.TotalSize())
}
