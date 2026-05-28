package main

import (
	"fmt"
	"test/functions"
)
func main() {
	v := "hello user"
	funct := functions.HelloUser(v)
	fmt.Println(funct)

	/* go mod init "name folder"
		go mod tidy
		go build .name
		если название с маленькой буквы то его нельзя будет импортировать и она будет
		анонимной что = инкапсуляция
	*/
}
