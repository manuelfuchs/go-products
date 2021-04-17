package main

func main() {
	a := App{}
	a.InitializeFromEnvironment()

	a.Run(80)
}
