// package main

// import (
//     "fmt"
//     "net/http"
//     "io/ioutil"
// )

// func main () {
//     resp, err := http.Get("https://www.example.com/")
// 	if err != nil {
// 		print(err)
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		print(err)
// 	}
// 	fmt.Print(string(body))
// }
package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	for z := 1.0; z != x; z++{
		if z -= (z*z - x) / (2*z); z*z != x {
			fmt.Println(z)
		}
        return z
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
