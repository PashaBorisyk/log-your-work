package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	fileName := getFileName()
	file :=getFile(fileName)
	startLogging(file)

}

func getFileName() string {

	now := time.Now()
	month := now.Month().String()
	day := strconv.Itoa(now.Day())
	dayOfWeek := now.Weekday().String()

	fileName := "./"+ dayOfWeek +"("+ day+" "+ month + ").log"
	return fileName
}

func getFile(fileName string) (file *os.File) {

	_, err := os.Stat(fileName)
	if err != nil && os.IsNotExist(err) {
		log.Println("No file found for today. Creating new one")
		file,err = os.Create(fileName)
	} else {
		file,err = os.OpenFile(fileName,os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	}

	if err != nil {
		log.Println("Error while opening or creating a file")
		log.Fatal(err)
	}
	return file
}

func startLogging(file *os.File){

	fmt.Println("Starting logging you work")
	fmt.Println("=========================")

	log.SetOutput(file)

	log.Println("Setting output to "+ file.Name())
	reader := bufio.NewReader(os.Stdin)
	for  {
		fmt.Print("->  ")
		text,_ := reader.ReadString('\n')
		text = strings.Replace(text,"\n","",-1)
		log.Println(text)
		fmt.Println("Commited (", time.Now().Format(time.Kitchen), ")")
		fmt.Println("")
	}
}
