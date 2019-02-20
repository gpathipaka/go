package main

import "fmt"

//ByteCouner is ...
type ByteCouner int

func (c *ByteCouner) Write(p []byte) (n int, err error) {
	*c += ByteCouner(len(p))
	return len(p), nil
}

//Stringer is
type Stringer interface {
	String() string
}

func main() {
	var c ByteCouner
	c.Write([]byte("Hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

	fmt.Printf("")
}
