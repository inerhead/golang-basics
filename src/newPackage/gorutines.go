package myPackage

import (
	"fmt"
	"sync"
)

func print(msg string) {
	fmt.Println(msg)
}

func MessGo() {

	mapi := make(map[int]string)
	mapi[1] = "Hola"
	mapi[2] = "Mundo"

	//wg := sync.WaitGroup
	var wg sync.WaitGroup

	for _, v := range mapi {
		wg.Add(1)
		fmt.Println("Intentando PRint: " + v)
		go func(m string, w *sync.WaitGroup) {
			defer w.Done()
			print(m)

		}(v, &wg)

	}
	wg.Wait()

}

func addMessageToChannel(m string, c chan<- string) {
	c <- m
}

func Channels() {

	const Filaauto int = 3

	c := make(chan string, 1)

	addMessageToChannel("HEYYYY", c)

	fmt.Println(<-c)

}
