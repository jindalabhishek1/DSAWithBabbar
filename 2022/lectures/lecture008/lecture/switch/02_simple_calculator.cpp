// simple calculator using switch
#include <iostream>

using namespace std;

int main ()
{
    int num1, num2;

    cout << "Enter first number: ";
    cin >> num1;

    cout << "Enter second number: ";
    cin >> num2;

    char operation;
    cout << "Enter the operation to perform: ";
    cin >> operation;

    switch (operation)
    {
        case '+':
            cout << "Sum: " << num1 + num2 << endl;
            break;
        case '-': 
            cout << "Difference: " << num2 - num1 << endl;
            break;
        case '*':
            cout << "Product: " << num1 * num2 << endl;
            break;
        case '/':
            cout << "Division: " << num1 / num2 << endl;
            break;
        default:
            cout << "Default case!" << endl;
    }

    return 0;
}