package myPackage

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type customLogOut struct{}

func (customLogOut) Write(b []byte) (n int, err error) {

	fmt.Println("Bytes read IS : ", len(b))
	fmt.Println(string(b))
	return len(b), nil
}

type customBody struct{}

func (customBody) Read(b []byte) (n int, err error) {
	fmt.Println("************************************ LEIDO WITH ***", len(b))
	os.Exit(1)
	return len(b), nil
}

func ShowMessageFromHttp() {

	//io.Copy(os.Stdout.w)
	resp, err := http.Get("http://google.com")
	if err != nil {
		os.Exit(1)
	}

	io.Copy(customLogOut{}, resp.Body)
}

func ShowMessageFromIoCustom() {
	//io.Copy(customLogOut{}, customBody{})
	b, e := io.ReadAll(customBody{})
	if e != nil {
		os.Exit(1)
	}

	fmt.Println("Byte Slice is: ", b)
}
