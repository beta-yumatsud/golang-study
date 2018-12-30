package main
import(
	"fmt"
	"./animals"
)

func main(){
	fmt.Println("Hello, World!")
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
	fmt.Println(AppName())
}
