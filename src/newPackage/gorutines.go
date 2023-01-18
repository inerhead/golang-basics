package myPackage

import (
	"fmt"
	"net/http"
	"sync"
	"time"
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

func ShowStatusWebSiteNet() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org2",
		"http://amazon.com",
	}

	var wg sync.WaitGroup
	for _, v := range links {
		wg.Add(1)
		go get(v, &wg)
	}

	wg.Wait()
}

func ShowStatusWebSiteNetWithChannel() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	ch := make(chan string)

	for _, v := range links {

		go getChan(v, ch)
	}

	// Option 1
	/*for {
		go getChan(<-ch, ch)
	}*/

	// Option 2
	for l := range ch {
		go func(u string) {
			time.Sleep(5 * time.Second)
			getChan(u, ch)
		}(l)
	}

}

func getChan(url string, ch chan<- string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("%q is Down\n", url)
		ch <- url
		return
	}
	ch <- url
	fmt.Printf("%q is Up\n", url)
}

func get(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("%q is Down\n", url)
		return
	}

	fmt.Printf("%q is Up\n", url)
}
