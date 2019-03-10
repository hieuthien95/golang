# Go Programming Language

The Go Programming Language, also commonly referred to as Golang, is a general-purpose programming language, developed by a team at Google. The Go language was conceived in September 2007 by Robert Griesemer, Rob Pike, and Ken Thompson at Google. Go first appeared in November 2009, and the first version of the language was released in December 2012. The official web site of the Go project is http://golang.org/. Go has its own elegance and programming idioms that make the language productive and fun to code. Go also provides a comprehensive standard library. The standard library provides all the core packages programmers need to build real-world, web and network-based programs. Go is a statically typed, natively compiled, garbage-collected, concurrent programming language that belongs primarily to the C family of languages in terms of basic syntax.

The Go Programming Language is an open source project that is distributed under a BSD-style license to make programmers more productive. Go is expressive, concise, clean, and efficient programming language. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a faster, statically typed, compiled language that feels like a dynamically typed, interpreted language.

Go is a compiled language, and like many languages, it makes heavy use of the command line. Go is both the name of the programming language and the name for the toolset used to build and interact with Go programs.

Go is a statically typed programming language. What that means is the compiler always wants to know what the type is for every value in the program. When the compiler knows the type information ahead of time, it can help to make sure that the program is working with values in a safe way. This helps to reduce potential memory corruption and bugs, and provides the compiler the opportunity to produce more perform-ant code. Go struct lets you create your own types by combining one or more types, including both built-in and user-defined types. Structs are the only way to create concrete user-defined types in Go. When you create your own types using struct, it is important to understand that Go does not provide support for inheritance in its type system, but it favors composition of types that lets you create larger types by combining smaller types. The design philosophy of Go is to create larger components by combining smaller and modular components. If you are a pragmatic programmer, you will appreciate the design philosophy of Go that favors composition over inheritance because of its practical benefits. The inheritance of types sometimes introduces practical challenges with regard to maintainability.

In the last decade, computer hardware has evolved to having many CPU cores and more power. Nowadays we heavily leverage cloud platforms for building and running applications where servers on the cloud have more power. Although modern computers and virtual machine instances on the cloud have more power and many CPU cores, we still can't leverage the power of modern computers using most of the existing programming languages and tools. Concurrency in Go is the ability for functions to run independent of each other. Its concurrency mechanisms make it easy to write programs that get the most out of multi core and networked machines, while its novel type system enables flexible and modular program construction. When a function is created as a goroutine, it's treated as an independent unit of work that gets scheduled and then executed on an available logical processor. Goroutines are created by calling the Go statement followed by the function or method that you want to run as an autonomous activity. The Go run-time scheduler is a sophisticated piece of software that manages all the goroutines that are created and need processor time. The scheduler sits on top of the operating system, binding operating system's threads to logical processors which, in turn, execute goroutines. The scheduler controls everything related to which goroutines are running on which logical processors at any given time.

# Data Type

Go is a statically typed programming language. This means that variables always have a specific type and that type cannot change. The keyword var is used for declaring variables of a particular data type. Here is the syntax for declaring variables:

```go
var name type = expression
```

On the left we use the var keyword to declare a variable and then assign a value to it. We can declare mutiple variables of the same type in a single statement as shown here:

```go
var fname,lname string
```

Multiple variables of the same type can also be declared on a single line: var x, y int makes x and y both int variables. You can also make use of parallel assignment: a, b := 20, 16 If you are using an initializer expression for declaring variables, you can omit the type using short variable declaration as shown here:

```go
country, state := "Germany", "Berlin"
```

We use the operator : = for declaring and initializing variables with short variable declaration. When you declare variables with this method, you can't specify the type because the type is determined by the initializer expression.

