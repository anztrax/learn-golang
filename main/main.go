package main

import (
  "fmt"
  "math"
)

//declare type
type Circle struct{
  x, y, radius float64
}

//define method for the Circle
func(c *Circle) area() float64{
  return math.Pi * c.radius * c.radius;
}

//end of the method
//====================================

func markSomething(){
  var grade string = "Z";
  var marks int = 90;

  switch(marks){
    case 90:
      grade = "A";
    case 80:
      grade = "B";
    case 50,60,70:
      grade = "C";
    default:
      grade = "D";
  }

  //switch without variable to check ?
  switch{
    //the variable is checked at case section
  case grade == "A":
    fmt.Println("Excellent");
  case grade == "B":
    fmt.Println("Well Done");
  case grade == "C":
    fmt.Println("You Passed");
  case grade == "D":
    fmt.Println("You Passed Too ");
  default:
    fmt.Println("Invalid grade");
  }

  fmt.Printf("Your grade = %s\n",grade);
}

func checkIntefaceOfVariable(){
  var x interface{}

  switch i := x.(type){
    case nil:
      /**
        %T <- a Go-syntax representation of the type of the value
        %#v	<- a Go-syntax representation of the value
      **/

      fmt.Printf("type of x : %T\n",i);
    case int:
      fmt.Println("x is int");
    case float64:
      fmt.Println("x is float64");
    case func(int) float64:
      fmt.Println("x is func(int) return float64");
    case bool, string:
      fmt.Println("x is bool or string");
    default:
      fmt.Println("unknown type");
  }
}

func tryLooping(){
  var b int = 15;
  var a int;

  //this is array declaration & initialization
  numbers := [6]int{1,2,3,4};

  //for loop section
  //this a scope is only at for statement
  for a := 0; a < 10; a++ {
    fmt.Printf("value of a : %d\n",a);
  }
  fmt.Println("==================================");

  for a < b{
    a++;
    fmt.Printf("value of a : %d\n",a);
  }
  fmt.Println("==================================");

  //place ( index, value )
  for i, x:= range numbers{
    fmt.Printf("the value of x = %d at %d\n",x,i);
  }
  fmt.Println("==================================");
  findPrimeNumber(50);

  var num1 int = 100;
  var num2 int = 200;
  maxResult := findMax(num1,num2);
  fmt.Println("==================================");
  fmt.Println("max result value : ",maxResult);

  fragment1 := "ananta";
  fragment2 := "andrew";

  swappedVal1, swappedVal2 := swap("ananta","andrew");
  fmt.Println("my name is : ", swappedVal1, swappedVal2);

  fmt.Println("before swap : ",fragment1, fragment2);
  swapCallByReference(&fragment1,&fragment2);
  fmt.Println("after swap : ",fragment1, fragment2);
}

//function return 2 values (call by value)
func swap(x,y string)(string,string){
  return y,x;
}

func swapCallByReference(x *string, y*string){
  var temp string;
  temp = *x;
  *x = *y;
  *y = temp;
}

//function with 1 return value
func findMax(num1, num2 int) int{
  var result int;
  if(num1 < num2){
    result = num1;
  }else{
    result = num2;
  }
  return result;
}

func findPrimeNumber(maxNumber int){
  var i,j int;
  for i=2; i< maxNumber;i++{
    for j=2; j <= (i/j);j++{
      if(i%j == 0){
        break;
      }
    }
    if(j > (i/j)){
      fmt.Printf("%d is prime\n",i);
    }
  }
}

func howToUseSelectKeyword(){
  var c1, c2 , c3 chan int;
  var i1, i2 int;

  select{
    case i1 = <-c1:
      fmt.Printf("received ",i1,"from c1\n");
    case c2 <- i2:
      fmt.Printf("sent ",i2, " to c2\n");
    case i3, ok:= (<-c3): //like initialize 2 variable
      if(ok){
        fmt.Printf("recieved ",i3," from c3\n");
      }else{
        fmt.Printf("c3 is closed\n");
      }
    default:
      fmt.Printf("no communication\n");
  }
}

func printAndPlayVariable(){
  //variable decration & initialization
  var number1 int;
  number1 = 100;

  var number2 float64 = 20.0;
  number3  := 100.56;

  //mixed variable
  var var1, var2, var3 = 3 , 10.56, "foo";
  total1 := var2 + number3;
  bool1 := true;

  const LENGTH int = 10;
  const WIDTH int = 5;
  area := LENGTH * WIDTH;
  fmt.Println("Area : ",area);


  /*this is comment*/
  fmt.Println(number1);
  fmt.Println(number2);
  fmt.Println(number3, var1, var2, var3);
  fmt.Println("Total :",total1);
  fmt.Printf("testing\n")
  fmt.Printf("Hello world\n");

  //this is if statement
  if(area <= 50){
    fmt.Println("area is less than or equals 50, that's not enough room for small studio");
  }else{
    fmt.Println("area is more than 50, that's enought room yo ");
  }

  if(var1 == 3 && bool1){
    fmt.Println("var1 value == 3 && bool1 == true");
    if(var2 == 10.56){
      fmt.Println("var2 == 10.56");
    }
  }
}

func tryFuncAsValueAndFunctionClosureAndMethod(){
  getSquareRoot := func(x float64)float64{
    return math.Sqrt(x);
  }
  fmt.Println("sqrt of 9 is : ",getSquareRoot(9));

  nextNumber := getSequence();
  fmt.Println("this is closure : ",nextNumber())
  fmt.Println("this is closure : ",nextNumber())

  nextNumber2 := getSequence();
  fmt.Println("this is closure : ",nextNumber2())
  fmt.Println("this is closure : ",nextNumber2())

  newCircle := Circle{x:0, y:0, radius:5};
  fmt.Printf("Circle area : %f\n",newCircle.area())
}

//closure
func getSequence() func() int{
  i := 0;
  return func() int{
    i += 1;
    return i;
  }
}

func main() {
  markSomething();
  printAndPlayVariable();
  checkIntefaceOfVariable();
  howToUseSelectKeyword();
  tryLooping();
  tryFuncAsValueAndFunctionClosureAndMethod();
}
