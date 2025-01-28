package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

//import "encoding/json"
//import "example/hello/MyStruct"

func summ(first int, second int) (sum int) {

	defer func() {
		sum *= 2
	}()

	sum = first + second
	return
}

func Taker(first int, second int, pidor func(int, int) int) int {

	return pidor(first, second)
}

func returner() func(int, int) int {

	return summ
}

func Age_Checker(age int) bool {
	return age < 18
}

type Person struct {
	name string
	age  int
}

type User struct {
	ID    int
	Name  string
	Email string
}

type Noob struct {
	Name string
}

var user User = User{ID: 1, Name: "Pidoras", Email: "sevaka777@gmail.com"}

type tester interface {
	GetName() string
	GetEmail() string
}

func (s *User) GetName() string {

	return s.Name
}

func (s *User) GetEmail() string {

	return s.Email
}

func (s *Noob) GetName() string {
	return s.Name
}

func (s *Noob) GetEmail() string {
	return "Noob can't have Email"
}

func ShowAllElements(slices []int) {
	slices = append(slices, 20, 50, 60)
	for _, value := range slices {
		fmt.Println(value)
	}
	slices[0] = 2
}

func IndexChecker(usermap []User) map[int]struct{} {
	mymap := make(map[int]struct{}, len(usermap))

	for _, value := range usermap {
		if _, ok := mymap[value.ID]; !ok {
			mymap[value.ID] = struct{}{}
		}
	}

	return mymap
}

func faktorial(value float64) float64 {
	if value == 1 {
		return 1
	}

	return value * faktorial(value-1)
}

func makeChan() <-chan string {
	responseChan := make(chan string)

	go func() {
		time.Sleep(time.Second)
		responseChan <- "Promise is working"
	}()

	return responseChan
}

func worker(ctx context.Context, toProcess <-chan int, processed chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return

		case value, ok := <-toProcess:
			if !ok {
				return
			}
			time.Sleep(time.Millisecond)
			processed <- value * value

		}
	}
}

func workerpool() {
	group := &sync.WaitGroup{}

	numbersToProcess, ProcessedNumbers := make(chan int, 5), make(chan int, 5)

	Test_Context, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < runtime.NumCPU(); i++ {
		group.Add(1)
		go func() {
			worker(Test_Context, numbersToProcess, ProcessedNumbers)
			group.Done()
		}()
	}

	go func() {
		for i := 0; i < 1000; i++ {
			numbersToProcess <- i
		}
		close(numbersToProcess)
	}()

	go func() {
		group.Wait()
		close(ProcessedNumbers)
	}()

	for value := range ProcessedNumbers {
		fmt.Println(value)
	}

}

func chanAsPromise(value int) <-chan int {

	time_chan := make(chan int, 1)

	go func() {
		time.Sleep(time.Second * 3)
		time_chan <- value * 3
	}()
	return time_chan
}

func chanAsMutex() {
	var counter int = 0
	wg := &sync.WaitGroup{}
	ctx := make(chan struct{}, 1)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx <- struct{}{}
			counter++
			<-ctx
		}()

	}
	wg.Wait()
	fmt.Println(counter)
}

type test interface {
	talker() string
}

type HumanMy struct {
	name string
}

type Animal struct {
	age int
}

func (value HumanMy) talker() string {
	return "Hello I am " + value.name
}

func (value Animal) talker() string {
	return fmt.Sprintf("I am %d", value.age)
}

