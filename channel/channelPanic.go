package channel

/*
panic
*/
func closeNilChannel() {
	var channel chan int
	close(channel)
}

func closeChannel() chan int {
	channel := make(chan int)
	close(channel)
	return channel
}

/*
deadlock
*/
func readNilChannel() {
	var channel chan int
	<-channel
}

/*
deadlock
*/
func writeNilChannel() {
	var channel chan int
	channel <- 1
}
