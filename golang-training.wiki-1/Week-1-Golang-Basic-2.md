# Functions

Each line of code we create to performs a specific task, and we combine these lines of code to carry out a desired result. Sometimes we desire to repeat the lines of code with different data, and in some actions our code becomes so long that keeping track of what each part does is hard. Here we use, Functions which serve as a system tools that keep your code spotless and orderly. 

> A function is a group of statements that exist within a program for the purpose of executing a specific task. In Golang, a function can be called multiple times within the package, and also from other packages if the functions are exported to other packages.

## Declaring Functions

Here is the syntax for writing functions in Go:

The first line of a function provides information about what the function will accept as an input and what to expect as an output and this we called as the function signature.

```go
func [function_name] (param1 type, param2 type...) (returned type1, returned type2...) { 
  //Function body 
} 
```

The func keyword signifies that this is the start point of a function. Next comes the name of the function. Then there is a set of brackets declares the expected variables (a list of parameters) for this function. After that there is closing bracket comes with an optional list of return types. The opening brace signifies the start of the function body, which is wrapped up by the closing bracket between them we write the logic for the function.

## Function with return value

In this example, the add() function takes input of two integer numbers and returns a integer value with a name of total.

```go
package main
 
import "fmt"
 
func main() {
  fmt.Println(add(20, 30))
}
 
func add(x int, y int) int {
  total := 0
  total = x + y
  return total
}
```

The types of input and return value must match with function signature. If we will modify the above program and pass some string value in argument then program will throw an exception "cannot use "test" (type string) as type int in argument to add".

## Function without return value

```go
package main
 
import "fmt"
 
func main() {
  add(20, 30)
}
 
func add(x int, y int) {
  total := 0
  total = x + y
  fmt.Println(total)
}
```

## Return types can have names

We can also name the return value by defining variables, here a variable total of integer type is defined in the function declaration for the value that the function returns.

```go
package main
 
import "fmt"
 
func main() {
  fmt.Println(add(20, 30))
}
 
func add(x int, y int) (total int) {
  total = x + y
  return total
}
```

# Deferred Functions Calls

Go has a special statement called defer that schedules a function call to be run after the function completes. Consider the following example:

```go
package main
import "fmt"
func first() {
  fmt.Println("First")
}
func second() {
  fmt.Println("Second")
}
func main() {
  defer second()
  first()
}
```

This program prints First followed by Second. 

A defer statement is often used with paired operations like open and close, connect and disconnect, or lock and unlock to ensure that resources are released in all cases, no matter how complex the control flow. The right place for a defer statement that releases a resource is immediately after the resource has been successfully acquired.

Below is the example to open a file and perform read/write action on it. In this example there are often spots where you want to return early.

## Without defer

```go
func ReadWrite() bool {
  file.Open("file")

  if failureX {
    file.Close()   //And here...
    return false
  }
  if failureY {
    file.Close()  //And here...
    return false
  }
  file.Close()  //And here...
  return true
}
```

A lot of code is repeated here. To overcome this Go has the defer statement. The code above could be rewritten as follows. This makes the function more readable, shorter and puts the Close right next to the Open.

## With defer

```go
func ReadWrite() bool {
  file.Open("file")
  defer file.Close()   //file.Close() is added to defer list
  // Do your thing
  if failureX {
    return false   // Close() is now done automatically
  }
  if failureY {
    return false   // And here too
  }
  return true   // And here
}
```

## This has various advantages

- It keeps our Close call near our Open call so it's easier to understand.
- If our function had multiple return statements (perhaps one in an if and one in an else), Close will happen before both of them.
- Deferred Functions are run even if a runtime panic occurs.
- Deferred functions are executed in LIFO order, so the above code prints: 4 3 2 1 0.
- You can put multiple functions on the "deferred list", like this example.

```go
package main
import "fmt"
func main() {
  for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
  }
}
```

# Variadic Functions

Variadic function is a function in which we can pass an infinite number of arguments to a function instead of just one argument at a time.

To declare a variadic function, the type of the final parameter is preceded by an ellipsis, "...", which shows that the function may be called with any number of arguments of this type.

## Select single argument from all arguments of variadic function

In below example we will are going to print `s[0]` the first and `s[3]` the forth, argument value passed to variadicExample() function.

```go
package main
 
import "fmt"
 
func main() {
  variadicExample("red", "blue", "green", "yellow")
}
 
func variadicExample(s ...string) {
  fmt.Println(s[0])
  fmt.Println(s[3])
}
```

Needs to be precise when running an empty function call, if the code inside of the function expecting an argument and absence of argument will generate an error "panic: run-time error: index out of range". In above example you have to pass at least 4 arguments.

