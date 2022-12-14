# Golang



## Basics

### Imports

"Factored" import statement.
```
import (
    "fmt"
    "math"
)
``` 
### Exported names

Capital letter
```
import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}
```

### Functions

```
func add(x, y int) int {
	return x + y
}
```

### Multiple results

```
func swap(x, y string) (string, string, int) {
	return y, x, 10
}
```

### Basic types

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

### Type conversions

```
i := 42
f := float64(i)
u := uint(f)
```

Explicit conversion <> C++

### Looping - only For
Go has only one looping construct, the **for** loop.

```
for i := 0; i < 10; i++ {
	sum += i
}
```
```
// while in other languages
for sum < 1000 {
	sum += sum
}
```
```
for {
	// forever loop (can break)
}
```

### If
```
// If with a short statement
if v := math.Pow(x, n); v < lim {
	return v
}
```
### Switch
```
// No need constant value
// Auto add break

func main() {
	fmt.Print("Go runs on ")
	val := 2
	val2 := 2
	switch val {
		case 1: fmt.Print("Okiela")
		case val2: fmt.Print("Okiela 2")
	}
}
```
```
// Switch with no condition
// clean way to write long if-then-else chains

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

### Defer
```
// A defer statement defers the execution of a function until the surrounding function returns.

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```
```
// Stacking defers

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```




## More types

### Pointers


```
// Unlike C/C++, Go has no pointer arithmetic.

// in C/C++:
int *c = new int[5];
c++

// in Go:
var c [5]int
c++ // compile error
```

### Structs
A struct is a collection of fields.

```
type Vertex struct {
	X int
	Y int
	Z float64
}
```

```
// Pointers to structs

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
```
```
// Struct Literals

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)
```

### Arrays - fixed size
```
var a [2]int
primes := [6]int{2, 3, 5, 7, 11, 13}
```

### Slices - dynamically-sized
In practice, slices are much more common than arrays.

```
func main() {
	var aa [5]int
	fmt.Println(aa)

	a := make([]int, 0, 5)
	printSlice("a", a)

	b := make([]int, 1, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
```

Detail => [Go Slices: usage and internals](https://go.dev/blog/slices-intro)

### Range

```
for i, v := range pow {
	fmt.Printf("2**%d = %d\n", i, v)
}
```

### Maps

```
m := make(map[string]int)

var m = map[string]int{
	"Bell Labs": 1,
	"Google":    2,
}

var m map[string]int // nil
```



## OOP in Golang

### Methods
```
// value receiver
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// pointer receiver
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

Methods and pointer indirection
```
// Function 

var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK

// Pointer indirection (both pointer & value receiver is accept)

var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```


NOTE:

There are two reasons to use a pointer receiver.
- Can modify the value
- Avoid copying the value on each method call

all methods on a given type should have either value or pointer receivers, but not a mixture of both.


### Interfaces

```
// Interfaces are implemented implicitly

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I
	describe(i)
	// i.M() - error
	
	var t *T
	i = t
	describe(i)
	i.M() // null pointer in other languages

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```
```
// Dynamic type in other languages
func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}
```

### Types

```
// Type assertions
t, ok := i.(T)
```

```
// Type switches
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```

## Concurrency

<mark>**Don't communicate by sharing memory, share memory by communicating**</mark>

### **Concurrency is not parallelism**
- Blogs: https://go.dev/blog/waza-talk
- Slides: https://go.dev/talks/2012/waza.slide#1

### **Concurrency patterns**
- Generator: function that returns a channel
- Multiplexing
- Restoring sequencing
- Select

### **Advanced concurrency**