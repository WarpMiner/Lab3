package hashtable

import (
	"encoding/binary"
	"fmt"
	"hash/fnv"
	"os"
	"strings"
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

func (ht *HashTable) SerializeText(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, pair := range ht.table {
		current := pair
		for current != nil {
			line := fmt.Sprintf("%s:%v\n", current.key, current.value)
			if _, err := file.WriteString(line); err != nil {
				return err
			}
			current = current.next
		}
	}

	return nil
}

func (ht *HashTable) DeserializeText(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n") // содержимое файла приобразуется в срез строку
	for _, line := range lines {               // обработка каждой строки
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ":", 2) // разделение строки на ключ значение
		if len(parts) != 2 {
			return fmt.Errorf("неправильный формат строки: %s", line)
		}

		key := parts[0]
		value := parts[1]

		ht.Insert(key, value)
	}

	return nil
}

func (ht *HashTable) SerializeBinary(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, head := range ht.table {
		current := head
		for current != nil {
			keyLen := uint32(len(current.key))
			valueLen := uint32(len(fmt.Sprintf("%v", current.value)))

			if err := binary.Write(file, binary.LittleEndian, keyLen); err != nil { // длина ключа
				return err
			}
			if _, err := file.Write([]byte(current.key)); err != nil { // ключ
				return err
			}
			if err := binary.Write(file, binary.LittleEndian, valueLen); err != nil { // длина значения
				return err
			}
			if _, err := file.Write([]byte(fmt.Sprintf("%v", current.value))); err != nil { // значения
				return err
			}

			current = current.next
		}
	}

	return nil
}

func (ht *HashTable) DeserializeBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for {
		var keyLen uint32
		if err := binary.Read(file, binary.LittleEndian, &keyLen); err != nil { // длина ключа
			break // Достигнут конец файла или произошла ошибка
		}

		keyBytes := make([]byte, keyLen)
		if _, err := file.Read(keyBytes); err != nil { // ключ
			return err
		}
		key := string(keyBytes)

		var valueLen uint32
		if err := binary.Read(file, binary.LittleEndian, &valueLen); err != nil { // длина значения
			return err
		}

		valueBytes := make([]byte, valueLen)
		if _, err := file.Read(valueBytes); err != nil { // значения
			return err
		}
		value := string(valueBytes)

		// Вставляем данные в хеш-таблицу
		ht.Insert(key, value)
	}

	return nil
}