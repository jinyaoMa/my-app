package main

import "majinyao.cn/my-app/backend/internal/app"

func main() {
	_ = app.H3S
	<-make(chan struct{}, 1)
}
