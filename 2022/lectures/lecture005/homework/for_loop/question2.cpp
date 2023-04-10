// second question on for loop, timestamp: 49:30

#include <iostream>

using namespace std;

int main()
{
    for (int i = 0; i <= 5; i--)
    {
        cout << i << " ";
        i++;
    }
    cout << endl;

    // output: infinite loop printing zeros

    return 0;
}