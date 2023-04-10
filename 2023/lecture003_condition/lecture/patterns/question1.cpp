/* Pattern question

*****
*****
*****

*/

#include <iostream>

using namespace std;

int main ()
{
    // will have random values
    int row, col;
    // cout << "Row: " << row << "\nCol: " << col << endl;

    cout << "Row: ";
    cin >> row;
    cout << "Col: ";
    cin >> col;

    for (int i = 0; i < row; i++)
    {
        for (int j = 0; j < col; j++)
        {
            cout << "*";
        }
        cout << "\n";
    }
}