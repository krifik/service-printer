// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"os/exec"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	// Replace "localhost:631" with the address of your CUPS server
// 	cmd := exec.Command("lpstat", "-a", "-h", "localhost:631")
// 	out, err := cmd.Output()
// 	if err != nil {
// 		panic(err)
// 	}

// 	printerNames := []string{}
// 	splitStr := strings.Split(string(out), "\n")
// 	for _, item := range splitStr {
// 		fields := strings.Fields(item)
// 		if len(fields) > 0 {
// 			printerNames = append(printerNames, fields[0])
// 		}
// 	}

// 	for index, name := range printerNames {
// 		fmt.Println(index+1, ":", name)
// 	}
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	pick := scanner.Text()
// 	pickInt, err := strconv.Atoi(pick)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defPrinter := printerNames[pickInt-1]
// 	fmt.Println("Using printer:", defPrinter)
// 	setDef := exec.Command("lpoptions", "-d", defPrinter)
// 	setDefOut, err := setDef.CombinedOutput()
// 	if err != nil {
// 		fmt.Println(err)
// 		fmt.Println(setDefOut)
// 	}
// 	def := exec.Command("lpstat", "-d")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defOut, err := def.CombinedOutput()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(defOut))
// 	fmt.Println("Print Now?")
// 	scanner.Scan()
// 	ans := scanner.Text()
// 	if ans == "y" {
// 		filePath := "/home/fikri/Downloads/asdasdasd.pdf" // Replace with the path of the file you want to print

// 		printCmd := exec.Command("lp", filePath)
// 		printCmd.Stdout = os.Stdout
// 		printCmd.Stderr = os.Stderr

// 		err := printCmd.Run()
// 		if err != nil {
// 			fmt.Println("Error printing file:", err)
// 			return
// 		}

// 		fmt.Println("File printed successfully.")
// 		cmd := exec.Command("lpstat", "-o", "-l")
// 		out, err := cmd.CombinedOutput()
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(string(out))

// 		printJobs := strings.Split(string(out), "\n\n")

// 		status := "completed" // Replace with the desired status ("completed", "pending", "held", "processing", etc.)

// 		fmt.Printf("Print jobs with status '%s':\n", status)
// 		for _, job := range printJobs {
// 			if strings.Contains(job, "Status:") && strings.Contains(job, status) {
// 				fmt.Println(job)
// 				fmt.Println("-----------------")
// 			}
// 		}

// 	}
// }
