// first MCQ question in the lecture timeframe: 24:16

#include <iostream>

using namespace std;

int main()
{
    int a, b = 1;
    a = 10;
    if (++a)
        cout << b << endl;
    else
        cout << ++b << endl;

    // output: 1
    // cout << "a: " << a << " b: " << b << endl;
    return 0;
}