package main

func main() {
	//for i := 0; i < 5; i++ {
	//	defer func() {
	//		println(i)
	//	}()
	//}

	for i := 0; i <= 5; i++ {
		defer func(num int) {
			println(num)
		}(i)
	}
}
