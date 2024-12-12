#include <iostream>
#include "../structures/include/array.h"

using namespace std;

int main() {
    
    Array<int> arr;
    arr.add(15);
    arr.add(23);
    arr.add(73);
    arr.add(82);
    arr.add(13);
    arr.add(400);

    cout << "Текстовый формат: ";
    arr.serializeText("files/array.txt");
    Array<int> textArr;
    textArr.deserializeText("files/array.txt");
    textArr.print();

    cout << "Бинарный формат: ";
    arr.serializeBinary("files/array.bin");
    Array<int> binArr;
    binArr.deserializeBinary("files/array.bin");
    binArr.print();

    return 0;
}