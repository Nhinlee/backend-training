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



## OOP in Golang

## Concurrency
