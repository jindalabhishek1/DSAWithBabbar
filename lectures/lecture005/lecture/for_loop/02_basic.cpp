// for loop example for missing arguments

#include <iostream>

using namespace std;

int main()
{
    int n;
    cout << "Enter a number: ";
    cin >> n;

    //for loop
    // for (initialization(s); condition(s); increment(s))

    int i = 0;
    for (; ; )
    {
        if (i < n)
        {
            cout << i + 1 << " ";
        }
        else
        {
            // for breaking out of loop
            break;
        }
        i++;
    }
    cout << endl;
    return 0;
}