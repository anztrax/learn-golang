package main

import (
  "fmt"
  "math"
  "strings"
)
//you can import package using single statement
import "errors"

//declare type
type Circle struct{
  x, y, radius float64
}

type Book struct{
  title string
  author string
  subject string
  book_id int
}

type Shape interface{
  area1() float64;
}

type Rectangle struct{
  width, height float64;
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

  variadicTotal := addValueWithVariadic(10,20,40,50);
  fmt.Printf("variadic add total : %d\n",variadicTotal);
}

func playWithStringAndArray(){
  var greeting = "Hello world";
  var myName = []string{"andrew","ananta"};
  for i:=0; i < len(greeting);i++{
    fmt.Printf("%c : %x\n",greeting[i],greeting[i]);
  }

  fmt.Println("my name ",strings.Join(myName, " "));

  //and play with array too
  var queueSlotNameList [10] string;
  queueSlotNameList[0] = "andrew";

  //declaring and initialze it
  var queueSlotNameList2 = [10]string{"crystal","kay","takeo","hideo"};
  fmt.Println(queueSlotNameList2);
  var balances = []float32{100.0,2.0,3.4,7.0,50.0};
  balances[4] = 100;

  salary1 := balances[4];
  fmt.Println(salary1);

  //raise employee salary by triple yeah :D
  fmt.Println("yeah salary increases by 3 times\n================");
  for i :=0; i < len(balances);i++{
    balances[i] = balances[i] * 3;
    fmt.Printf("balances[%d] = %f\n",i,balances[i]);
  }
  fmt.Println("=======================");

  //multi dimentional array
  var multiDimentioanlArray [5][10] int;
  multiDimentioanlArray[0][0] = 0;

  var multiDimentionalArray2 = [3][4] int{
    {0,0,0,0},
    {0,0,0,0},
    {0,0,0,0},    //this is a must yeah
  };
  fmt.Println(multiDimentionalArray2);
  for i:=0; i < len(multiDimentionalArray2);i++{
    for j:=0; j < len(multiDimentionalArray2[i]);j++{
      fmt.Printf("%d",multiDimentionalArray2[i][j]);
    }
    fmt.Println();
  }

  fmt.Println("the avarage of 10,20,30,40 : ",getAverage([]uint64{10,20,30,40}));

  //pointers
  var number1 = 100;
  var float1 = 100.0;
  var ip *int;
  var fp *float64;
  ip = &number1;
  fp = &float1;
  fmt.Printf("number1 value : %d (address : %d) \n, number2 value : %f (address : %d) \n",*ip,ip,*fp,fp);

  //array of pointer
  var arrayOfNumber = [3]int{10,20,30};
  var arrayOfPtr [3]*int;
  for i:=0;i < len(arrayOfNumber);i++{
    arrayOfPtr[i] = &arrayOfNumber[i];
  }

  //if we change the pointer this will changes
  *arrayOfPtr[0] = 100;
  for i:=0; i < len(arrayOfPtr);i++{
    fmt.Println("array of the pointer :",*arrayOfPtr[i]);
  }

  //pointer of pointer become make sense now :)
  var ptr1 *int;
  var pptr1 **int;
  var pptr2 ***int;

  ptr1 = &number1;
  pptr1 = &ptr1;
  pptr2 = &pptr1;
  fmt.Println("========= pointer stuff ==============");
  fmt.Printf("value of number1 : %d\n",number1);
  fmt.Printf("value of ptr1 : %d (address : %d)\n",*ptr1,ptr1);
  fmt.Printf("value of pptr1 : %d\n",**pptr1);
  fmt.Printf("value of pptr2 : %d\n",***pptr2);

  var intPtr *int;
  if(intPtr == nil){
    fmt.Println("this intPtr is not inialized yet");
  }

}

func printSlices(anArray []int){
  fmt.Printf("len = %d , cap = %d , slice : %v \n",len(anArray),cap(anArray),anArray);
}

