// switch case

#include <iostream>

using namespace std;

int main()
{
    int num = 10;
    int temp = 5;
    switch (num-5){
        case 3:
            cout << "first case" << endl;
            break;
        case 5:
            cout << "second case" << endl;
            break;
            // we can only use continue in loop
        // can't add duplicate case in cpp
        // case 5:
        //     cout << "duplicate second case" << endl;
        //     break;
        default:
            cout << "Default case" << endl;
    }
}