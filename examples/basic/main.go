//go:generate go run github.com/wesovilabs/goa main.go

package main

import (
	"fmt"
	"github.com/wesovilabs/goa/context"
	"github.com/wesovilabs/goa/examples/basic/aspects"
	"github.com/wesovilabs/goa/goa"
	"strings"
)

func Goa() *goa.Aspects {
	return goa.New().
		WithAspect("*.*(*)*", aspects.LogAspect).
		WithAspect("*.*(string)*", ToUpper)

}

func main() {
	fmt.Println("\n[test]")
	test1("pepe")
	fmt.Println("\n[test2]")
	test2(2)
	fmt.Println("\n[test3]")
	test3("John", "Doe")
	fmt.Println("----------")
	aspects.PrintCounter()

}

func test1(name string) {
	fmt.Printf("    name is %v\n", name)
}

func test2(value int) {
	fmt.Printf("    value is %v\n", value)
}

func test3(name string, surname string) {
	fmt.Printf("    %s %s\n", name, surname)
}

func AspectMultiply(ctx *context.Ctx) {
	ctx.In().List()[0].Update(ctx.In().List()[0].Value().(int) * 20)
}

func ToUpper(ctx *context.Ctx) {
	ctx.In().List()[0].Update(strings.ToUpper(ctx.In().List()[0].Value().(string)))
}

func ToLower(ctx *context.Ctx) {
	ctx.In().List()[0].Update(strings.ToLower(ctx.In().List()[0].Value().(string)))
	ctx.In().List()[1].Update(strings.ToLower(ctx.In().List()[1].Value().(string)))
}
