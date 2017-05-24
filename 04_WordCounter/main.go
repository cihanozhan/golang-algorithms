package main
 
import (
    "fmt"
    "strconv"
    "strings"
)
 
func main() {
    var result = Count("eskiden buralar hep dutluktu")
    fmt.Println("Toplam : " + strconv.Itoa(result))
}
 
func Count(input string) int {
 
    result := 0
    x := strings.TrimSpace(input)
 
    if x == "" {
        return 0
    }
 
    for strings.Contains(x, "  ") {
        x = strings.Replace(x, "  ", " ", 2)
    }
 
    for _, _ = range strings.Split(x, " ") {
        result++
    }
 
    return result
}
