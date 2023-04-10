// sum of n numbers using for loop

#include <iostream>

using namespace std;

int main()
{
    int n, sum = 0;
    cout << "Enter a number: ";
    cin >> n;

    //for loop
    for (int i = 0; i < n; i++)
    {
        // i+1 because, i = 0
        sum += (i + 1);
    }

    cout << "Sum: " << sum << endl;
    return 0;
}