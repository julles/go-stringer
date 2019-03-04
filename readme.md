# Stringer for GoLang
Go Stringer is GoLang package for string manipulation with an expresive syntax.
## How to Install ?

``` sh
go get github.com/julles/go-stringer
```

### Or if you using DEP 

``` sh
dep ensure -add github.com/julles/go-stringer

dep ensure

```

# Example 

### Basic

Import package

``` go
import (
	..
	s "github.com/julles/go-stringer"
)
```

Here is the basic usage

``` go
var kata s.Words = "REZA"
kata.LowerFirst()
fmt.Println(kata) // the output "rEZA"

var kata2 s.Words = "REZA"

kata2.Reverse().
	Replace("A","I")

fmt.Println(kata2) // the output "IZER"

var kata3 s.Words = "Muhamad Reza Abdul Rohim"
kata3.CamelCase(" ")
fmt.Println(kata3) // the output "muhamadRezaAbdulRohim"

```

Using chaining method

``` go
var kata s.Words = "REZA"

kata.LowerFirst().
	LowerLast().
	Reverse().
	Replace("a", "i").
	Repeat(2)

fmt.Println(kata) // the output "iZEriZEr"

```

## Available Methods

| Methods   |
|----------|
| .CamelCase(separator string) |
| .Lower() |
| .LowerFirst() |
| .LowerLast() |
| .Repeat(count int) |
| .Replace(search string,replace string) |
| .Reverse() |
| .Substring(start int , end int) |
| .Upper() |
| .UpperFirst() |
| .UpperLast() |

## License

https://reza.mit-license.org/

