#include <iostream>

using namespace std;

int main()
{
    int i = 3;
    // you can skip the condition from for loop, but you have to use break in loop to come out
    // this is example of infinite loop
    for (;;)
    {
        if (i > 0) 
        {
            cout << i << endl;
            i--;
        }
        
    }
}