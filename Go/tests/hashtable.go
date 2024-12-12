package hashtable

import (
	"hash/fnv"
)

const TABLE_SIZE = 100 // Размер таблицы

type KeyValuePair struct {
	key   string
	value interface{}
	next  *KeyValuePair
}

type HashTable struct {
	table     [TABLE_SIZE]*KeyValuePair
	sizetable int
}

func NewHashTable() *HashTable {
	return &HashTable{}
}

func (ht *HashTable) hashFunction(key string) int { // хеш-функция
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() % uint32(TABLE_SIZE))
}

func (ht *HashTable) Insert(key string, value interface{}) { // вставка

	index := ht.hashFunction(key)
	newPair := &KeyValuePair{key: key, value: value}

	if ht.table[index] == nil {
		ht.table[index] = newPair
		ht.sizetable++
	} else {
		current := ht.table[index]
		for current != nil {
			if current.key == key {
				current.value = value // Обновляем значение, если ключ существует
				return
			}
			if current.next == nil { // Если достигли конца цепочки
				current.next = newPair // Добавляем новый элемент в конец цепочки
				ht.sizetable++
				return
			}
			current = current.next
		}
	}
}

func (ht *HashTable) Get(key string) (interface{}, bool) { // возвращение элемента по ключу

	index := ht.hashFunction(key)
	current := ht.table[index]
	for current != nil {
		if current.key == key {
			return current.value, true // Возвращаем значение
		}
		current = current.next
	}
	return nil, false // Ключ не найден
}

func (ht *HashTable) Remove(key string) bool { // удаление элемента по ключу

	index := ht.hashFunction(key)
	current := ht.table[index]
	var previous *KeyValuePair

	for current != nil {
		if current.key == key {
			if previous != nil {
				previous.next = current.next // Удаляем элемент из цепочки
			} else {
				ht.table[index] = current.next // Если это первый элемент в цепочке
			}
			ht.sizetable--
			return true // Успешное удаление
		}
		previous = current
		current = current.next
	}
	return false // Ключ не найден
}