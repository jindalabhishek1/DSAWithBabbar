// program to convert negative decimal to binary

#include <iostream>

using namespace std;

int main()
{
    int decimal;
    cout << "Enter a negative decimal number: ";
    cin >> decimal;

    // checking if the decimal is negative
    if (decimal >= 0)
    {
        cout << "The decimal number is not negative.\n";
        return 1;
    }

    // converting decimal to binary
    int size = sizeof(decimal) * 8;
    int binary[size];

    int count = 0;

    // converting decimal to positive
    decimal *= -1;

    while (count < size)
    {
        int bit = decimal & 1;
        
        // doing ones complement
        // cout << "Bit: " << bit << "\n";
        bit = bit ^ 1;
        // cout << "~Bit: " << bit << "\n";
        binary[size - count - 1] = bit;
        decimal = decimal >> 1;
        count++;
    }

    // doing twos complement
    int i = 0;
    while (true)
    {
        if (binary[count - i - 1] == 0)
        {
            binary[count - i - 1] = 1;
            break;
        }
        else
        {
            binary[count - i - 1] = 0;
        }
        i++;
    }

    // printing the binary
    cout << "The binary representation is: ";
    for (int i = 0; i < size; i++)
    {
        cout << binary[i];
    }
    cout << "\n";
    return 0;
}