## Passing multiple string arguments to a variadic function

The parameter s accepts an infinite number of arguments. The tree-dotted ellipsis tells the compiler that this string will accept, from zero to multiple values.

```go
package main
 
import "fmt"
 
func main() {
  variadicExample()
  variadicExample("red", "blue")
  variadicExample("red", "blue", "green")
  variadicExample("red", "blue", "green", "yellow")
}
 
func variadicExample(s ...string) {
  fmt.Println(s)
}
```

In the above example, we have called the function with single and multiple arguments; and without passing any arguments.

## Normal function parameter with variadic function parameter

```go
package main
 
import "fmt"
 
func main() {
  fmt.Println(calculation("Rectangle", 20, 30))
  fmt.Println(calculation("Square", 20))
}
 
func calculation(str string, y ...int) int {
  area := 1
  for _, val := range y {
    if str == "Rectangle" {
      area *= val
    } else if str == "Square" {
      area = val * val
    }
  }
  return area
}
```

## Pass different types of arguments in variadic function

In the following example, the function signature accepts an arbitrary number of arguments of type slice.

```go
package main
 
import (
    "fmt"
    "reflect"
)
 
func main() {
    variadicExample(1, "red", true, 10.5, []string{"foo", "bar", "baz"},
        map[string]int{"apple": 23, "tomato": 13})
}
 
func variadicExample(i ...interface{}) {
    for _, v := range i {
        fmt.Println(v, "--", reflect.ValueOf(v).Kind())
    }
}
```

# Panic and Recover

The built-in type system of GO Language catches many mistakes at compile time, but unable to check mistakes like an out-of-bounds array, access or nil pointer deference which require checks at run time. GO does not have an exception mechanism you can't throw exceptions. During the execution when Go detects these mistakes, it panics and stops all normal execution, all deferred function calls in that goroutine are executed and finally program crashes with a log message. This log message usually has enough information to analyze the root cause of the problem without running the program repeatedly, so it should always be included in a bug report about a panicking program.

Panic is a built-in function that stops the ordinary flow of control and begins panicking. When the function X calls panic, execution of X stops, any deferred functions in X are executed normally, and then X returns to its caller. To the caller, X then behaves like a call to panic. The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes. Panics can be initiated by invoking panic directly. They can also be caused by run-time errors, such as out-of-bounds array accesses.

Not all panics come from the run-time. The built-in panic function may be called directly; it accepts any value as an argument. A panic is usually the best thing to do when some "impossible" situation happens, for instance, execution reaches a case that logically can't happen:

```go

package main
import "fmt"
func main() {
  var action int
  fmt.Println("Enter 1 for Student and 2 for Professional")
  fmt.Scanln(&action)
  /*  Use of Switch Case in Golang */  
  switch action {
    case 1:
      fmt.Printf("I am a  Student")
    case 2:
      fmt.Printf("I am a  Professional")
    default:
      panic(fmt.Sprintf("I am a  %d",action))
  }   
  fmt.Println("")
  fmt.Println("Enter 1 for US and 2 for UK")
  fmt.Scanln(&action)
  /*  Use of Switch Case in Golang */  
  switch   {
    case 1:
      fmt.Printf("US")
    case 2:
      fmt.Printf("UK")
    default:
      panic(fmt.Sprintf("I am a  %d",action))
  }
}
```

In above program program will stop execution after first switch-case if user enters any other value other that 1 or 2.

Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.

```go
package main
import "fmt"
func main() {
  var action int
  fmt.Println("Enter 1 for Student and 2 for Professional")
  fmt.Scanln(&action)
  /*  Use of Switch Case in Golang */  
  switch action {
    case 1:
      fmt.Printf("I am a  Student")
    case 2:
      fmt.Printf("I am a  Professional")
    default:
      panic(fmt.Sprintf("I am a  %d",action))
  }    
  defer func() {
    action := recover()
    fmt.Println(action)
  }()
}
```

# Arrays

An array is a data structure that consists of a collection of elements of a single type. An array in Go is a fixed-length data type that contains a contiguous block of elements of the same type. This could be a built-in type such as integers and strings, or it can be a struct type.

## Declaration and Initialization

There are four common ways of declaring arrays.

### 1. A basic declaration of an array goes like this

```
[length]element_type
```

The type definition of an array is composed of its length, enclosed within brackets, followed by the type of its stored elements.

For example:

```go
var intArray [5]int
```

he name of this array is intArray and the type structure is a `[5]int`. The type includes the length of the array(number of items it must contain). A string array with a length of 10 items is of type `[10]string`. In this example, the array is automatically populated by compiler, which assigns to each of the 5 index positions, a default value 0(int array type); space(string array type).

