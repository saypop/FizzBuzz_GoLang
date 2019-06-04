# FizzBuzz in Go

This was my first attempt at TDDing FizzBuzz in Go.

I am still learning about the folder structuring in this very opinionated language.

For now I have just written my tests in the same file that my code is in.

All in all I'm happy with the result, I have TDD'ed the code and it works, to run it enter this into the command line: `$ go run main.go`.

I need to improve the structuring though so the way to run my tests is convoluted, to do this you need to go into the main.go file and comment out the bottom line in the main function and uncomment the rest then enter this into the command line: `$ go run main.go`

In the next iteration I am going to structure thing so I can run tests and the program. 
