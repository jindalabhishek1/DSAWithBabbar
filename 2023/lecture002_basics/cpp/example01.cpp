// size of void

#include <iostream>

using namespace std;

int main()
{
    // you can't check size of void
    // cout << sizeof (void) << endl;

    cout << sizeof(__null) << endl;
}