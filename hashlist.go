package main

import (
  "fmt"
  "bufio"
  "os"
  h "hashlist"
)

func main() {
  reader := bufio.NewReader(os.Stdin)

  fmt.Print("Source data: ")
  sourcedata, _ := reader.ReadSlice('\n')
  sourceHashlist := h.New(sourcedata)

  fmt.Print("Target data: ")
  targetData, _ := reader.ReadSlice('\n')
  targetHashlist := h.New(targetData)

  if sourceHashlist.Checksum == targetHashlist.Checksum {
    fmt.Println("Data is the same")
  } else {
    fmt.Println("Data is different")
  }
}