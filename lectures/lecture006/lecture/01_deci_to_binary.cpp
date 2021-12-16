// program to convert decimal to binary

#include <iostream>
#include <math.h>

using namespace std;

int main()
{
    int decimal = 0;

    cout << "Enter a decimal number: ";
    cin >> decimal;
    // This solution will not work proper for large numbers and negative numbers
    // int count = 0;
    // long long binary = 0;
    // while (decimal > 0)
    // {
    //     int bit = decimal & 1;
    //     cout << "Bit: " << bit << endl;
    //     binary = (bit * pow(10, count)) + binary;
    //     cout << "Binary: " << binary << endl;
    //     decimal = decimal >> 1;
    //     cout << "Decimal: " << decimal << endl;
    //     count++;
    // }

    int size = sizeof(int) * 8;
    int binary[size];
    int count = 0;

    // checking if decimal is negative
    if (decimal < 0)
    {
        cout << "Decimal is negative" << endl;
        return 1;
    }
    
    while (decimal > 0)
    {
        int bit = decimal & 1;
        binary [size - count - 1] = bit;
        decimal = decimal >> 1;
        count++;
    }

    // printing binary
    // binary[0 to (size-count) will be garbage]
    for (int i = size - count; i < size; i++)
    {
        cout << binary[i];
    }
    cout << endl;
    // cout << "The binary equivalent is: " << binary << endl;
    // cout << "The number of bits is: " << count << endl;
    // cout << "Size of long long is: " << sizeof(long long) << endl;
    return 0;
}