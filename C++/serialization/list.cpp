#include <iostream>
#include "../structures/include/list.h"

using namespace std;

int main() {

    cout << "Односвязный список: " << endl;
    SinglyLinkedList<int> listSL;
    listSL.push_back(1);
    listSL.push_back(2);
    listSL.push_back(3);

    cout << "Текстовый формат: ";
    listSL.serializeText("files/list.txt");
    SinglyLinkedList<int> textListSL;
    textListSL.deserializeText("files/list.txt");
    textListSL.print();

    cout << "Бинарный формат: ";
    listSL.serializeBinary("files/list.bin");
    SinglyLinkedList<int> binListSL;
    binListSL.deserializeBinary("files/list.bin");
    binListSL.print();


    cout << endl << "Двухсвязный список: " << endl;
    DoublyLinkedList<int> listDL;
    listDL.push_back(3);
    listDL.push_back(2);
    listDL.push_back(1);

    cout << "Текстовый формат: ";
    listDL.serializeText("files/list.txt");
    SinglyLinkedList<int> textListDL;
    textListDL.deserializeText("files/list.txt");
    textListDL.print();

    cout << "Бинарный формат: ";
    listDL.serializeBinary("files/list.bin");
    SinglyLinkedList<int> binListDL;
    binListDL.deserializeBinary("files/list.bin");
    binListDL.print();

    return 0;
}