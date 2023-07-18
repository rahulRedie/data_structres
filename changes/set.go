// set in go

//add, exists, remove, size

package main

import "fmt"

type customSet struct {
	container map[string]struct{}
}

func (c *customSet) Add(key string) {
	c.container[key] = struct{}{}
}

func (c *customSet) exists(key string) bool {
	if _, contains := c.container[key]; contains {
		return true
	}
	return false
}

func (c *customSet) remove(key string) error {
	if c.exists(key) {
		delete(c.container, key)
	}
	return fmt.Errorf("element not found")
}

func makeSet() *customSet {
	return &customSet{
		container: map[string]struct{}{},
	}
}

func (c *customSet) size() int {
	return len(c.container)
}
func main() {
	setContainer := makeSet()

	setContainer.Add("a")
	setContainer.Add("b")
	setContainer.Add("c")
	fmt.Println(setContainer.size())

	fmt.Println("checking a exists", setContainer.exists("a"))
	err := setContainer.remove("a")
	if err != nil {
		fmt.Println("checking a exists", setContainer.exists("a"))
	}
}
