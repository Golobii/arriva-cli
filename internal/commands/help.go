package commands

import "fmt"

func PrintHelp() {
	fmt.Println("Commands:")
	fmt.Println("./arriva-cli help   # Prints the help section (WIP)")
	fmt.Println("./arriva-cli date   # Prints the current date")
	fmt.Println("./arriva-cli list   # Lists all stations")
	fmt.Println("./arriva-cli \"[start station]\" \"[end station]\" [date] # Displays busses")
	fmt.Println()
	fmt.Println("GitHub: https://github.com/Golobii/arriva-cli")
}
