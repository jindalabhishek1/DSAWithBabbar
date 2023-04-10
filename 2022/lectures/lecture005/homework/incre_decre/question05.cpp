// fifth MCQ question in lecture, timeframe: 24:24

#include <iostream>

using namespace std;

int main()
{
    int a = 1;
    int b = a++;
    int c = ++a;
    cout << b << " ";
    cout << c << endl;

    // output: 1 3
}