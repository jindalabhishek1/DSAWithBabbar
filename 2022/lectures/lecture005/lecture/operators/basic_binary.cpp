// program for basic binary operators

#include <iostream>

using namespace std;

int main()
{
    int number1 = 4;
    int number2 = 6;

    // And operato(&): returns true if both operands are true
    cout << "number1 & number2: " << (number1 & number2) << '\n';

    // Or operato(|): returns true if either operand is true
    cout << "number1 | number2: " << (number1 | number2) << '\n';

    // Xor operato(^): returns true if operands are different
    cout << "number1 ^ number2: " << (number1 ^ number2) << '\n';

    // Not operato(~): returns the negation of the operand (invert the bits)
    cout << "~number1: " << ~number1 << '\n';

    /*
        Left shift operato(<<): returns the left operand shifted by the number of bits specified by the right operand
        BE CAREFUL: 
            - the right operand must be a non-negative integer
            - padding will be done by 0 for positive numbers and compiler dependent for negative numbers
            - large number, for example, 01000000000000000110
    */
    cout << "number1 << 2: " << (number1 << 2) << '\n';
    cout << "number1 << -1: " << (number1 << -1) << '\n';

    // Right shift operato(>>): returns the left operand shifted by the number of bits specified by the right operand
    cout << "number1 >> 2: " << (number1 >> 2) << '\n';
    cout << "number1 >> -1: " << (number1 >> -1) << '\n';

    /*
        postfix increment operato(++): returns the operand before increment
        prefix increment operato(++): returns the operand after increment
        postfix decrement operato(--): returns the operand before decrement
        prefix decrement operato(--): returns the operand after decrement
    */
    int i = 7;
    cout << "i++: " << i++ << '\n'; // Post increment
    // print 7, i = 8
    cout << "++i: " << ++i << '\n'; // pre increment
    // print 9, i = 9
    cout << "i--: " << i-- << '\n'; // Post decrement
    // print 9, i = 8
    cout << "--i: " << --i << '\n'; // Pre decrement
    // print 7, i = 7




    return 0;
}