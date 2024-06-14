package main

import (
  "fmt"
  "os"
)

//this is a comment

func main() {
  fmt.Println("Hello, my name is Himani")
  fmt.Println("This is a message before exiting.")

  // Exit the program with a status code

  os.Exit(1)

  /* This line will not be executed because os.Exit terminates the program */
    fmt.Println("This message will not be printed.")
}
