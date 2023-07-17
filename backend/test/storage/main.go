package main

import (
	"fmt"
	"my-app/backend/pkg/storage"
	"strings"
)

func main() {
	s := storage.New()
	added, _ := s.AddPaths("C:/", "D:/")
	u, _ := s.GetMountpointUsage()
	println("Added:", added)
	println("Availables:", strings.Join(u.AvailableMountPoints(), " | "))
	println("Total free:", u.TotalFree())
	println("Total used:", u.TotalUsed())
	println("Total calc:", u.TotalFree()+u.TotalUsed())
	println("Total size:", u.TotalSize())
	fmt.Printf("Total used (%%): %.2f %%", u.TotalUsedPercent())
}
