package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/data/azappconfig"
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

type Stest struct {
	Test string
}

func pFoo(p *int) {
	fmt.Printf("value: %#v\n", p) // value: (*int)(0xc420080008)
	fmt.Printf("addr: %#v\n", &p) // addr: (**int)(0xc420088028)
	*p = 11
}

func main() {
	fmt.Println("main")

	srvLookup()

	a := 10
	pa := &a
	fmt.Printf("address of a: %#v\n", pa)   // value: (*int)(0xc420080008)
	fmt.Printf("address of pa: %#v\n", &pa) // addr: (**int)(0xc420088018)
	pFoo(pa)

	fmt.Printf("address of a: %#v\n", pa)   // value: (*int)(0xc420080008)
	fmt.Printf("address of pa: %#v\n", &pa) // addr: (**int)(0xc420088018)

	fmt.Printf("value of a: %d\n", a)
	//Map()
	//ExampleGroup_justErrors()
	//ExampleGroup_noGoRoutine()
	//RunGeneric()
	credential, _ := azidentity.NewDefaultAzureCredential(nil)
	client, _ := azappconfig.NewClient("https://junbchenconfig-test.azconfig.io", credential, nil)
	var etag azcore.ETag = "AwjTinIYc1X8jH50-Gcvn6g0XwmPszOs660nnwv_e98"
	setting, err := client.GetSetting(context.Background(), "aaaa", &azappconfig.GetSettingOptions{Label: nil, OnlyIfChanged: &etag})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(*setting.ETag)

	// HttpHandler()
	// fmt.Println(&a)
	// var test Stest

	// if &test == nil {
	// 	fmt.Println("test is nil")
	// }
	// fmt.Println(test)
	// fmt.Println(&test)

	// fmt.Printf((*test).Test)

	// channelA := make(chan int)
	// go goroutineA(&channelA)
	// go goroutineB(&channelA)
	// fmt.Println("main is running...")
	// time.Sleep(time.Second * 10)
	// fmt.Println("complete")
	// credential, _ := azidentity.NewDefaultAzureCredential(nil)
	// client, _ := azappconfig.NewClient("https://junbchenconfig.azconfig.io", credential, nil)

	// keyFilter := "*"
	// // //labelFilter := "\x00"
	// labelFilter := "perftest1"

	// pager := client.NewListSettingsPager(azappconfig.SettingSelector{KeyFilter: &keyFilter, LabelFilter: &labelFilter}, nil)
	// ctx := context.TODO()
	// for pager.More() {

	// 	page, err := pager.NextPage(ctx)

	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}

	// 	for _, item := range page.Settings {
	// 		fmt.Println(*item.Key)
	// 	}
	// }

	// arrs := strings.Split("tsee\\,sfdsf,,est", ",")

	// strings.Contains("test", "test")

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
	// chanUsage()
}

func getFuncAddress(a string) {
	fmt.Println(&a)
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
