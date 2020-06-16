package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    var this_year int
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter your year of birth: ")
    // supposed to get input from terminal
    yob, _ := reader.ReadString('\n')

    fmt.Print("Enter your age: ")
    // also supposed to get input from terminal
    age, _ := reader.ReadString('\n')

    // remove newlines
    yob = strings.Replace(yob,"\n","",-1)
    age = strings.Replace(age,"\n","",-1)

    // convert strings to ints
    num1, e := strconv.Atoi(yob)
        if e != nil {
    	fmt.Println("conversion error:", yob)
        }
    num2, e := strconv.Atoi(age)
        if e != nil {
    	fmt.Println("conversion error:", age)
        }

    this_year = num2 + num1
    fmt.Println("The current year is ")
    fmt.Println(this_year)

}

// strconv.Itoa(int_var) converst int to str
