package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// version var (doesn't work well for now)
var vers string

// help show how to use todo
func help() {
	clear()
	fmt.Println(`Usage :
	write something to add in the list
	write a number in the list to delete it
	write s to save the todo
	write all to clear the todo list
	write -1 to save and end the program`)
	fmt.Printf("\nVersion : %s\n", vers)
	fmt.Print("Press Enter to continue...")
	fmt.Scanln()
}

// remove an element in an array
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

// clear the screen
func clear() {
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

// load the todo with the file
func load() []string {
	path, _ := os.UserHomeDir() // get user home folder
	path += "/.todo.mine"
	f, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	// decode the file
	data, err := base64.StdEncoding.DecodeString(string(f))
	if err != nil {
		return nil
	}
	return strings.Split(string(data), "\n")
}

// save to the todo file in base64
func save(todo []string) {
	path, _ := os.UserHomeDir() // get user home folder
	path += "/.todo.mine"
	data := base64.StdEncoding.EncodeToString([]byte(strings.Join(todo, "\n")))
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString(data)
	f.Close()
}

// show the todo list
func show_list(todo []string) {
	for n, do := range todo {
		fmt.Printf("  %d. %s\n", n+1, do)
	}
}

// add/remove/error if user entered a number
func action_number(todo *[]string, scan *bufio.Scanner) {
	num, _ := strconv.Atoi(scan.Text())
	if num == -1 {
		save(*todo)
		os.Exit(0)
	} else if num > len(*todo) || num <= 0 {
		println("number not in the list...")
		time.Sleep(1 * time.Second)
	} else {
		*todo = remove(*todo, num-1)

	}
}

// save/help/remove all if the user entered a string
func action_str(todo *[]string, scan *bufio.Scanner) {
	if scan.Text() == "h" {
		help()
	} else if scan.Text() == "s" {
		save(*todo)
	} else if scan.Text() == "all" {
		*todo = nil
		clear()
	} else if scan.Text() == "" {

	} else {
		*todo = append(*todo, scan.Text())
	}
}

func main() {
	var todo []string = load()
	scan := bufio.NewScanner(os.Stdin)
	clear()
	for {
		fmt.Println("To do (h for help):")
		show_list(todo)
		fmt.Print("> ")
		scan.Scan()
		if _, err := strconv.Atoi(scan.Text()); err == nil { // if number
			action_number(&todo, scan)
		} else { // if string
			action_str(&todo, scan)
		}
		clear()
	}
}
