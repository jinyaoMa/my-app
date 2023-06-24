package main

import "my-app/backend/pkg/storage"

func main() {
	s := storage.New()
	added := s.AddPaths("C:/", "C:/Users/Administrator/Downloads")
	println("Added:", added)
	println("Total avil:", s.TotalAvailable())
	println("Total size:", s.TotalSize())
}
