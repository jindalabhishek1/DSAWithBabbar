// fibbonacci series using for loop

#include <iostream>

using namespace std;

int main()
{
    int n;
    cout << "Enter the number of terms: ";
    cin >> n;

    // first two terms
    int a = 0, b = 1;

    for (int i = 0; i < n; i++)
    {
        // if user wants to print the first two terms
        if (i == 0)
        {
            cout << "in if " << endl;
            cout << a << " ";
        }
        // if user wants to print second term
        else if (i == 1)
        {
            cout << b << " ";
        }
        // calculating remaining terms
        else
        {
            int nextNumber = a + b;
            cout << nextNumber << " ";
            a = b;
            b = nextNumber;
        }
    }
    cout << endl;

    return 0;
}