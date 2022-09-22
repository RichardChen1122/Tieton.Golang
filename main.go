package main

import (
	"errors"
	"fmt"
)

type Controller struct {
	name string
	age  int
}

func (controller *Controller) TestC() {
	controller.name = "C"
	fmt.Println(controller.name)
}

func (controller Controller) TestD() {
	controller.name = "D"
	fmt.Println(controller.name)
}

// func goroutineA(channel *chan int) {
// 	fmt.Println("From goroutineA")

// 	for {
// 		time.Sleep(time.Second * 1)
// 		fmt.Println("goroutineA is running..")
// 		select {
// 		case t := <-*channel:
// 			fmt.Printf("goroutineA get %d", t)
// 			fmt.Println()
// 			break
// 		default:
// 			fmt.Println("goroutineA default..")
// 		}
// 	}
// }

// func goroutineB(channel *chan int) {
// 	for i := 0; i < 3; i++ {
// 		time.Sleep(time.Second * 1)
// 		fmt.Println("goroutineB is running..")
// 	}
// 	*channel <- 45
// 	fmt.Println("From goroutineB")
// 	fmt.Println()
// }

func foo() (err error) {
	defer func() {
		err = errors.New("defer") // change value at the very last moment
	}()
	return errors.New("dddd")
}

func setup(test *map[string]string, src *map[string]string) {
	if *test == nil {
		fmt.Println("is nil")
		test = src
	}
	fmt.Println((*test)["key"])
}

func main() {
	// fmt.Println("main")

	// channelA := make(chan int)
	// go goroutineA(&channelA)
	// go goroutineB(&channelA)
	// fmt.Println("main is running...")
	// time.Sleep(time.Second * 10)
	// fmt.Println("complete")
	// credential, _ := azidentity.NewDefaultAzureCredential(nil)
	// client, _ :=azappconfig.NewClient("https://junbchenconfig.azconfig.io", credential, nil)

	// arrs := strings.Split("tsee\\,sfdsf,,est", ",")

	// strings.Contains("test","test")

	// for i := 0; i < len(arrs); i++ {
	// 	fmt.Println(arrs[i])
	// }

	// fmt.Println(foo())
	// people := &People{
	// 	Name: "Linda Chen",
	// 	Age:  3,
	// }

	// lele := &Student{
	// 	People: *people,
	// 	Score:  100,
	// 	Class:  3,
	// }

	// peo := People{
	// 	Name: "Richard",
	// 	Age:  33,
	// }

	// richard := Educator{
	// 	Job: "Programmer",
	// 	Student: Student{
	// 		Score:  99,
	// 		People: peo,
	// 	},
	// }

	// richard.Student.SayName()
	// lele.SayName()
	//Demo(lele)

	// yu := make(map[string]string)

	// yu["key"] = "value"

	// setup(nil, &yu)
	chanUsage()
}

func Demo(student StudentInterface) {
	student.SayName()
}

type People struct {
	Name string
	Age  int
}

type Student struct {
	Score int
	Class int
	People
}

type Educator struct {
	StudentInterface
	Student
	Job string
}

type PeopleInterface interface {
	SayName() string
}

type StudentInterface interface {
	PeopleInterface
	SayScore() int
}

func (p *People) SayName() string {
	fmt.Println(p.Name)
	return p.Name
}

func (s *Student) SayScore() int {
	fmt.Println(s.Score)
	return s.Score
}

// func (s *Student) SayName() string {
// 	fmt.Println("This is from Student" + strconv.Itoa(s.Score))
// 	return strconv.Itoa(s.Score)
// }

// func (s *Educator) SayScore() int {
// 	fmt.Println("This is from Educator " + strconv.Itoa(s.Score))
// 	return s.Score
// }

// func (s *Educator) SayName() string {
// 	fmt.Println("This is from Educator " + s.Name)
// 	return s.Name
// }