The default values - set when we declare an array - can then be reassigned manually to whatever data we want to have, as long as this data is of the intended type:

For example:

```go
intArray[0] = 10
```

`index[0]` assigns a value 10 to the first element of the array intArray. Like other programming languages, indexing in array starts with zero.

### 2. Declare an array with some or all values to it at the time of declaration

```
element_type{comma-separated list of element values}
```

The literal value for an array is composed of the array type definition followed by a set of comma-separated values, enclosed in curly brackets.

For example:

```go
var intArray = [5]int{10,20,30}
```

In the above example, element[0] contains 10, element[1] contains 20 and remaining elements consist 0.

When arrays are initialized using an array literal, we have the option of customizing the index positions of the elements initially assigned to the array.

You can provide values for specific elements as shown here:

```go
var intArray = [5]int{0:10,2:30,4:50}
```

Index element 0,2 and 4 were assigned to values 10,30 and 50; remaining values will be 0.

### 3. Declare an array inside function using the := shortcut

We can also declare array inside of a function using short variable declaration := like in the below example:

```go
intArray := [5]int{10, 20, 30, 40, 50}
```

Index element 0,1...4 so on were assigned to values 10,20...50

### 4. Declaring an array with ellipses ...

The length of an array may be omitted and replaced by ellipses during initialization. Capacity is determined based on the number of values initialized.

The following will assign type `[5]int` to variable intArray:

```go
intArray := [...]int{10, 20, 30, 40, 50}
```

## Exercise

### 1. for loop to iterate over the Array

Create integer array named intArray, with a length of 5. Using a for loop to iterate over the array, display the contents of the array.

```go
package main
 
import "fmt"
 
func main() {
  intArray := [5]int{10, 20, 30, 40, 50}
  for i := 0; i < len(intArray); i++ {
    fmt.Println(intArray[i])
  }]
}
```

### 2. Assigning an array by value and reference

The values "Japan", "Australia", "Germany" are assigned to strArray1. strArray1 is then assigned to strArray2, which means that a copy of those values is created. The reassignment of the strArray1 to strArray3 from a value assignment to a reference assignment.

```go
package main
 
import "fmt"
 
func main() {
  strArray1 := [3]string{"Japan", "Australia", "Germany"}
  fmt.Printf("strArray1: %v\n", strArray1)
  strArray2 := strArray1 // data is passed by value (copied)
  fmt.Printf("strArray2: %v\n", strArray2)
  strArray1[0] = "Canada"
  fmt.Printf("strArray1: %v\n", strArray1)
  fmt.Printf("strArray2: %v\n", strArray2)

  strArray3 := &strArray1
  fmt.Printf("strArray3: %v\n", strArray3)
  fmt.Printf("&strArray3: %v\n", &strArray3)
  fmt.Printf("*strArray3: %v\n", *strArray3)
}
```

strArray3 is a pointer string array, which means strArray3 is a pointer to the memory address of strArray1. In order to dereference strArray3, we have to add the * operator as a prefix `*strArray3`.

### 3. An example program that explores the array type

```go
package main
import "fmt"
 
func main(){
  var x[5] int    // Array Declaration
  x[0]=10         // Assign the values to specific Index
  x[4]=20         // Assign Value to array index in any Order 
  x[1]=30
  x[3]=40
  x[2]=50
  fmt.Println("Values of Array X: ",x)

  // Array Declartion and Intialization to specific Index
  y := [5]int{0:100,1:200,3:500}
  fmt.Println("Values of Array Y: ",y)

  // Array Declartion and Intialization
  Country := [5]string{"US","UK","Australia","Russia","Brazil"}
  fmt.Println("Values of Array Country: ",Country)

  // Array Declartion without length and Intialization
  Transport := [...]string{"Train","Bus","Plane","Car","Bike"}
  fmt.Println("Values of Array Transport: ",Transport)
}
```

# Slices

A slice is a flexible and extensible data structure to implement and manage collections of data. A slice is a segment of dynamic arrays that can grow and shrink as you see fit. Like arrays, slices are indexable and have a length. Unlike arrays they're flexible in terms of length because they have their own built-in function called append, which can grow a slice quickly with efficiency. You can also trim the size of a slice by slicing out a part of the hidden memory.

Slices helps us in indexing, iteration, and garbage collection optimizations because the hidden memory is allocated in adjoining blocks.

## Creating Slices

> There are three common ways to declare slices.

### 1. A basic declaration of an slice using make() function

```
make([]Type, length,capacity)
```

A slice can be initialized at runtime using the built-in function `make`

