package main

import (
    "fmt"
    "linkedlist"
)

func main() {
    list := dlb.CreateLinkedList()
    s := []int{1, 2, 3, 4, 5, 6, 7}
    for i, v := range s {
        list.Insert(i+1, v)
    }
    list.Print()
    ok := list.Insert(2, 99)
    if !ok {
        fmt.Println("insert fail")
    }
    list.Print()
    fmt.Println(list.Len())

}