```go
package main
 
import "fmt"
 
//Global variable declaration
var (m int
    n int)
 
func main(){
    var x int = 1 // Integer Data Type
    var y int    //  Integer Data Type 
    fmt.Println(x)
    fmt.Println(y)
     
    var a,b,c = 5.25,25.25,14.15 // Multiple float32 variable declaration
    fmt.Println(a,b,c)
    
    city:="Berlin" // String variable declaration
    Country:="Germany" // Variable names are case sensitive
    fmt.Println(city) 
    fmt.Println(Country) // Variable names are case sensitive
     
    food,drink,price:="Pizza","Pepsi",125  // Multiple type of variable declaration in same line
    fmt.Println(food,drink,price)
    m,n=1,2
    fmt.Println(m,n)
}
```

# Integer Float String Boolean

Go provides both signed and unsigned integer arithmetic. There are four distinct sizes of signed integers—8, 16, 32, and 64. Bits represented by the types int8, int16, int32, and int64 and corresponding unsigned versions uint8, uint16, uint32, and uint64.

```go

package main
 
import "fmt"
 
func main(){
    var n1 uint8 // Unsigned 8-bit integers (0 to 255)
    n1 = 200
    fmt.Println(n1)
     
    var n2 uint16 // Unsigned 16-bit integers (0 to 65535)
    n2 = 54200
    fmt.Println(n2)
     
    var n3 uint32 // Unsigned 32-bit integers (0 to 4294967295)
    n3 = 98765214
    fmt.Println(n3)
     
    var n4 uint64 // Unsigned 64-bit integers (0 to 18446744073709551615)
    n4 = 1844674073709551615
    fmt.Println(n4)
     
    var n5 int8 //Signed 8-bit integers (-128 to 127)
    n5 = -52
    fmt.Println(n5)
    fmt.Println(n5*-1)
     
    var n6 int16 // Signed 16-bit integers (-32768 to 32767)
    n6 = -32552
    fmt.Println(n6)
    fmt.Println(n6*-1)
     
    var n7 int32 // Signed 32-bit integers (-2147483648 to 2147483647)
    n7 = -98658754
    fmt.Println(n7)
    fmt.Println(n7*-1)
     
    var n8 int64 // Signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
    n8 = -92211111111111111
    fmt.Println(n8)
    fmt.Println(n8*-1)
}
 
 
Output 
 
200  
54200         
98765214        
1844674073709551615        
-52                                                                                                                                                     
52  
-32552        
32552 
-98658754
98658754
-92211111111111111
92211111111111111
```

Go provides two sizes of floating-point numbers, float32 and float64. Their arithmetic properties are governed by the IEEE 754 standard implemented by all modern CPUs. Go provides two sizes of complex numbers, complex64 and complex128, whose components are float32 and float64 respectively.

```go
package main
 
import "fmt"
 
func main(){
    var f1 float32 = 1677.7216 // IEEE-754 32-bit floating-point numbers
    fmt.Println(f1)
     
    var f2 float64 = 18787677.878716 // IEEE-754 64-bit floating-point numbers
    fmt.Println(f2)
     
    var f3 complex64 = 8789579877.721 // Complex numbers with float32 real and imaginary parts
    fmt.Println(f3)
     
    var f4 complex128 = 985214745.7216 // Complex numbers with float64 real and imaginary parts
    fmt.Println(f4)
}
 
 
Output 
 
1677.7216
1.8787677878716e+07                                                                                                                                    
(8.78958e+09+0i) 
(9.852147457216e+08+0i)   
```

A boolean value (named after George Boole) is a special 1-bit integer type used to represent true and false (or on and off).A boolean type represents the set of boolean truth values denoted by the predeclared constants true and false. The boolean type is bool.

