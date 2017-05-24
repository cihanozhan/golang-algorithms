package main
 
import (
    "fmt"
    "strings"
)
 
func main() {
    result1 := RemoveDuplicateCharacters("dijibil")
    result2 := RemoveDuplicateCharacters("cihanözhan")
    fmt.Println(result1)
    fmt.Println(result2)
}
 
func RemoveDuplicateCharacters(input string) string {
 
    var data = ""
    var result = ""
 
    for _, val := range input {
        var _val = string(val)
        if strings.Index(data, _val) == -1 {
            data += _val
            result += _val
        }
    }
 
    return result
}
