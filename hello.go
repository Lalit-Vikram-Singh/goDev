package main

import "fmt"

func swap(m1, m2 *int){

	temp := 0
	temp = *m1
	*m1 = *m2
	*m2 = temp
}

type car struct{
	name string
	drive int
}

func (c car) drivekilometer(){
	fmt.Println(c.drive)
}

func main() {
	m1, m2 := 2, 3
	fmt.Println(m1 , m2)
	swap(&m1, &m2)
	fmt.Println(m1 , m2)

	c := car{name: "honda", drive: 150,}
	c.drivekilometer()
}
