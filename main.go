package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Input your score : ")
	i, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)                //줄바꿈 제거
	score, _ := strconv.ParseInt(i, 10, 32) //정수형 10진수 변환
	if score >= 60 {
		fmt.Println("A")
		//status := "passing"

	} else {
		//status := "failing"
		fmt.Println("bcdf")
	}

}
