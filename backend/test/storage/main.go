package main

import (
	"crypto/md5"
	"crypto/sha512"
	"fmt"
	"my-app/backend/pkg/storage"
	"os"
	"strings"
)

func main() {
	s := storage.New()
	added, _ := s.AddPaths("C:\\Users\\jinya\\Desktop\\my-app\\backend\\test\\storage\\tmp")
	u, _ := s.GetMountpointUsage()
	println("Added:", added)
	println("Availables:", strings.Join(u.AvailableMountPoints(), " | "))
	println("Total free:", u.TotalFree())
	println("Total used:", u.TotalUsed())
	println("Total calc:", u.TotalFree()+u.TotalUsed())
	println("Total size:", u.TotalSize())
	fmt.Printf("Total used (%%): %.2f %%\n", u.TotalUsedPercent())
	//fmt.Printf("Pick a path with %d size: %s", 409640964096, u.PickAPath(409640964096))

	file, _ := os.Open("C:\\Users\\jinya\\Desktop\\my-app\\backend\\test\\storage\\test.zip")
	fileInfo, _ := file.Stat()

	buffer := make([]byte, fileInfo.Size())
	n, _ := file.Read(buffer)
	checksum := fmt.Sprintf("%x-%x-%d", md5.Sum(buffer), sha512.Sum512(buffer), n)
	println("test.zip checksum =", checksum)

	file.Seek(0, 0)

	verifyMap := make(map[string]bool)

	size := 0
	buffer = make([]byte, 4096)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			break
		}

		temp := buffer[0:n]

		_, path, _ := s.Cache(checksum+".zip", temp, int64(size), int64(size+n), fileInfo.Size(), true)
		verifyMap[path] = false

		size += n
	}

	ok, paths, _ := s.VerifyChecksum(checksum+".zip", true, checksum)
	for _, p := range paths {
		println(p)
		verifyMap[p] = true
	}
	for k, v := range verifyMap {
		if !v {
			println(k, "=", v)
		}
	}

	if ok {
		ok, path, _ := s.Persist(checksum+".zip", paths, fileInfo.Size())
		println(ok, path)

		s.ClearCache(checksum + ".zip")

		ok, _, _ = s.VerifyChecksum(checksum+".zip", false, checksum)
		if ok {
			println("Persist file checksum ok")
		}

		size = 0
		var persistData []byte
		for int64(size) <= fileInfo.Size() {
			temp, _ := s.Load(checksum+".zip", int64(size), int64(size+4096))
			persistData = append(persistData, temp...)

			size += 4096
		}
		persistChecksum := fmt.Sprintf("%x-%x-%d", md5.Sum(persistData), sha512.Sum512(persistData), len(persistData))
		println("Persist Checksum =", persistChecksum)
		if persistChecksum == checksum {
			println("Persist file loaded successfully")
		}
	}
}
