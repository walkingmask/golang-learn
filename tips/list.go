/********************
 * list.go
 * 2016/11/11(金)
 * walkingmask
 * Stack/Queue using container/list sample
 ********************/

package main

import (
  "fmt"
  "container/list"
)

type block struct {
  key int
  value int
}

func main() {
  // Stack
  fmt.Println("Stack:")
  stack := list.New()
  for i := 0; i < 10; i++ {
    stack.PushFront(block{i, i*(i+1)})
  }
  for i := 0; i < 10; i++ {
    v := stack.Remove(stack.Front())
    fmt.Print(v, " ")
  }
  // Queue
  fmt.Println("\nQueue:")
  queue := list.New()
  for i := 0; i < 10; i++ {
    queue.PushBack(block{i, i*(i+1)})
  }
  for i := 0; i < 10; i++ {
    v := queue.Remove(queue.Front())
    fmt.Print(v, " ")
  }
}
