// Program to convert binary to decimal

#include <iostream>
#include <math.h>

using namespace std;

int main()
{
    // we are considering the binary number to be in the form of an int
    int binary, decimal = 0;
    cout << "Enter a binary number: ";
    cin >> binary;

    // converting the binary number to decimal
    int i = 0;
    while (binary > 0)
    {
        int remainder = binary % 10;

        // checking if input is a binary number, i.e. 0 or 1
        if (remainder != 1 && remainder != 0)
        {
            cout << "Invalid binary number" << endl;
            return 1;
        }

        /*
            Case 1: if remainder = 1: 1 * 2^i, i.e. this will contribute to the decimal number
            Case 2: if remainder = 0: 0 * 2^i, i.e. this will not contribute to the decimal number
        */
        decimal += remainder * pow(2, i);
        binary /= 10;
        i++;
    }

    cout << "The decimal number is: " << decimal << "\n";
    return 0;
}