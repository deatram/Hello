package main

import (
	"fmt"
	"math"
	"strings"
	"net/http"
	"html/template"
)
type Page struct {
	Name string
}

func main() {



	var (
		x         float64 = 16.0
		y                 = 2.8
		a         bool    = true
		b         bool    = false
		division  int
		square    float64 = squareroot(x)
		g, h, k           = 3, "dit", "00dfd"
		greetings         = "hello"
	)
	/*var ret int
	ret = max(20,40)*/

	fmt.Println(x)
	fmt.Printf("x is of type %T \n", x)
	fmt.Printf("y is of type %T \n", y)
	fmt.Printf("g is of type %T \n", g)
	fmt.Printf("h is of type %T \n", h)
	fmt.Printf("k is of type %T \n", k)
	fmt.Println("Hello, World! \n")
	var (
		LENGTH int = 8
		WIDTH  int = 0
	)
	if WIDTH != 0 && LENGTH != 0 {
		division = LENGTH / WIDTH
		fmt.Printf("division is %d \n", division)
	} else {
		fmt.Printf("cannot divide number by zero \n")
	}
	if a || b {
		fmt.Printf("condition is true \n")
	} else {
		fmt.Printf("Condition is not true \n")
	}
	var (
		marks int = 90
		grade string
	)
	switch marks {
	case 90:
		grade = "A"
	default:
		grade = "Failed"
	}
	fmt.Printf("Grade is %s \n", grade)
	fmt.Printf("maximum value is %d \n", max(50, 20))
	fmt.Printf("Square root of 16 is %f \n", square)
	var (
		nextsequence = getsequence()
	)
	fmt.Println(nextsequence())
	fmt.Println(nextsequence())
	fmt.Printf("function closure is of type %T \n", nextsequence())
	fmt.Printf("%d \n", nextsequence())
	fmt.Printf("%s \n", greetings)
	fmt.Printf("quoted string: %+q \n", greetings)
	fmt.Printf("hexa bytes ")

	fmt.Printf("%x ", greetings[0])
	fmt.Printf("%x ", greetings[1])
	fmt.Printf("%x ", greetings[2])
	fmt.Printf("%x ", greetings[3])
	fmt.Printf("%x \n", greetings[4])
	var firstname Page
	firstname.Name = "Dit"
	fmt.Println(firstname.Name)
	/*for ( z := 0; z < len(greetings); z++ ) {
		fmt.Printf("%x ", greetings[i])
	}*/
	greet := []string{"oi", "vp"}
	fmt.Println(strings.Join(greet, " "))
	var balance = []int{1, 2, 3}
	fmt.Printf("%d \n", balance[2])
	templates := template.Must(template.ParseFiles("templates/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := Page{Name: "Dit"}
		if name :=r.FormValue("name"); name != "" {
			p.Name = name
		}
		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		//fmt.Fprintf(w, "Hello, Go Web Development")

	})
	fmt.Println(http.ListenAndServe(":9090", nil))


}
func max(num1 int, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
func squareroot(x float64) float64 {
	return math.Sqrt(x)
}

//function closure
func getsequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
