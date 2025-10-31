package main

func main() {
	defer func() {
		println("程序结束")
	}()

	// select {}
	// for {}
	// <-make(chan struct{})
	// go func() {
	// 	for {

	// 		if time.Now().Unix()%7 == 0 {
	// 			return
	// 		} else {
	// 			println(time.Now().Format(time.DateTime))
	// 		}

	// 		time.Sleep(1 * time.Second)
	// 	}

	// }()
	// go func() {}()
	// select {}
	// <-make(chan struct{})

}
