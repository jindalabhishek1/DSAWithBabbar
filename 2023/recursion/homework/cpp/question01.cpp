// factorial using recursion

#include <iostream>

using namespace std;

int factorial(int);

int main()
{
    int input, output = 0;
    cout << input << " " << output << "\n";
    cout << "Enter the number for which you want to find factorial: ";
    cin >> input;

    output = factorial(input);
    cout << "Factorial of " << input << " is: " << output << "\n";
}

int factorial (int num)
{
    if (num == 1)
    {
        return num;
    }

    return num * factorial(num - 1);
}