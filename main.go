package main

import (
  "strconv"
  "testing"
)

func FizzBuzz(number int) string {
  result := ""
  for i := 1; i <= number; i++ {
    if i%3 == 0 { result += "Fizz" }
    if i%5 == 0 { result += "Buzz" }
    if i%3 != 0 && i%5 != 0 { result += strconv.Itoa(i) }
    result += "\n"
  }
  return result
}

func TestFizzBuzz(t *testing.T) {
  got := FizzBuzz(15)
  want := "1\n2\nFizz\n4\nBuzz\nFizz\n7\n8\nFizz\nBuzz\n11\nFizz\n13\n14\nFizzBuzz\n"

  if got != want {
    t.Errorf("FizzBuzz(15) \n got: \n%v \n want: \n%v", got, want)
  }
}

func main() {
  tests := []testing.InternalTest{{"TestFizzBuzz", TestFizzBuzz}}
  matchAll := func(t string, pat string) (bool, error) { return true, nil }
  testing.Main(matchAll, tests, nil, nil)
}
