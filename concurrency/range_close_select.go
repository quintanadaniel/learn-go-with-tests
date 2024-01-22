package concurrency

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func RangeCloseSelect() {
	c := make(chan string, 2)

	c <- "Message1"
	c <- "Message2"

	fmt.Println(len(c), cap(c))
	// Range and close
	close(c)

	for message := range c{
		fmt.Println("Using range to iterate channel messages:",message)
	}

	//Select
	email1 := make(chan string)
	email2 := make(chan string)
	fmt.Println("Using go routine with channels.")
	go message("mensaje1", email1)
	go message("message2", email2)

	for i:=0;i< 2; i++ {
		select {
		case m1 := <- email1:
			fmt.Println("Email recibido de email1:", m1)
		case m2 := <- email2:
			fmt.Println("Email recibido de email2:", m2)
		}
	}



}