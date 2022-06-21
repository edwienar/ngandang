package main

func main() {
	e := routes.Index()
	e.Logger.Fatal(e.Start(":1234"))
}
