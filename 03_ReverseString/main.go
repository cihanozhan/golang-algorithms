package main
 
import "fmt"
 
func main() {
    result := Reverse("cihan")
    fmt.Println(result)
}
 
func Reverse(input string) string {
    result := ""
    for i := len(input) - 1; i >= 0; i-- {
        result += string(input[i])
    }
    return result
}
