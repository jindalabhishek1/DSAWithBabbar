// demonstrate continue

#include <iostream>

using namespace std;

int main()
{
    int n;
    cout << "Enter a number: ";
    cin >> n;

    //for loop
    for (int i = 0; i < n; i++)
    {
        // skipping even numbers
        if (i % 2 == 0)
        {
            continue;
        }
        cout << i << " ";
    }
    cout << endl;
    return 0;
}