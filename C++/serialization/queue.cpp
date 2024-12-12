#include <iostream>
#include "../structures/include/queue.h"

using namespace std;

int main() {

    Queue<int> queue(10);
    queue.push(1);
    queue.push(2);
    queue.push(3);

    cout << "Текстовый формат: ";
    queue.serializeText("files/queue.txt");
    Queue<int> textQueue(10);
    textQueue.deserializeText("files/queue.txt");
    while (!textQueue.isEmpty()) {
        cout << textQueue.peek() << " ";
        textQueue.pop();
    }
    cout << endl;

    cout << "Бинарный формат: ";
    queue.serializeText("files/queue.bin");
    Queue<int> binQueue(10);
    binQueue.deserializeText("files/queue.bin");
    while (!binQueue.isEmpty()) {
        cout << binQueue.peek() << " ";
        binQueue.pop();
    }
    cout << endl;

    return 0;
}