An example program that declares slices in golang:

```go
package main
 
import "fmt"
 
func main() {
  var intSlice1 = make([]int,10)  // when length and capacity is same
  var intSlice2 = make([]int,10,20) // when length and capacity is different

  fmt.Printf("intSlice1 \tLen: %v \tCap: %v\n", len(intSlice1), cap(intSlice1))
  fmt.Printf("intSlice1 \tLen: %v \tCap: %v\n", len(intSlice2), cap(intSlice2))
}
```

Creates an underlying array of type `[10]int`

Creates the slice value with length and capacity of `10`

Third parameter that specifies the maximum capacity of the slice is optional.

### 2. Declaration using new keyword

```go
package main
 
import "fmt"
 
func main() {
  var intSlice1 = new([50]int)[0:10]
  fmt.Printf("intSlice1 \tLen: %v \tCap: %v\n", len(intSlice1), cap(intSlice1))   
}
```

In above declaration slice starts with a length of `10`, going from `0` to `50`.

Capacity is `50` so it can expand up-to 50 without requesting a new array.

### 3. Literal declaration

> []slice_type{comma-separated list of element values}

The following code snippet illustrates slice variables initialized with composite literal values:

```go
package main
 
import "fmt"
 
func main() {
  var intSlice = []int{10,20,30,40}
  var strSlice = []string{"India","Canada","Japan"}
  fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
  fmt.Println(intSlice)   
  fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))  
  fmt.Println(strSlice)
}
```

The length and capacity are of 3 and 4 elements respectively.

Number of elements provided in the literal is not surrounded by a fixed size. This suggests that the literal can be as large as needed.

## Exercise

### 1. Enlarge a slice using the append function

Append adds elements onto the end of a slice. If there's sufficient capacity in the underlying array, the element is placed after the last element and the length is incremented. However, if there is not sufficient capacity, a new array is created, all of the existing elements are copied over, the new element is added onto the end, and the new slice is returned.

The following code snippet enlarge a slice using the append Function:

```go
package main
 
import "fmt"
 
func main() {
  // Create a smaller slice
  a := make([]int, 2, 5)
  a[0] = 10  
  a[1] = 20  
  fmt.Println("Slice A:", a)
  fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))
    
  // Create a bigger slice
  a = append(a, 30, 40, 50, 60, 70, 80, 90)
  fmt.Println("Slice A after appending data:", a)
  fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a)) 
}
```

Using `make()`, create a slice of type `[]int` and name it `a`

Declare a length of `2` and capacity of `5` elements.

Idiomatic way in Go is to use `:=` opertor instead of `var` when creating slice inside `main` function.

`7` elements added in slice using append function which increased the length to `9` and capacity to `12`.

### 2. Enlarge a slice using the copy function

Go's built-in copy function is used to copy data from one slice to another. copy takes two arguments: dst and src. All of the entries in src are copied into dst overwriting whatever is there. If the lengths of the two slices are not the same, the smaller of the two will be used.

The following code snippet enlarge a slice using the copy Function:

```go
package main
 
import "fmt"
 
func main() {
  // Create a smaller slice
  a := []int{5, 6, 7}
  fmt.Printf("[Slice:A] Length is %d Capacity is %d\n", len(a), cap(a))
  // Create a bigger slice
  b := make([]int, 5, 10)
  copy(b, a)  // copy function
  fmt.Printf("[Slice:B] Length is %d Capacity is %d\n", len(b), cap(b))

  fmt.Println("Slice B after copying:", b)
  b[3] = 8
  b[4] = 9
  fmt.Println("Slice B after adding elements:", b)
}
```

Using `make()`, create a slice of type `[]in`t and name it `b`

Declare a length of `5` and capacity of `10` elements.

copy funcation copies all elements for `a` into `b`

### 3. Slice tricks

The following code snippet illustrates various slicing tricks:

```go
package main
 
import "fmt"
 
func main() {   
  var countries = []string{"india", "japan", "canada", "australia", "russia"}
    
  fmt.Printf("Countries: %v\n", countries)
  fmt.Printf(":2 %v\n", countries[:2])
  fmt.Printf("1:3 %v\n", countries[1:3])
  fmt.Printf("2: %v\n", countries[2:])
  fmt.Printf("2:5 %v\n", countries[2:5])
  fmt.Printf("0:3 %v\n", countries[0:3])
    
  fmt.Printf("Last element: %v\n", countries[4])
  fmt.Printf("Last element: %v\n", countries[len(countries)-1])
  fmt.Printf("Last element: %v\n", countries[4:])
  fmt.Printf("All elements: %v\n", countries[0:len(countries)])
  fmt.Printf("Last two elements: %v\n", countries[3:len(countries)])
  fmt.Printf("Last two elements: %v\n", countries[len(countries)-2:len(countries)])
    
  fmt.Println(countries[:])
  fmt.Println(countries[0:])
  fmt.Println(countries[0:len(countries)])
}
```

