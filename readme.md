# Pocker

`pocker` is a program which helps to find best 5 card pocker hand.

## Installation

```
go get -u github.com/firstrow/pocker/cmd/pocker
```


## Usage

To process many hands and decks cli application can be used.
```
pocker ~/input.txt
```

See [data/input.txt](https://github.com/firstrow/pocker/tree/master/data/input.txt) for example input file.

Also, `pocker` library can be used in your golang project:

``` go
package main

import (
	"fmt"
	"github.com/firstrow/pocker"
)

func main() {
	hand := []string{"AC", "2D", "6C", "3S", "KD"}
	deck := []string{"5S", "4D", "KS", "AS", "4C"}

	best := pocker.BestHand(hand, deck)
	fmt.Println("Best hand:", best) // straight
}
```
