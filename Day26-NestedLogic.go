package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
  var actualDay int
  var actualMonth int
  var actualYear  int
  var expectedDay int
  var expectedMonth int
  var expectedYear int
  fmt.Scanln(&actualDay, &actualMonth, &actualYear)
  fmt.Scanln(&expectedDay, &expectedMonth, &expectedYear)
  monthsLate := actualMonth - expectedMonth;
  daysLate := actualDay - expectedDay;
  yearDiference := actualYear - expectedYear;

  if yearDiference > 0 {
    fmt.Println(10000)
  }else if monthsLate > 0 && yearDiference == 0 {
    fmt.Println(500*monthsLate)
  }else if daysLate > 0 && monthsLate == 0{
    fmt.Println(15*daysLate)
  }else{
    fmt.Println(0)
  }
  return
}