Strings in Go are a sequence of UTF-8 characters enclosed in double quotes (") "Hello, World" or backticks `Hello, World`. If you use the single quote (') you mean one character (encoded in UTF-8) — which is not a string in Go. Once assigned to a variable the string can not be changed: strings in Go are immutable.A space is also considered a character.

```go
package main
 
import "fmt"
 
func main(){
    var a,b bool // Boolean type declaration
    a=true
    b=false
    fmt.Println(a,b)
     
    var city = "Washington" // " is necessary for string assignmnet
    fmt.Println(city)
    var country string 
    country = "Germany"
    fmt.Println(country)
}
 
Output 
 
true false
Washington
Germany 
```

# Golang Variables

> In this tutorial you will learn how to store information in a variable in Golang.

## What is Variable in Golang

Variables are "containers" for storing information, like string of text, numbers, etc. Variable values can change over the execution of a program. Here're some important things to know about Golang variables:

- Golang is statically typed language, which means that when variables are declared, they either explicitly or implicitly assigned a type even before your program runs.
- Golang requires that every variable you declare inside main() get used somewhere in your program.
- You can assign new values to existing variables, but they need to be values of the same type.
- A variable declared within brace brackets {}, the opening curly brace { introduces a new scope that ends with a closing brace }.

## Declaring an integer and string variable

The assignment of a value inline with the initialization of the variable.

```go
package main
 
import (
    "fmt"
)
 
func main() {
    var i int = 10
    var s string = "Japan"
    fmt.Println(i)
    fmt.Println(s)
}
```

A variable in Go is initialized by `var` keyword.
In above example variables are assigned by name `i, s` and type `integer, string` respectively.
The assignment operator `=` signifies that the variable should be assigned a value of whatever is to the right of `=` .
The `fmt` standard library package uses the variable name i and s as a reference to the value of i and s.

## Assignment after initialization

First a variable is initialized and then assigned a value later in program.

```go
package main
 
import (
    "fmt"
)
 
func main() {
    var intVar int
    var strVar string
 
    intVar = 10
    strVar = "Australia"
 
    fmt.Println(intVar)
    fmt.Println(strVar)
}
```

## Short Variable Declaration

The `:=` short variable assignment operator indicates that short variable declaration is being used.

```go
package main
 
import (
  "fmt"
)
 
func main() {
  s := "Japan"
  fmt.Println(s)
}
```

## Scope of Variables Defined by Brace Brackets

Inner block can access its outer block defined variables, but outer block cannot access inner block defined variables.

```go
package main
 
import (
  "fmt"
)
 
var s = "Japan"
 
func main() {
  fmt.Println(s)
  x := true
 
  if x {
    y := 1
    if x != false {
      fmt.Println(s)
      fmt.Println(x)
      fmt.Println(y)
    }
  }
  fmt.Println(x)
}
```

## Naming Conventions for Golang Variables

These are the following rules for naming a Golang variable:

- A name must begin with a letter, and can have any number of additional letters and numbers.
- A variable name cannot start with a number.
- A variable name cannot contain spaces.
- If the name of a variable begins with a lower-case letter, it can only be accessed within the current package this is considered as unexported variables.
- If the name of a variable begins with a capital letter, it can be accessed from packages outside the current package one this is considered as exported variables.
- If a name consists of multiple words, each word after the first should be capitalized like this: empName, EmpAddress, etc.
- Variable names are case-sensitive (car, Car and CAR are three different variables).

# Constants

Go also has support for constants. Constants are essentially variables whose values cannot be changed later. They are created in the same way you create variables, but instead of using the var keyword we use the const keyword:

```go
const x string = "Hello, World"
```

A const declaration defines named values that look syntactically like variables but whose value is constant, which prevents accidental (or nefarious) changes during program execution.

```go
package main
import "fmt"
const (
  x=10
  y=20
  z=30
)
func main(){
  const name string ="John Carry" // Constant with data type
  fmt.Println(name)
  const age = 35 // Constant without data type
  fmt.Println(age)
  fmt.Println(x,y,z)
}
```

The underlying type of every constant is a basic type: boolean, string, or number.

# Operators

Operators combine operands into expressions. Go language is rich in built-in operators and provides the following types of operators:

## Arithmetic Operators

Arithmetic operators apply to numeric values and yield a result of the same type as the first operand. The four standard arithmetic operators (+, -, *, /) apply to integer, floating-point, and complex types; + also applies to strings.

```go
package main
import "fmt"
func main() {
  var m,n int = 100,25
  var a,b string = "John","Carry"
  var sum,difference,product,quotient,remainder int
  var sumstring string
  sum = m+n // +   sum   integers, floats, complex values, strings
  sumstring = a+" "+b
    difference = m-n // -   difference   integers, floats, complex values
    product = m*n // *   product   integers, floats, complex values
    quotient = m/n // /   quotient   integers, floats, complex values
    remainder = m%n //%   remainder  integers
    fmt.Println(sum)
    fmt.Println(sumstring)
    fmt.Println(difference)
    fmt.Println(product)
    fmt.Println(quotient)
    fmt.Println(remainder)
}
```

## Bitwise Operators

The bitwise logical and shift operators apply to integers only. Bitwise operator works on bits and perform bit-by-bit operation.

```go
package main
import "fmt"
func main() {
  var x,y,z int
  x = 75
  y = 25
  z = x & y       // &    bitwise AND            integers
  fmt.Println(z)
  z = x | y       // |    bitwise OR             integers
  fmt.Println(z)
  z = x ^ y       // ^    bitwise XOR            integers
  fmt.Println(z)
  z = x &^ y      // &^   bit clear (AND NOT)    integers
  fmt.Println(z)
}
```

## Comparison operators

Comparison operators like == and < can also be used to compare a value of a named type to another of the same type, or to a value of the underlying type.

```go
package main
import "fmt"
func main() {
  var num1 int = 50
  var num2 int = 60
  fmt.Println(num1==num2) // ==    equal
  fmt.Println(num1!=num2) // !=    not equal
  fmt.Println(num1<num2)  // <     less
  fmt.Println(num1<=num2)  // <     less or equal
  fmt.println(num1>num2)  // >     greater
  fmt.Println(num1>=num2) // >=    greater or equal
}
```

## Logical Operators

Logical operators apply to boolean values and yield a result of the same type as the operands. The right operand is evaluated conditionally.

```go
package main
import "fmt"
func main() {
  var num1 int = 50
  var num2 int = 60
    
  if(num1!=num2 && num1<=num2){ // &&  Called Logical AND operator.
      fmt.Println("True")
  }
    
  if(num1!=num2 || num1<=num2){ // ||  Called Logical OR Operator
      fmt.Println("True")
  }
    
  if(!(num1==num2)){ // !  Called Logical NOT Operator. Use to reverses the logical state of its operand.
      fmt.Println("True")
  }
}
```

## Assignment Operators

Assignment statement that concatenates the old value of X with Y and assigns it back to X.
X=X+Y
X+=Y
The operator += is an assignment operator. Each arithmetic and logical operator like + or * has a corresponding assignment operator.

```go
package main
import "fmt"
func main() {
  var X int = 50
  var Y int = 60
    
  X+=Y    // += Add AND assignment operator
  fmt.Println(X)
    
  X=50
  Y=60
  X-=Y   // -= Subtract AND assignment operator
  fmt.Println(X)
    
  X=50
  Y=60
  X*=Y  // *= Multiply AND assignment operator
  fmt.Println(X)
    
  X=4
  Y=44
  X%=Y  // %= Modulus AND assignment operator
  fmt.Println(X)
    
  X=50
  Y=200
  X/=Y  // /= Divide AND assignment operator
  fmt.Println(X)
}
```

# Golang If...Else...Else If Statements

> In this tutorial you'll learn how to write decision-making conditional statements used to perform different actions in Golang.

## Golang Conditional Statements

Like most programming languages, Golang borrows several of its control flow syntax from the C-family of lang
uages. In Golang we have the following conditional statements:

- The **if** statement - executes some code if one condition is true
- The **if...else** statement - executes some code if a condition is true and another code if that condition is false
- The **if...else if....else** statement - executes different codes for more than two conditions
- The **switch...case** statement - selects one of many blocks of code to be executed

We will explore each of these statements in the coming sections.

## Golang - if Statement

The `if` statement is used to execute a block of code only if the specified condition evaluates to true.

### Syntax

```go
if  condition { 
  // code to be executed if condition is true
}
```

The example below will output "Japan" if the X is true:

```go
package main
 
import (
  "fmt"
)
 
func main() {
  var s = "Japan"
  x := true
  if x {
    fmt.Println(s)
  }
}
```

## Golang - if...else Statement

The `if....else` statement allows you to execute one block of code if the specified condition is evaluates to true and another block of code if it is evaluates to false.

### Syntax

```go
if  condition { 
  // code to be executed if condition is true
} else {
  // code to be executed if condition is false
}
```

The example below will output "Japan" if the X is 100:

```go
package main
 
import (
  "fmt"
)
 
func main() {
  x := 100
 
  if x == 100 {
    fmt.Println("Japan")
  } else {
    fmt.Println("Canada")
  }
}
```

## Golang - if...else if...else Statement

The `if...else if...else` statement allows to combine multiple if...else statements.

### Syntax

```go
if  condition-1 { 
  // code to be executed if condition-1 is true
} else if condition-2 {
  // code to be executed if condition-2 is true
} else {
  // code to be executed if both condition1 and condition2 are false
}
```

The example below will output "Japan" if the X is 100:

```go
package main
 
import (
  "fmt"
)
 
func main() {
  x := 100
 
  if x == 50 {
    fmt.Println("Germany")
  } else if x == 100 {
    fmt.Println("Japan")
  } else {
    fmt.Println("Canada")
  }
}
```

## Golang - if statement initialization

The `if` statement supports a composite syntax where the tested expression is preceded by an initialization statement.

### Syntax

```go
if  var declaration;  condition { 
  // code to be executed if condition is true
}
```

The example below will output "Germany" if the X is 100:

```go
package main
 
import (
  "fmt"
)
 
func main() {
  if x := 100; x == 100 {
    fmt.Println("Germany")
  }
}
```

# Golang Switch…Case Statements

> In this tutorial you will learn how to use the switch-case statement to perform different actions based on different conditions in Golang.

Golang also supports a switch statement similar to that found in other languages such as, Php or Java. Switch statements are an alternative way to express lengthy if else comparisons into more readable code based on the state of a variable.

## Golang - switch Statement

The `switch` statement is used to select one of many blocks of code to be executed.

Consider the following example, which display a different message for particular day.

```go
package main
 
import (
  "fmt"
  "time"
)
 
func main() {
  today := time.Now()
 
  switch today.Day() {
  case 5:
    fmt.Println("Today is 5th. Clean your house.")
  case 10:
    fmt.Println("Today is 10th. Buy some wine.")
  case 15:
    fmt.Println("Today is 15th. Visit a doctor.")
  case 25:
    fmt.Println("Today is 25th. Buy some food.")
  case 31:
    fmt.Println("Party tonight.")
  default:
    fmt.Println("No information available for that day.")
  }
}
```

The `default` statement is used if no match is found.

## Golang - switch multiple cases Statement

The `switch with multiple case` line statement is used to select common block of code for many similar cases.

```go
package main
 
import (
  "fmt"
  "time"
)
 
func main() {
  today := time.Now()
  var t int = today.Day()
 
  switch t {
  case 5, 10, 15:
    fmt.Println("Clean your house.")
  case 25, 26, 27:
    fmt.Println("Buy some food.")
  case 31:
    fmt.Println("Party tonight.")
  default:
    fmt.Println("No information available for that day.")
  }
}
```

## Golang - switch fallthrough case Statement

The `fallthrough` keyword used to force the execution flow to fall through the successive case block.

```go
package main
 
import (
  "fmt"
  "time"
)
 
func main() {
  today := time.Now()
 
  switch today.Day() {
  case 5:
    fmt.Println("Clean your house.")
    fallthrough
  case 10:
    fmt.Println("Buy some wine.")
    fallthrough
  case 15:
    fmt.Println("Visit a doctor.")
    fallthrough
  case 25:
    fmt.Println("Buy some food.")
    fallthrough
  case 31:
    fmt.Println("Party tonight.")
  default:
    fmt.Println("No information available for that day.")
  }
}
```

## Golang - swith conditional cases Statement

The `case` statement can also used with conditional operators.

```go
package main
 
import (
  "fmt"
  "time"
)
 
func main() {
  today := time.Now()
 
  switch {
  case today.Day() < 5:
    fmt.Println("Clean your house.")
  case today.Day() <= 10:
    fmt.Println("Buy some wine.")
  case today.Day() > 15:
    fmt.Println("Visit a doctor.")
  case today.Day() == 25:
    fmt.Println("Buy some food.")
  default:
    fmt.Println("No information available for that day.")
  }
}
```

## Golang - switch initializer Statement

The `switch` keyword may be immediately followed by a simple initialization statement where variables, local to the switch code block, may be declared and initialized.

```go
package main
 
import (
  "fmt"
  "time"
)
 
func main() {
  switch today := time.Now(); {
  case today.Day() < 5:
    fmt.Println("Clean your house.")
  case today.Day() <= 10:
    fmt.Println("Buy some wine.")
  case today.Day() > 15:
    fmt.Println("Visit a doctor.")
  case today.Day() == 25:
    fmt.Println("Buy some food.")
  default:
    fmt.Println("No information available for that day.")
  }
}
```

# Golang For Loops

> In this tutorial you will learn how to repeat a block of code execution using loops in Golang.

A `for` loop is used for iterating over a sequence (that is either a slice, an array, a map, or a string.

As a language related to the C-family, Golang also supports for loop style control structures.

Golang has no while loop because the for loop serves the same purpose when used with a single condition.

## Golang - traditional for Statement

The `for` loop is used when you know in advance how many times the script should run.

Consider the following example, display the numbers from 1 to 10 in three different ways.

```go
package main
 
import "fmt"
 
func main() {
 
  k := 1
  for ; k <= 10; k++ {
    fmt.Println(k)
  }
 
  k = 1
  for k <= 10 {
    fmt.Println(k)
    k++
  }
 
  for k := 1; ; k++ {
    fmt.Println(k)
    if k == 10 {
      break
    }
  }
}
```

## Golang - for range Statement

The `for` statement supports one additional form that uses the keyword range to iterate over an expression that evaluates to an array, slice, map, string, or channel

```go
package main
 
import "fmt"
 
func main() {
 
  // Example 1
  strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
  for index, element := range strDict {
    fmt.Println("Index :", index, " Element :", element)
  }
 
  // Example 2
  for key := range strDict {
    fmt.Println(key)
  }
 
  // Example 3
  for _, value := range strDict {
    fmt.Println(value)
  }
}
```

## Golang - range loop over string

The `for` loop iterate over each character of string.

Consider the following example, display "Hello" five times.


```go
package main
 
import "fmt"
 
func main() {
  for range "Hello" {
    fmt.Println("Hello")
  }
}
```

## Golang - Infinite loop

The `for` loop runs infinite times unless until we can't break.

Consider the following example, display "Hello" several times.

```go
package main
 
import "fmt"
 
func main() {
  i := 5
  for {
    fmt.Println("Hello")
    if i == 10 {
      break
    }
    i++
  }
}
```

# Source and Credit:

- http://www.golangprograms.com
- https://github.com/quii/learn-go-with-tests
- https://github.com/GoesToEleven/GolangTraining

