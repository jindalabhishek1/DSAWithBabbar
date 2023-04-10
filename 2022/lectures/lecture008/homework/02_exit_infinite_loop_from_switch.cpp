#include <iostream>
using namespace std;

int main()
{
    int num = 2;

    cout << "Start" << endl;
    while (1)
    {
        cout << "Inside while loop :)" << endl;
        switch (num)
        {
            cout << "Inside switch" << endl;
            case 2: 
                cout << "Inside case 2" << endl;
                // exit(<exitCode>) exits the program with the exit code provided
                exit(1);
                break;
            cout << "Ouside case" << endl;
        }
        cout << "Outside Switch" << endl;
    }
    cout << "Below while loop" << endl;
}