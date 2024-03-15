package main

import "fmt"

func main()  {
	buah := []string{"Apel", "Jeruk", "Anggur", "Melon"}

	buah = append(buah, "Pepaya")

	for i:=0; i<len(buah); i++ {
		fmt.Println("Element ke - :", i, buah[i])
	}

}