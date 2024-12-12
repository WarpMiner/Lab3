#include <iostream>
#include "../structures/include/stack.h"

using namespace std;

int main() {

    Stack<int> stack(10);
    stack.push(1);
    stack.push(2);
    stack.push(3);

    cout << "Текстовый формат: ";
    stack.serializeText("files/stack.txt");
    Stack<int> textStack(10);
    textStack.deserializeText("files/stack.txt");
    while (!textStack.isEmpty()) {
        cout << textStack.peek() << " ";
        textStack.pop();
    }
    cout << endl;

    cout << "Бинарный формат: ";
    stack.serializeBinary("files/stack.bin");
    Stack<int> binStack(10);
    binStack.deserializeBinary("files/stack.bin");
    while (!binStack.isEmpty()) {
        cout << binStack.peek() << " ";
        binStack.pop();
    }
    cout << endl;

    return 0;
}