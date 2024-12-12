package main

import (
	"fmt"
	hashtable "serialization/mypackage"
)

func main() {
	ht := hashtable.NewHashTable()
	ht.Insert("key1", "value1")
	ht.Insert("key2", "value2")

	fmt.Print("Текстовый формат: ")
	err := ht.SerializeText("hashtable.txt")
	if err != nil {
		fmt.Println("Ошибка при сериализации в текст:", err)
		return
	}
	htnew := hashtable.NewHashTable()
	err = htnew.DeserializeText("hashtable.txt")
	if err != nil {
		fmt.Println("Ошибка при десериализации из текста:", err)
		return
	}
	value, found := htnew.Get("key1")
	if found {
		fmt.Printf("key1:%s ", value)
	}
	value, found = htnew.Get("key2")
	if found {
		fmt.Printf("key2:%s\n", value)
	}

	fmt.Print("Бинарный формат: ")
	err = ht.SerializeBinary("hashtable.bin")
	if err != nil {
		fmt.Println("Ошибка при сериализации в бинарный формат:", err)
		return
	}
	htanother := hashtable.NewHashTable()
	err = htanother.DeserializeBinary("hashtable.bin")
	if err != nil {
		fmt.Println("Ошибка при десериализации из бинарного формата:", err)
		return
	}
	value, found = htanother.Get("key1")
	if found {
		fmt.Printf("key1:%s ", value)
	}
	value, found = htanother.Get("key2")
	if found {
		fmt.Printf("key2:%s\n", value)
	}
}