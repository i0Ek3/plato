# plato

plato is a small compiler which compile Lisp expression to C expression.


## Getting Started

### TL;DR

Just run command `make`.

### Install

`$ go get github.com/i0Ek3/plato`

### Import

```Go
package main

import (
    "fmt"
    "github.com/i0Ek3/plato"
)

func main() {
    expression := "(add 10 (subtract 10 (add 6 (subtract 4 2))))"
    result := plato.Compiler(expression)
    fmt.Println(result) // add(10, subtract(10, add(6, subtract(4, 2))));
}
```


### Test

```Shell
$ go test -v .
$ go test -bench .
```


More details please check comments in the source code.

## Credit

- [https://github.com/jamiebuilds](https://github.com/jamiebuilds)
- [https://github.com/hazbo](https://github.com/hazbo)
