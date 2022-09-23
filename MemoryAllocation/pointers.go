package main

import "fmt"

/*
	There are different types of memory allocation.

	when you declare a variable or function of any piece of data
	the computer will allocate a piece of memory to that point and
	then associate a memory address for where that data will be stored.

	there are different types of memory allocation and different data types associated with it

	these are.

	Heap:


	Pointers:
		a Pointer is the memory address rather then the value itself.
		it can introduce concepts such as lazy loading where instead of
		reading a piece of data you just know where it lives and manipulate/read
		the actual value when it is needed.
		This can be far more efficient to pass around a large program rather then
		passing large chunks of data at at a time.

	Stack of memory (not as efficient but very safe):
		goroutines -> a light weight thread takes place.
		when a function call or variables is declared a part
		of the stack is given to the call. this is called a frame.
		the goroutine can only operate within this frame
		(the amount of given memory for the function call)
		this gives immutability giving a safe programme operation.

		however this means you would need to copy parameters that might come from
		a function call outside of the frame. i.e another piece of the stack with its
		own immutable frame. This is where pointers become handy for efficient operations.



*/

func assignValues () {
	i, j := 47, 2701;
	// below the actual values will be printed out to the terminal
	fmt.Println(i,j);

	// below the memory addresses where these number are stored will appear on the terminal
	// Note the & represents memory address or id of the memory
	fmt.Println(&i, &j);

	// These two variables have been declared as pointers
	pointerI, pointerJ := &i, &j;
	// now only the memory address is associated with the variables pointerI and pointerJ
	fmt.Println(pointerI, pointerJ);


	// the * operator or pointer type
	// so if you *int -> this is a pointer type with base of an integer

	// to print a value of a pointer you use the * operand to get the actual value
	// this is also known as dereferencing
	fmt.Println(*pointerI, *pointerJ);


	// mutating a pointer --> how to change or update pointer values
	/*
		use the pointer operand to mutate the value --> if you did not include you'll get an error
		the error would be because the memory address is a char array or string - and you can't divide an int by this type
		we need the value at the memory address -- not the memory address.
	*/
	*pointerJ = *pointerJ / 9

	fmt.Println(*pointerJ)
}

// inefficient function below using standard or raw variables
func squareVal(v int) {
	v *= v
	fmt.Println(v, &v);
}


// efficient function using pointers
func squarePointer(p *int) {
	*p *= *p;
	fmt.Println(*p, p);
}

func main() {
	// calling a frame with a regular value and function
	a:=4
	squareVal(a);
	fmt.Println(a);

	// calling a function with a pointer



}