Declare a slice with 5 elements.

Using square bracket we can create different slices of a `slices`

`len` function used to calculate the length of slice.

### 4. Assign parts of slice to another slice

This technique illustrates the usage of ranges to select elements of a slice.

```go

package main
 
import "fmt"
 
func main() {   
  var oldStr = []string{"india", "japan", "canada", "australia", "russia"}
  var strSlice []string
  newStr := oldStr[0:3]
  strSlice = append(strSlice, oldStr[:1]...)
  fmt.Printf("newStr: %v\n", newStr)
  fmt.Printf("strSlice: %v\n", strSlice)
    
  fmt.Printf("oldStr length: %v\tcapacity: %v\n", len(oldStr), cap(oldStr))
  fmt.Printf("newStr length: %v\tcapacity: %v\n", len(newStr), cap(newStr))
  fmt.Printf("strSlice length: %v\tcapacity: %v\n", len(strSlice), cap(strSlice))
    
  newStr[0] = "china"
  fmt.Printf("newStr: %v\n", newStr)
  fmt.Printf("oldStr: %v\n", oldStr)
    
  newStr = append(newStr, "brazil")
  fmt.Printf("newStr: %v\n", newStr)
  fmt.Printf("oldStr: %v\n", oldStr)
    
  oldStr = append(oldStr, "us")
  newStr = append(newStr, "uk")
  fmt.Printf("newStr: %v\n", newStr)
  fmt.Printf("oldStr: %v\n", oldStr)
}
```

Declare a `slice` with `5` elements.

Declare `slice` `newStr` assign first three values of `oldStr`

### 5. Append a slice to an existing slice

This technique illustrates the usage of triple-dot `...` ellipsis to append a slice.

```go
package main
 
import "fmt"
 
func main() {   
  var slice1 = []string{"india", "japan", "canada"}
  var slice2 = []string{"australia", "russia"}
  slice2 = append(slice2, slice1...)
  fmt.Printf("slice1: %v\n", slice1)
  fmt.Printf("slice2: %v\n", slice2)
}
```

Declare two slices with `2` and `3` items.

Append, the first slice `slice1` to the second slice `slice2`

### 6. Append part of slice to an existing slice

This technique illustrates the usage of ellipsis with range to append a slice.

```go
package main
 
import "fmt"
 
func main() {   
  var slice1 = []string{"india", "japan", "canada", "us", "uk", "italy", "germany"}
  var slice2 = []string{"australia", "russia"}
  slice2 = append(slice2, slice1[3:]...)  
  fmt.Printf("slice1: %v\n", slice1)
  fmt.Printf("slice2: %v\n", slice2)
}
```

Declare two slices with 2 and 3 items.

Append, the first slice `slice1` to the second slice `slice2`

# Maps

A map is a data structure that provides you with an unordered collection of key/value pairs (maps are also sometimes called associative arrays in Php, hash tables in Java, or dictionaries in Python). Maps are used to look up a value by its associated key. You store values into the map based on a key. The strength of a map is its ability to retrieve data quickly based on the key. A key works like an index, pointing to the value you associate with that key.

A map is implemented using a hash table, which is providing faster lookups on the data element and you can easily retrieve a value by providing the key. maps are unordered collections, and there's no way to predict the order in which the key/value pairs will be returned. Every iteration over a map could return a different order.

## Map declaration in golang goes like this:

```
make[Key-type] [Value-type]
```

The Key-type specifies the type of a value that will be used to index the stored elements(Value-type) of the map. Map keys can be of any type numeric, string, Boolean, pointers, arrays, struct, and interface.

## Maps Examples

### 1. Empty Map declaration

```go
package main
import "fmt"
var employee = map[string]int{}

func main() {
  fmt.Println(employee)
}
```

Map employee created having string as key-type and int as value-type

### 2. Map initialization

The literal mapped values are specified using a colon-separated pair of key and value as shown in below example.

```go
package main
import "fmt"
var employee = map[string]int{"Mark":10,"Sandy":20}

func main() {
  fmt.Println(employee)
}
```

The type of each key and value pair must match that of the declared elements in the map.

### 3. Creating Map using the make function

A `map` value can also be initialized using the `make` function.

