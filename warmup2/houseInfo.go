import (
    "fmt"
    "bufio"
    )

type House struct {
    numberOfRooms int
    city string
    address string
    price int
}

func main(){
    reader := bufio.NewReader(os.Stdin)

    var ahouse House

    fmt.Print("Number of rooms in house: ")
    fmt.Print("Street Address: ")
    fmt.Print("City: ")
    fmt.Print("Selling Price: ")
}

// in another class and will do this when I can focus fully on it for the 30 minutes
// less about 5 mins for what I've already done
