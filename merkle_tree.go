package main

import (
  "fmt"
  "bufio"
  "os"
  "mrktree"
)

func main() {
  reader := bufio.NewReader(os.Stdin)

  fmt.Print("Source data: ")
  sourcedata, _ := reader.ReadSlice('\n')
  sourcedata = sourcedata[:len(sourcedata)-1]
  sourceMerkelTree := mrktree.New(sourcedata)

  fmt.Print("Target data: ")
  targetData, _ := reader.ReadSlice('\n')
  targetData = targetData[:len(targetData)-1]
  targetMerkelTree := mrktree.New(targetData)

  if mrktree.Match(sourceMerkelTree, targetMerkelTree) {
    fmt.Println("Data is the same")
  } else {
    fmt.Println("Data is different")
  }
}