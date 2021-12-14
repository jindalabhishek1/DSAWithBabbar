// for loop for printing numbers from 1 to n

#include <iostream>

using namespace std;

int main()
{
    int n;
    cout << "Enter a number: ";
    cin >> n;

    //for loop
    // for (initialization(s); condition(s); increment(s))
    for (int i = 0; i < n; i++)
    {
        cout << i + 1 << " ";
    }
    cout << endl;
    return 0;
}