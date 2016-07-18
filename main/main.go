package main

import "fmt"

func main() {
  //variable decration & initialization
  var number1 int;
  number1 = 100;

  var number2 float64 = 20.0;
  number3  := 100.56;

  //mixed variable
  var var1, var2, var3 = 3 , 10.56, "foo";
  total1 := var2 + number3;

  /*this is comment*/
  fmt.Println(number1);
  fmt.Println(number2);
  fmt.Println(number3, var1, var2, var3);
  fmt.Println("Total :",total1);
  fmt.Printf("testing\n")
  fmt.Printf("Hello world\n");
}
