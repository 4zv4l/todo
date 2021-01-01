package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

func help() { // show how to use todo
	clear()
	fmt.Println(`Usage :
	write something to add in the list
	write a number in the list to delete it
	write -1 to end the program
	`)
	fmt.Print("Press Enter to continue...")
	fmt.Scanln()
}

func remove(slice []string, s int) []string { // remove an element from a slice
	return append(slice[:s], slice[s+1:]...)
}

func clear() { // clear the screen depend on the os
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func load() []string {
	path, _ := os.UserHomeDir() // make the path for the save file
	path += "/.todo.mine"
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0600) //open the file in read only
	if err != nil {
		print(err)
		os.Exit(4)
	}
	var buff []string
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		if scan.Text() == "" { // remove blank lines

		} else {
			buff = append(buff, scan.Text())
		}
	}
	f.Close()
	return buff
}

func save(b []string) {
	path, _ := os.UserHomeDir()
	path += "/.todo.mine"
	f, err := os.OpenFile(path, os.O_WRONLY, 0600)
	f.Truncate(0)
	if err != nil {
		print(err)
		os.Exit(5)
	}
	for _, e := range b {
		f.WriteString(e + "\n")
	}
}

func main() {
	var todo []string = load() // load previous todo
	scan := bufio.NewScanner(os.Stdin)
	clear()
	for {
		fmt.Println("To do (h for help):")
		for n, do := range todo {
			fmt.Printf("  %d. %s\n", n+1, do)
		}
		fmt.Print("> ")
		scan.Scan()
		if _, err := strconv.Atoi(scan.Text()); err == nil {
			num, _ := strconv.Atoi(scan.Text())
			if num == -1 {
				save(todo)
				os.Exit(0)
			} else if num > len(todo) || num <= 0 {
				println("number not in the list...")
				time.Sleep(1 * time.Second)
			} else {
				todo = remove(todo, num-1)

			}
		} else {
			if scan.Text() == "h" {
				help()
			} else if scan.Text() == "" {

			} else {
				todo = append(todo, scan.Text())
			}
		}
		clear()
	}
}
