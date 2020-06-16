package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    var avgWt int
    reader := bufio.NewReader(os.Stdin)

    // get inputs from terminal
    fmt.Print("Enter first weight: ")
    w1, _ := reader.ReadString('\n')

    fmt.Print("Enter second weight: ")
    w2, _ := reader.ReadString('\n')

    fmt.Print("Enter third weight: ")
    w3, _ := reader.ReadString('\n')

    fmt.Print("Enter fourth weight: ")
    w4, _ := reader.ReadString('\n')

    fmt.Print("Enter fifth weight: ")
    w5, _ := reader.ReadString('\n')



    // remove newlines
    w1 = strings.Replace(w1,"\n","",-1)
    w2 = strings.Replace(w2,"\n","",-1)
    w3 = strings.Replace(w3,"\n","",-1)
    w4 = strings.Replace(w4,"\n","",-1)
    w5 = strings.Replace(w5,"\n","",-1)
    // NOTE: strings.TrimSuffix(wx,"\n") works also

    // convert strings to ints
    num1, e := strconv.Atoi(w1)
        if e != nil {
    	fmt.Println("conversion error:", w1)
        }
    num2, e := strconv.Atoi(w2)
        if e != nil {
    	fmt.Println("conversion error:", w2)
        }
    num3, e := strconv.Atoi(w3)
        if e != nil {
        fmt.Println("conversion error:", w3)
        }
    num4, e := strconv.Atoi(w4)
        if e != nil {
        fmt.Println("conversion error:", w4)
        }
    num5, e := strconv.Atoi(w5)
        if e != nil {
    	fmt.Println("conversion error:", w5)
        }

    avgWt = (num2 + num1 + num3 + num4 + num5)/5
    fmt.Println("The average weight is ")
    fmt.Println(avgWt)

}
