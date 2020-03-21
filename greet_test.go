package greet

import (
	"fmt"
	"testing"
)

func TestSayHello(t *testing.T) {
	want := "Hello, World!"
	if got := SayHello(); got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func ExampleSayHello() {
	fmt.Println(SayHello())
	// Output: Hello, World!
}
