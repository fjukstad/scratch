# scratch
Simple pkg for getting information about scratch projects at scratch.mit.edu

# example 
```
package main

import (
	"github.com/fjukstad/scratch"
)

func main() {
	id := "86536890"
	p, err := scratch.GetProject(id)
   ...
}

``` 
