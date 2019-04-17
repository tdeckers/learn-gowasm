package main

import (
	"fmt"

	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/html"
)

func main() {
	// get the element
	divElement := webapi.GetWindow().Document().GetElementById("app")
	app := html.HTMLDivElementFromJS(divElement)
	app.SetInnerHTML("Hello, WebAssembly!")

	btnElement := webapi.GetWindow().Document().GetElementById("button")
	btn := html.HTMLButtonElementFromJS(btnElement)

	webapi.GetWindow().Document().SetCookie("cookie content")

	// register a callback that will display a counter
	count := 1
	callback := domcore.EventHandlerToJS(func(event *domcore.Event) interface{} {
		btn.SetInnerText(fmt.Sprint("Count: ", count))
		count++
		return nil
	})
	btn.SetOnclick(callback)

	fmt.Println("Hello, WebAssembly World!") // goes to console.

	// prevent to program to terminate
	c := make(chan struct{}, 0)
	<-c
}
