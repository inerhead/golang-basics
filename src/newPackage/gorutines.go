package myPackage

import (
	"fmt"
	"sync"
)

func print(msg string) {
	fmt.Println(msg)
}

func ShowGoRoutinesWithWG() {

	mapi := make(map[int]string)
	mapi[1] = "Hola"
	mapi[2] = "Mundo"

	//wg := sync.WaitGroup
	var wg sync.WaitGroup
	fmt.Printf("\n\n")
	fmt.Println("GO ROUTINE WITH WAIT GROUP")
	for _, v := range mapi {
		wg.Add(1)
		fmt.Println("Intentando PRint: " + v)
		go func(m string, w *sync.WaitGroup) {
			defer w.Done()
			print(m)

		}(v, &wg)

	}
	fmt.Printf("\n\n")
	wg.Wait()

}

func addMessageToChannel(m string, c chan<- string) {
	c <- m
}

func ShowGoRoutinesWithChannels() {

	const Filaauto int = 3

	c := make(chan string, 1)
	fmt.Printf("\n\n")
	fmt.Println("GO ROUTINE WITH CHANNEL")
	addMessageToChannel("HEYYYY", c)

	fmt.Println(<-c)
	fmt.Printf("\n\n")

}