```go
package main
import "fmt"

func main() {
  var employee = make(map[string]int)
  employee["Mark"] = 10
  employee["Sandy"] = 20
  fmt.Println(employee)
  
  employeeList := make(map[string]int)
  employeeList["Mark"] = 10
  employeeList["Sandy"] = 20
  fmt.Println(employeeList)
}
```

The `make` function takes as argument the type of the map and it returns an initialized map.

### 4. Find length of Map

The built-in len() function returns the number of elements in a map.

```go
package main

import "fmt"

func main() {
  var employee = make(map[string]int)
  employee["Mark"] = 10
  employee["Sandy"] = 20

  // Empty Map
  employeeList := make(map[string]int)

  fmt.Println(len(employee))     // 2
  fmt.Println(len(employeeList)) // 0
}
```

The len function will return zero for an uninitialized map.

### 5. Delete element from a Map

The built-in delete function deletes an element from a given map associated with the provided key.

```go
package main

import "fmt"

func main() {
  var employee = make(map[string]int)
  employee["Mark"] = 10
  employee["Sandy"] = 20  
  employee["Rocky"] = 30
  employee["Josef"] = 40
  
  fmt.Println(employee)
  
  delete(employee,"Mark")
  fmt.Println(employee)
}
```

In above example `delete` function used to delete first element from employee map by passing key Mark as second argument in delete function.

### 6. Adding and Editing elements in Map

```go
package main
import "fmt"

func main() {
  var employee = map[string]int{"Mark":10,"Sandy":20}
  fmt.Println(employee)   // Initial Map
  
  employee["Rocky"] = 30  // Add element
  employee["Josef"] = 40  
  employee["Mark"] = 50   // Edit element
  fmt.Println(employee)
}
```

2 elements added and 1 edited in employee map after initialization.

### 7. for range loop to iterate over a Map

The `for…range` loop statement can be used to fetch the index and element of a map.

```go
package main

import "fmt"

func main() {
  var employee = map[string]int{"Mark": 10, "Sandy": 20,
    "Rocky": 30, "Rajiv": 40, "Kate": 50}
  for key, element := range employee {
    fmt.Println("Key:", key, "=>", "Element:", element)
  }
}
```

Each iteration returns a key and its correlated element content.

# Struct

Go have the ability to declare and create own data types by combining one or more types, including both built-in and user-defined types. The declaration of new data type is constructed to provide the compiler with size and representation information, similar like built-in data types.

Structs are the only way to create concrete user-defined types in Go. Struct types are declared by composing a fixed set of unique fields. Each field in a struct is declared with a known type, which could be a built-in type or another userdefined type.

## Struct Declaration

The syntax of a struct is as follows:

```
type identifier struct{
  field1 data_type
  field2 data_type
  field3 data_type
}
```

The `struct` type is composed by defining the keyword `struct` followed by a set of field declarations enclosed within curly brackets.

An example program that declares `struct` in golang:

```go

package main
 
import "fmt"
 
type rectangle struct {
  length  float64
  breadth float64
  color   string
}
 
func main() {
  fmt.Println(rectangle{10.5,25.10,"red"})
}
```

As an example, a `struct` named `rectangle` containing two fields of type `float64` and a third field of type `string`.

Different fields in the same struct can have different data types.

Data type for each field can be of any type int, float, string etc...

## Different ways of Struct Instantiation

### 1. Dot notation

A struct uses a selector expression (or dot notation) to access the values stored in fields.

The following code snippet illustrates struct instantiating:

```go
package main
 
import "fmt"
 
type rectangle struct {
  length  int
  breadth int
  color   string
  geometry struct{
    area int
    perimeter int
  }
}
 
func main() {
  var rect rectangle 
  rect.length =  10
  rect.breadth=  20
  rect.color  =  "Green"
    
  rect.geometry.area   =  rect.length * rect.breadth
  rect.geometry.perimeter =  2 * (rect.length + rect.breadth)
    
  fmt.Println(rect)
  fmt.Println("Area:\t", rect.geometry.area)
  fmt.Println("Perimeter:", rect.geometry.perimeter)
}
```

In the above example, selectors dot can be used as a `rect.geometry.perimeter `chain to fetch field values that are nested inside a struct.

### 2. var keyword and := operator

The following code snippet illustrates struct instantiating using var and := :

```go
package main
 
import "fmt"
 
type rectangle struct {
  length  int
  breadth int
  color   string
}
 
func main() {
  var rect1 = rectangle{10,20,"Green"}
  fmt.Println(rect1)
    
  var rect2 = rectangle{length:10,color:"Green"} // breadth value skipped
  fmt.Println(rect2)
    
  rect3 := rectangle{10,20,"Green"}
  fmt.Println(rect3)
    
  rect4 := rectangle{length:10,breadth:20,color:"Green"}
  fmt.Println(rect4)  
        
  rect5 := rectangle{breadth:20,color:"Green"}    // length value skipped
  fmt.Println(rect5)
}
```

