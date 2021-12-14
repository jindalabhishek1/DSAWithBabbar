// prime number or not using for

#include <iostream>

using namespace std;

int main()
{
    int n;
    cout << "Enter a number: ";
    cin >> n;

    bool isPrime = true;

    for (int i = 2; i < n; i++)
    {
        // if n is divisible by i, then it is not prime
        if (n % i == 0)
        {
            isPrime = false;

            // no need to iterate further
            break;
        }
    }

    if (isPrime)
    {
        cout << n << " is a prime number." << endl;
    }
    else
    {
        cout << n << " is not a prime number." << endl;
    }
    return 0;
}