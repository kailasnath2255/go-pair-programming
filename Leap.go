
package main

import "fmt"

func main() {
var date int
fmt.Println("Enter the Year to Check")
fmt.Scan(&date)

if (date%4 == 0 && date%100 != 0) || (date%400 == 0) {
fmt.Println("The year is leap year")
}else {
fmt.Println("The year is not leap year")
}	
}