You can skip value when you are specified fields name.

### 3. Using new keyword

The following code snippet illustrates struct instantiating using new keyword :

```go
package main
 
import "fmt"
 
type rectangle struct {
    length  int
    breadth int
    color   string
}
 
func main() {
  rect1 := new(rectangle)
  rect1.length = 10
  rect1.breadth = 20
  rect1.color  = "Green"
  fmt.Println(rect1)
    
  var rect2 = new(rectangle)
  rect2.length = 10  
  rect2.color  = "Red"
  fmt.Println(rect2)  // breadth skipped
}
```

`rect2` is a pointer to an instance of `rectangle`

The fields are initially assign as a default of data type then we can reassign the field by using dot operator `rect2.length`.

### 4. Using &

The following code snippet illustrates struct instantiating using `&`:

```go
package main
 
import "fmt"
 
type rectangle struct {
  length  int
  breadth int
  color   string
}
 
func main() {
  var rect1 = &rectangle{10,20,"Green"} // Can't skip any value
  fmt.Println(rect1)
    
  var rect2 = &rectangle{}
  rect2.length = 10  
  rect2.color  = "Red"
  fmt.Println(rect2)  // breadth skipped
    
  var rect3 = &rectangle{}
  (*rect3).breadth = 10
  (*rect3).color  = "Blue"
  fmt.Println(rect3)  // length skipped
}
```

In the above example, when you instantiating like `var rect1 = &rectangle{10,20,"Green"}` you have to pass all field values and as per there data type otherwise compiler will give error like too few values in struct initializer or cannot use XXXXXX (type string) as type int in field value.

The fields are initially assign as a default of data type then we can reassign the field by using dot operator `rect2.length`.

## Struct Exercise

### 1. Use field tags in the defination of Struct type

The following code snippet illustrates struct that id tagged with JSON annotation which can be interpreted by Go's JSON encoder and decoder:

```go
package main
 
import (
  "fmt"
  "encoding/json"
)
 
type Employee struct {
  FirstName  string `json:"firstname"`
  LastName   string `json:"lastname"`
  City string `json:"city"`
}
 
func main() {
  json_string := `
  {
    "firstname": "Rocky",
    "lastname": "Sting",
    "city": "London"
  }`

  emp1 := new(Employee)
  json.Unmarshal([]byte(json_string), emp1)
  fmt.Println(emp1)

  emp2 := new(Employee)
  emp2.FirstName = "Ramesh"
  emp2.LastName = "Soni"
  emp2.City = "Mumbai"
  jsonStr, _ := json.Marshal(emp2)
  fmt.Printf("%s\n", jsonStr)
}
```

In the above example, the tags are represented as raw string values which are wrapped within a pair of ``.

### 2. Nested Struc Type

Creating a Struct type using other struct types as the type for the fields of structs:

```go
package main
 
import "fmt"
 
func main() {
  type Salary struct {
    Basic, HRA, TA float64
  }

  type Employee struct {
    FirstName, LastName, Email string
    Age                        int
    MonthlySalary              []Salary
  }

  e := Employee{
    FirstName: "Mark",
    LastName:  "Jones",
    Email:     "mark@gmail.com",
    Age:       25,
    MonthlySalary: []Salary{
      Salary{
        Basic: 15000.00,
        HRA:   5000.00,
        TA:    2000.00,
      },
      Salary{
        Basic: 16000.00,
        HRA:   5000.00,
        TA:    2100.00,
      },
      Salary{
        Basic: 17000.00,
        HRA:   5000.00,
        TA:    2200.00,
      },
    },
  }
  fmt.Println(e.FirstName, e.LastName)
  fmt.Println(e.Age)
  fmt.Println(e.Email)
  fmt.Println(e.MonthlySalary[0])
  fmt.Println(e.MonthlySalary[1])
  fmt.Println(e.MonthlySalary[2])
}
```

The Employee struct has been expanded by adding a new field MonthlySalary, for which the type is specified as a slice of a struct named Salary. Using the MonthlySalary field, you can specify monthly salary for an employee.

### 3. Adding Methods to Struct Types

The following code snippet illustrates an example of Struc Type with Method:

