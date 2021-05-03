package main
 import "fmt"

 func main()  {
	 name:= "eum602";
	 tpl:= `
	 <!DOCTYPE html>
	 <html lan="en">
	 <head>
	 <meta charset="UTF-8">
	 <title>Hello World!</title>
	 </head>
	 <body>
	 <h1>` + name + `</h1>
	 </body> 
	 </html>
	 `

	 fmt.Println(tpl);
 }