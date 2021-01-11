package main

func main() {
	c := make(chan bool)
	c <- true
	// <-c
	// 由于没有其他的 goroutine 可以从 c 接收数据或者发送数据到 c, 导致 main goroutine 永远阻塞
}