```go
package main
 
import "fmt"
 
type Salary struct {
  Basic, HRA, TA float64
}
 
type Employee struct {
  FirstName, LastName, Email string
  Age                        int
  MonthlySalary              []Salary
}
 
func (e Employee) EmpInfo() string {
  fmt.Println(e.FirstName, e.LastName)
  fmt.Println(e.Age)
  fmt.Println(e.Email)
  for _, info := range e.MonthlySalary {
    fmt.Println("===================")
    fmt.Println(info.Basic)
    fmt.Println(info.HRA)
    fmt.Println(info.TA)
  }
  return "----------------------"
}

func main() {
  e := Employee{
    FirstName: "Mark",
    LastName:  "Jones",
    Email:     "mark@gmail.com",
    Age:       25,
    MonthlySalary: []Salary{
      Salary{
        Basic: 15000.00,
        HRA:   5000.00,
        TA:    2000.00,
      },
      Salary{
        Basic: 16000.00,
        HRA:   5000.00,
        TA:    2100.00,
      },
      Salary{
        Basic: 17000.00,
        HRA:   5000.00,
        TA:    2200.00,
      },
    },
  }
  fmt.Println(e.EmpInfo())
}
```

Go language type system allows you to add methods to struct types using a method receiver. The method receiver specifies which type has to associate a function as a method to that type.

# Interface

## What is an interface?

Go doesn't have classic Object Oriented concept of classes and inheritance. Interfaces in Go provide a way to specify the behavior of an object: do like this and reuse it.

In Go programming language Interfaces are types that just defines a set of methods (the method set), but these methods do not contain code. This mthods is never implemented by the interface type directly. Also an interface cannot contain variables. Go's interface type provides lot of extensibility and composability for your Go applications. Interface types express generalizations or abstractions about the behaviors of other types.

Interfaces are particularly useful as software projects grow and become more complex. They allow us to hide the incidental details of implementation the basic concept of Oops.

## Declaration of Interface Types

An interface type is declared with the keyword interface.

An interface is declared in the format:

```
type Name interface {
  Method1(param_list) return_type
  Method2(param_list) return_type
  …
}
```

where Name is an interface type.

```go
type Information interface {
  General()
  Attributes()
  Inventory()
}
```

The interface type `Information` is a contract for creating various `Product` types in a catalog.

The `Information` interface provides three behaviors in its contract: `General`,`Attributes` and `Inventory`.

## Implemention of Interface into Concrete Types

When a userdefined type implements the set of methods declared by an interface type the values of the user-defined type can be assigned back to the values of the interface type. In this process the value of the user-defined type getting stores into the interface value.

```go
type Product struct {
  Name, Description string
  Weight,Price float64
  Stock int
}
 
func (prd Product) General() {
  fmt.Printf("\n%s",prd.Name)
  fmt.Printf("\n%s\n",prd.Description)
  fmt.Println(prd.Price)
}
 
func (prd Product) Attributes(){
  fmt.Println(prd.Weight)
}
```

A struct Product is declared with fields for holding its state and methods implemented based on the behaviors defined in the Information interface.Multiple types can implement the same interface. A type that implements an interface can also have other functions. A type can implement many interfaces.

## Type Embedding

Go allows you to take existing types and both extend and change their behavior. This capability is important for code reuse and for changing the behavior of an existing type to suit a new need. This is accomplished through type embedding. It works by taking an existing type and declaring that type within the declaration of a new struct type. The type that is embedded is then called an inner type of the new outer type.

### A struct Mobile is declared in which the type Product is embedded.

```go
type Mobile struct{
  Product
  DisplayFeatures []string
  ProcessorFeatures []string
}
```

A struct Mobile is declared in which the type Product is embedded. The type Product is an implementation of the Information interface, the type Mobile is also an implementation of the Information interface. All fields and methods defined in the Type Product types are also available in the Mobile type. In addition to the embedded type of Product, the Mobile struct p

## Method Overriding

Since any user-defined type can implement any interface, method calls against an interface value are polymorphic in nature. The userdefined type in this relationship is often called a concrete type, since interface values have no concrete behavior without the implementation of the stored user-defined value.

```go
func (mob Mobile) Attributes(){
  mob.Product.Attributes()
  fmt.Println("\nDisplay Features:")
  for _, key := range mob.DisplayFeatures{
    fmt.Println(key)    
  }
  fmt.Println("\nProcessor Features:")
  for _, key := range mob.ProcessorFeatures{
    fmt.Println(key)    
  }
}
```

The Mobile struct is a concrete implementation of the Information interface than an Product type. The Product type is defined for type embedding for making a more concrete implementation of the Information interface, like the Mobile struct. At this moment, the Mobile struct uses the methods that were defined in the Product struct. Because the Mobile struct is more of a concrete implementation, it might have its own implementations for its methods. Here the Mobile struct might need to override the methods defined in the Product struct to provide extra functionalities.



