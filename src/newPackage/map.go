package myPackage

import "fmt"

func ShowMap() {

	// option 2
	m1 := map[string]string{
		"name": "Gustavo",
	}

	// Option 1
	var m map[string]string
	m = m1

	// Option 3
	m2 := make(map[int]string)
	m2[0] = "Example"

	fmt.Printf("\n\n")
	fmt.Println("MAPS")
	fmt.Println("Options: ", m1, m)

	fmt.Printf("\n\n")

}

func loop() {
	m := make(map[string]int)
	m["Gustavo"] = 39
	m["Maxi"] = 5
	m["Naty"] = 40

	for i, v := range m {
		fmt.Printf(
			"La edad de %s es %d \n", i, v)
	}
}

func checkValueMap(m map[string]int) {
	existe, err := m["Maxi"]
	fmt.Println(
		existe,
		err,
	)
}
