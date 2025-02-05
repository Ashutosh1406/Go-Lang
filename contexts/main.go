package main

import (
	"context"
	"fmt"
	"time"
)

// context starter

func doSomething(ctx context.Context) {
	fmt.Printf("doing something with %s, Context address: %p\n", ctx.Value("mykey"), ctx) //first print

	//creating new context
	anotherContext := context.WithValue(ctx, "mykey", "DrChase")
	fmt.Printf("New context in doSomething, Address: %p\n", anotherContext)
	doAnother(anotherContext)

	fmt.Printf("doing something 2nd time with %s, Context address: %p\n", ctx.Value("mykey"), ctx) //third print
}

func doAnother(ctx context.Context) {
	fmt.Printf("doing Anotherwith,  %s, Context address: %p\n", ctx.Value("mykey"), ctx) //second print
}

func main() {

	// ctx := context.TODO()
	// doSomething(ctx)

	//context.Background() is best to use by default if none of the information is available
	//context.Todo() also returns the same empty context.Context but is generally used to start a context

	ctx := context.Background()
	ctx = context.WithValue(ctx, "mykey", "DrHouse")
	doSomething(ctx)

}

//code 2 [merge main function for using different context functions like WithCancel, WithDeadline, WithTimeOut]
// context with cancel for webServers and for saving and efficiently managing computer cpu-cycles, memory runtime and various
// other resources

func doSomething(ctx context.Context) {
	ctx, cancelCtx := context.WithCancel(ctx)

	printCh := make(chan int)       //making new channel named as printCh that is read only channel
	go doAnotherThing(ctx, printCh) //spinning another go routine

	for num := 1; num <= 1000; num++ {
		printCh <- num
	}
	cancelCtx()

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("DoSomething!!! FINISHED\n")
}

func doAnotherThing(ctx context.Context, printCh <-chan int) {
	fmt.Println("entered in doAnother")
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("Do another error %s\n", err)
			}
			fmt.Printf("DoAnother: Finished!!\n")
			return
		case num := <-printCh:
			fmt.Printf("DoAnother: %d\n", num)
		}
	}
}

func main() {
	ctx := context.Background()
	doSomething(ctx)
}
