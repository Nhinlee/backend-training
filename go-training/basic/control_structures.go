package basic

type ControlStructures struct{}

// Function demonstrates how to use if-else statement in Go.
func (c *ControlStructures) IfElse() {
	// Declare a variable.
	x := 10

	// If-else statement.
	if x > 0 {
		println("x is greater than 0")
	} else {
		println("x is less than or equal to 0")
	}
}

// Function demonstrates how to use switch statement in Go.
func (c *ControlStructures) Switch() {
	// Declare a variable.
	x := 10

	// Switch statement.
	switch x {
	case 1:
		println("x is 1")
	case 2:
		println("x is 2")
	default:
		println("x is neither 1 nor 2")
	}

	// Switch statement with multiple cases.
	switch x {
	case 1, 2:
		println("x is either 1 or 2")
	default:
		println("x is neither 1 nor 2")
	}

	// Switch statement with expression.
	switch {
	case x > 0:
		println("x is greater than 0")
	case x < 0:
		println("x is less than 0")
	default:
		println("x is equal to 0")
	}

	// Switch statement with fallthrough.
	switch x {
	case 1:
		println("x is 1")
		fallthrough
	case 2:
		println("x is 2")
	default:
		println("x is neither 1 nor 2")
	}

	// Switch statement with type assertion.
	var i interface{} = 1
	switch i.(type) {
	case int:
		println("i is an integer")
	case string:
		println("i is a string")
	default:
		println("i is neither an integer nor a string")
	}
}

// Function demonstrates how to use for loop in Go.
func (c *ControlStructures) ForLoop() {
	// For loop with a single condition.
	for i := 0; i < 3; i++ {
		println(i)
	}

	// For loop with a single condition.
	i := 0
	for i < 3 {
		println(i)
		i++
	}

	// For loop with no condition.
	i = 0
	for {
		if i >= 3 {
			break
		}
		println(i)
		i++
	}
}
