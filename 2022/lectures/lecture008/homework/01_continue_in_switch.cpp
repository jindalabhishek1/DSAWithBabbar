// try continue in switch

#include <iostream>

using namespace std;

int main()
{
    int num = 2;

    // switch (<expersion>)
    switch (num)
    {
        // case <int or char>:
        case 1:
            cout << "One" << endl;
            // break is used to break out of the switch block
            break;
        case 2:
            cout << "Two" << endl;
            /*** 
             * continue is only used for loops
             * Continue is used to jump to next iteration immediately, in switch we don't have iteration
             ***/
            conitnue;
            // If we don't use break, all the cases below this will be executed
        
        // default will be executed if no case matches
        default:
            cout << "default" << endl;
    }
}