#include <iostream>

using namespace std;

void tuker(string &stac,int i, int j){
    char temp  = stac[i];
    stac[i] = stac[j];
    stac[j] = temp;
}

void permutasi(string stac, int m, int n){
    if(m == n){
        cout << stac << endl;
    }
    else{
        for(int i = 1; i <= n;i++){
            tuker(stac,m,i);
            permutasi(stac,m+1,n);
            tuker(stac,m,i);
        }
    }

}
int main()
{
    string input = "1234";
    permutasi(input,0,input.length() - 1);
    return 0;
}
