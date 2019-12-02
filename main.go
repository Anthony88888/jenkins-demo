package main

// Import the fmt for formatting strings
// Import os so we can read environment variables from the system
import (
	"fmt"
	"os"
)

func main() {
<<<<<<< HEAD
<<<<<<< HEAD
	fmt.Println("Hello, Kubernete！I'm from Jenkins CI！")
=======
	fmt.Println("Hello, Kubernete！I'm from Jenkins CI！")
<<<<<<< HEAD
>>>>>>> dev
	fmt.Println("BRANCH:", os.Getenv("branch"))
=======
	fmt.Println("Hello, Kubernete！I'm from Jenkins CI！")
	fmt.Println("BRANCH:", os.Getenv("branch"))
>>>>>>> dev
=======
	fmt.Println("BRANCH:", os.Getenv("branch"))
>>>>>>> dev
}