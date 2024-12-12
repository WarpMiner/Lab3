#!/bin/bash

# Компилируем тесты с поддержкой покрытия
g++ -g -O0 --coverage -o test_array test_array.cpp
g++ -g -O0 --coverage -o test_list test_list.cpp
g++ -g -O0 --coverage -o test_stack test_stack.cpp
g++ -g -O0 --coverage -o test_queue test_queue.cpp
g++ -g -O0 --coverage -o test_hashtable test_hashtable.cpp
g++ -g -O0 --coverage -o test_tree test_tree.cpp

# Запускаем тесты
./test_array
./test_list
./test_stack
./test_queue
./test_hashtable
./test_tree


# Генерируем отчет о покрытии
lcov --capture --directory . --output-file coverage.info --rc geninfo_unexecuted_blocks=1 --ignore-errors inconsistent

# Генерируем HTML-отчет
genhtml coverage.info --output-directory coverage_html

# Удаляем лишнее
rm coverage.info
rm test_array test_array.gcda test_array.gcno
rm test_list test_list.gcda test_list.gcno
rm test_stack test_stack.gcda test_stack.gcno
rm test_queue test_queue.gcda test_queue.gcno
rm test_hashtable test_hashtable.gcda test_hashtable.gcno
rm test_tree test_tree.gcda test_tree.gcno tree.txt tree.bin

# Открываем отчет
xdg-open coverage_html/index.html