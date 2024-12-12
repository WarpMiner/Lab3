#include <iostream>
#include "../structures/include/tree.h"

using namespace std;

int main() {

    BinaryTree tree;
    tree.insert(9);
    tree.insert(6);
    tree.insert(17);
    tree.insert(8);
    tree.insert(3);
    tree.insert(20);

    cout << "Текстовый формат:" << endl;
    tree.serializeText("files/tree.txt");
    BinaryTree textTree;
    textTree.deserializeText("files/tree.txt");
    textTree.print();
    
    cout << "Бинарный формат:" << endl;
    tree.serializeBinary("files/tree.bin");
    BinaryTree binTree;
    binTree.deserializeBinary("files/tree.bin");
    binTree.print();

    return 0;
}