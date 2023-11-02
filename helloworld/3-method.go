package main

type Person struct {
	name string
	sex  uint32
}

func (p Person) Run(distance string) {
	println(p.name, "run", distance)
}

func main() {
	p := &Person{
		name: "Tom",
		sex:  1,
	}
	p.Run("100")
}
