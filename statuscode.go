package main
 
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"net/http"
)
 
func main() {
	file, err := os.Open("input.txt")
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
 
	file.Close()
 
	for _, eachline := range txtlines {
		{

    resp, err := http.Get(eachline)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(eachline)
    fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
}
}
