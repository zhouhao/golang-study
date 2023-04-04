package main

type People struct {
	name string
	age  int
}

func main() {
	p := People{"Tom", 18}
	println(p.name)

	peopleArr := make([]People, 2)
	peopleArr[0] = People{"Tom", 18}
	peopleArr[1] = People{"Jerry", 20}

	for _, v := range peopleArr {
		println(v.name)
	}
}
