#include <iostream>

using namespace std;

int main(void)
{
    int n;
    cout << "Enter the number of rows: ";
    cin >> n;

    // loop for the number of rows
    int i = 1;
    while (i <= n)
    {
        int j = 1;

        // loop for first triangle
        while (j <= n - i + 1)
        {
            cout << j << " ";
            j++;
        }

        // loop for second triangle
        j = 1;
        while (j < i)
        {
            cout << "* ";
            j++;
        }

        // loop for third triangle
        j = 1;
        while (j < i)
        {
            cout << "* ";
            j++;
        }

        // loop for fourth triangle
        j = 1;
        while (j <= n - i + 1)
        {
            cout << n - i - j + 2 << " ";
            j++;
        }
        cout << endl;
        i++;
    }
}