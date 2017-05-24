package main
 
import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
)
 
type SortRunes []rune
 
func main() {
 
    var pattern = "*********"
 
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("First Word: ")
    word1, _ := reader.ReadString('\n')
    fmt.Print("Second Word: ")
    word2, _ := reader.ReadString('\n')
 
    word1 = SortString(strings.ToLower(word1))
    word2 = SortString(strings.ToLower(word2))
 
    // fmt.Print("Word1 : " + word1)
    // fmt.Print("Word2 : " + word2)
 
    if word1 == word2 {
        fmt.Println(pattern)
        fmt.Println("Bu iki kelime Anagram'dır")
        fmt.Println(pattern)
    } else {
        fmt.Println(pattern)
        fmt.Println("Bu iki kelime Anagram değildir")
        fmt.Println(pattern)
    }
}
 
func SortString(s string) string {
    r := []rune(s)
    sort.Sort(SortRunes(r))
    return string(r)
}
 
func (s SortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}
 
func (s SortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
 
func (s SortRunes) Len() int {
    return len(s)
}