func main() {
	Andrei := HumanMy{name: "Andrei"}
	Wolf := Animal{age: 4}

	var Dota2 test = Andrei
	fmt.Println(Dota2.talker())
	Dota2 = Wolf
	fmt.Println(Dota2.talker())

	var number int = 0
	myMutex := sync.Mutex{}

	for i := 0; i < 100; i++ {
		go func(i int) {
			myMutex.Lock()
			defer myMutex.Unlock() // defer освобождает мьютекс даже в случае паники

			if number == 0 {
				number = 1
				fmt.Printf("It was made by goroutine number: %d", i)
				fmt.Println("\n")
			}
		}(i)
	}

	// chanAsMutex()

	// muuchan := chanAsPromise(10)

	// fmt.Println("It works well")

	// fmt.Println(<-muuchan)

	// workerpool()

	ctx := context.Background()

	context_value := context.WithValue(ctx, "name", "pidor")

	fmt.Println(context_value.Value("name"))

	deadline, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
	defer cancel()
	fmt.Println(deadline.Deadline())
	fmt.Println(deadline.Err())
	fmt.Println(<-deadline.Done())
	fmt.Println(deadline.Err())

	fmt.Println("Started")

	// ResultChan := make(chan int)
	// timer := time.After(time.Second)
	// go func(){
	// 	defer close(ResultChan)

	// 	for i := 0; i < 1000; i++{
	// 		select{
	// 		case <- timer:
	// 			fmt.Printf("Time's up")
	// 			return
	// 		default:
	// 			time.Sleep(time.Millisecond)
	// 			fmt.Println("Is deadlock?")
	// 			ResultChan <- i
	// 			fmt.Println("No")

	// 		}
	// 	}

	// }()

	// for{
	// 	v,ok := <- ResultChan
	// 	fmt.Println(v,"\n")
	// 	if(ok == false){
	// 		 break
	// 	}
	// }

	// unbufchan := make(chan int)

	// go func(){
	// 	time.Sleep(time.Second)
	// 	unbufchan <- 1
	// }()

	// select{
	// case val := <- unbufchan:
	// 	fmt.Println("blocking reading: ",val)
	// case <- time.After(time.Millisecond * 999):
	// 	fmt.Println("Time's up")
	// }
	// channel := make(chan int,5)
	// myslice := make([]int,0,5)
	// myslice = append(myslice,2,5,4,8,10)
	// go func(b int){
	// 	for _,value := range myslice{
	// 		channel <- value
	// 	}
	// 	close(channel)
	// }(2)

	// for{
	// 	v,ok := <- channel
	// 	fmt.Println("Value: ",v)
	// 	if(ok == false){
	// 		break
	// 	}
	// }

	// for i := 0; i < 10; i++ {
	// 	defer fmt.Println("first",i)
	// }

	// for i := 0;i<10;i++{
	// 	defer func(){
	// 		fmt.Println("second",i)
	// 	}()
	// }

	Users := []User{
		{
			50,
			"Seva",
			"no",
		},
		{
			50,
			"Daniil",
			"no",
		},
		{
			35,
			"Kirill",
			"no",
		},
	}

	User_map := make(map[int][]User, len(Users))
	// for _,value := range Users{
	// 	if _,ok := User_map[value.ID]; !ok{
	// 		User_map[value.ID] = value
	// 	}
	// }

	for _, user := range Users {
		User_map[user.ID] = append(User_map[user.ID], user)
	}

	for _, value := range User_map {
		for _, value2 := range value {
			fmt.Println(value2.ID, " ", value2.Name)
		}
	}

	var usermap []User = []User{
		{ID: 52, Name: "Seva", Email: "0"}, {ID: 52, Name: "Artur", Email: "0"}, {ID: 48, Name: "Ronaldo", Email: "0"}}

	fmt.Println(IndexChecker(usermap))

	mymap := make(map[int64]string, 5)

	mymap[1] = "Seva"
	mymap[0] = "Daniil"
	mymap[2] = "Artur"
	mymap[3] = "Dimon"
	mymap[5] = "Dasha"

	fmt.Println(mymap[1])

	slices := []int{1, 2, 4, 6, 8}

	Short_Slices := append(slices[0:2], slices[3:5]...)

	fmt.Println(Short_Slices)

	slices2 := []int{1, 2, 4, 6, 8}

	Short_Slices = slices2[:2+copy(slices2[2:5], slices2[3:5])]

	fmt.Println(Short_Slices)

	// var slices []int = make([]int, 0, 5)

	// for i := 0;i<cap(slices);i++{
	// 	slices = append(slices,i*2)
	// }
	// slices = append(slices,120)

	// ShowAllElements(slices)

	// //fmt.Printf("Length %d Capacity %d", len(slices),cap(slices))

	// for _,value := range slices{
	// 	fmt.Println(value)
	// }

	//builder := mystruct.Builder{mystruct.Person{"Anton",40},mystruct.WoodBuilder{"WoodBuilder"},mystruct.StoneBuilder{"None"}}

	// fmt.Println(builder.Get_Job_Type())

	// unnamed := &Noob{"\nSomeOne"}
	// var checker tester = unnamed
	// switch v := checker.(type){
	// case *User:
	// 	fmt.Printf("%T, %#v, %#v",v,v,v.Name)
	// case *Noob:
	// 	fmt.Printf("%#v",v.GetEmail())
	// }

	// fmt.Println(user.GetName())
	// fmt.Println(user.GetEmail())

	// Seva := struct {
	// 	Name, LastName, BirthDate string
	// }{
	// 	Name: "Seva",
	// 	LastName: "Pashetov",
	// 	BirthDate: "15.08.2005",
	// }

	// fmt.Printf("%#v %#v %#v \n", Seva, Seva.Name, Seva.LastName)

	// userJSON, err := json.Marshal(user)

	// if err != nil {
	// 	fmt.Println("Mistake of Marshaling", err)
	// }

	// fmt.Println(string(userJSON))

	// // age := 17

	// // // if AgeCheck := Age_Checker(age); AgeCheck == true{
	// // // 	fmt.Println("You are too young")
	// // // }	else{
	// // // 		fmt.Println("You are an adult")
	// // // 	}

	// var number int = 1

	// switch number{
	// case 1:
	// 	fmt.Println("Number is 1")
	// 	fallthrough
	// case 2:
	// 	fmt.Println("Number is 2")
	// case 3:
	// 	fmt.Println("Number is 3")
	// case 4:
	// 	fmt.Println("Number is 4")
	// case 5:
	// 	fmt.Println("Number is 5")
	// default:
	// 	fmt.Println("Another number is here")
	// }

	// var my_num int = 5
	// var pointer *int = &my_num
	// fmt.Printf("%T %#v %#v \n", pointer, pointer, *pointer)

	// fmt.Printf("%T", user)

	// Label1:
	// 	for i:=0;i<10;i++{
	// 		for j := 0; j < 10; j++ {
	// 			fmt.Println(i," ",j)
	// 			if i > 5 {
	// 			continue Label1
	// 			}
	// 			fmt.Println(" No")
	// 		}
	// 	}

}
