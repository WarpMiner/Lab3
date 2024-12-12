#include <iostream>
#include "../structures/include/hashtable.h"

using namespace std;

int main() {

    HashTable<int> table;
    table.insert("key1", 1);
    table.insert("key2", 2);
    table.insert("key3", 3);
    table.insert("key4", 4);
    table.insert("key5", 5);
    table.insert("key6", 6);

    cout << "Текстовый формат: ";
    table.serializeText("files/hashtable.txt");
    HashTable<int> textTable;
    textTable.deserializeText("files/hashtable.txt");
    int value;
    textTable.get("key1", value);
    cout << "key1:" << value << " ";
    textTable.get("key2", value);
    cout << "key2:" << value << " ";
    textTable.get("key3", value);
    cout << "key3:" << value << " ";
    textTable.get("key4", value);
    cout << "key4:" << value << " ";
    textTable.get("key5", value);
    cout << "key5:" << value << " ";
    textTable.get("key6", value);
    cout << "key6:" << value;
    cout << endl;

    cout << "Бинарный формат: ";
    table.serializeBinary("files/hashtable.bin");
    HashTable<int> binTable;
    binTable.deserializeBinary("files/hashtable.bin");
    binTable.get("key1", value);
    cout << "key1:" << value << " ";
    binTable.get("key2", value);
    cout << "key2:" << value << " ";
    binTable.get("key3", value);
    cout << "key3:" << value << " ";
    binTable.get("key4", value);
    cout << "key4:" << value << " ";
    binTable.get("key5", value);
    cout << "key5:" << value << " ";
    binTable.get("key6", value);
    cout << "key6:" << value;
    cout << endl;

    return 0;
}