#include <iostream>
using namespace std;

int main(){
    int x,y,z;
    int *p1;
    int *p2;
    int *p3;

    x = 11;
    y = 16;
    p2 = &x;
    p1 = p2;
    z = 28;
    p1 = &z;
    x = *p1;
    p3 = &y;
    *p2 = 9;
    

    cout << p1 << endl;

}