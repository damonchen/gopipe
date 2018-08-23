# gopipe
golang pipe implement 


## get the package


`go get -u github.com/damonchen/gopipe`


## usage

It is very simple for use, so code only.


```golang
package main

import (
	"fmt"
	"github.com/damonchen/gopipe"
)

func main() {
	p := pipe.NewPipe("cat", "pipe.go")
	p.P("grep", "func").P("awk", "{print $3}").P("grep", "string")
	
	stdout, err := p.Call()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(stdout.String())
}
```

It is the same as 

`cat pipe.go | grep func | awk '{print $3}' | grep string`


current, the code will output:

`string,`