func tryAndExperimentWithSliceAndRange(){
  var arrayOfNumbers = make([]int,3,5); //this make method is make array with parameter (length & capacity)
  arrayOfNumbers[2] = 10;
  printSlices(arrayOfNumbers);

  var nilNumbers []int;
  if(nilNumbers == nil){
    fmt.Println("Slice is nil");
  }

  numbers := []int{0,1,2,3,4,5,6};
  fmt.Println("numbers[1:4]",numbers[1:4]);
  fmt.Println("numbers[1:]",numbers[1:]);
  fmt.Println("numbers[:4]",numbers[:4]);

  fmt.Println("cap =",cap(numbers), "len =",len(numbers));
  numbers2 := append(numbers,1);
  fmt.Println("numbers2 : ",numbers2);

  //using range
  for i := range numbers2{
    fmt.Println("index : ",i," is : ",numbers2[i]);
  }

  //define map and loop into it
  gundamAndUserList := map[string]string{"a" : "zaku", "b" : "gundam strike", "c": "gundam buster"};

  fmt.Println("===============================");
  for player,gundam := range gundamAndUserList{
    fmt.Println("user : ",player," gundam name : ",gundam);
  }
  gundamAndUserList["d"] = "zeong";

  fmt.Println("===============================");
  playerA, ok := gundamAndUserList["c"];
  if(!ok){
    fmt.Println("player C is not exist");
  }else{
    fmt.Println("player C is exist, the gundam is : ",playerA);
  }

  //for deleting map data
  delete(gundamAndUserList,"d");
  fmt.Println("===============================");
  playerD, ok := gundamAndUserList["d"];
  if(!ok){
    fmt.Println("player D is not exists");
  }else{
    fmt.Println(playerD," is still there");
  }
}

func playWithStruct(){
  var book1 Book;
  book1.title = "Go Programming";
  book1.author = "S Kumar";
  book1.subject = "Programming"
  book1.book_id = 1;

  var book2 Book;
  book2.title = "C++ Programming";
  book2.author = "Daniel M Ritchie";
  book2.subject = "Programming"
  book2.book_id = 2;

  //temp struct pointer;
  changeBookSubject(&book1,"Fun");
  changeBookSubject(&book2,"This is Fun");

  printBookDetail(book1);
  printBookDetail(book2);
}


func printBookDetail(book Book){
  fmt.Println("=====================");
  fmt.Println("Book title : ",book.title);
  fmt.Println("Book author : ",book.author);
  fmt.Println("Book subject : ",book.subject);
  fmt.Println("Book id : ",book.book_id);
}

func changeBookSubject(book *Book,subject string){
  book.subject = subject;
}

func getAverage(arr []uint64) float32{
  var sum uint64;
  var lenOfArray = len(arr);
  for i:=0;i < lenOfArray;i++{
    sum += arr[i];
  }
  return float32(sum / uint64(lenOfArray));
}

//closure
func getSequence() func() int{
  i := 0;
  return func() int{
    i += 1;
    return i;
  }
}

func findFactorial(i int)int{
  if( i <= 1){
    return 1;
  }
  return i * findFactorial(i -1);
}

func recursiveExample(){
  fmt.Println("10 ! " ,findFactorial(10));
}

func testInterface(){
  circle1 := Circle{x :0, y:0, radius: 10};
  rectangle1 := Rectangle{width : 100,height : 100};
  fmt.Printf("circle area : %f\n",getArea(circle1))
  fmt.Printf("rectangle area : %f\n",getArea(rectangle1))

  sqrtResult, sqrtResultErr := calculateSqrt(-1);
  if(sqrtResultErr != nil){
    fmt.Println("err : ",sqrtResultErr);
  }else{
    fmt.Println("result : ",sqrtResult);
  }
}

//test variadic function
func addValueWithVariadic(args ... int) int{
  total := 0;
  for _, v := range args{
    total += v;
  }

  return total;
}

func calculateSqrt(value float64)(float64,error){
  if(value < 0){
    return 0, errors.New("Math : negative number can't passed to Sqrt");
  }

  return math.Sqrt(value), nil;
}


//inteface method
func (circle Circle) area1() float64{
  return math.Pi * circle.radius * circle.radius;
}
func (rect Rectangle) area1() float64{
  return rect.width * rect.height;
}

func getArea(shape Shape) float64{
  return shape.area1();
}

func testFuncFeatures(){
  first := func(){
    fmt.Println("1st");
  }

  second := func(){
    fmt.Println("2nd");
  }

  testDeferFeature := func(){
    defer first();
    second(); //this will be execute first
  }

  testDeferFeature();
}

func main() {
  markSomething();
  printAndPlayVariable();
  checkIntefaceOfVariable();
  howToUseSelectKeyword();
  tryLooping();
  tryFuncAsValueAndFunctionClosureAndMethod();
  playWithStruct();
  playWithStringAndArray();
  tryAndExperimentWithSliceAndRange();
  recursiveExample();
  testInterface();
  testFuncFeatures();
}
