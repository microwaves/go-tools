# Randomizer

Randomizer generate random strings. To use it, install the package:

```
$ go get github.com/microwaves/go-utils/randomizer
```

Then, import on your code:

```
package main

import (
    "fmt"
    
    "github.com/microwaves/go-utils/randomizer"
)

func main() {
    n := 10
    fmt.Println(randomizer.GenerateRandomString(n))
}
```

Where `n` is the amount of characters for the `string` output.

## Author

Stephano Zanzin <sz@shitty.pizza>
