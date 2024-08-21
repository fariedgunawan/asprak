#include <iostream>

using namespace std;

typedef char infotype;

struct Stack{
  infotype info[15];
  int top;
};

void createStack(Stack &S){
    S.top = 0;
}

bool isFull(Stack S){
    return S.top == 15;
}

bool isEmpty(Stack S){
    return S.top == 0;
}

infotype popStack(Stack &S){
    infotype x;
    x = S.info[S.top];
    S.top--;
    return x;
}

void pushStack(Stack &S,infotype x){
    if(!isFull(S)){
        S.top++;
        S.info[S.top] = x;
    }
}

bool isMatching(char x, char y){
    if(x == '(' && y == ')') return true;
    if(x == '[' && y == ']') return true;
    if(x == '{' && y == '}') return true;
    return false;
}

bool areBalanced(string x){
    int i;
    Stack S;
    createStack(S);
    for (i= 0; i < x.length(); i++){
        if(x[i]== '(' || x[i]== '{' || x[i]== '['){
            pushStack(S,x[i]);
        }else if(x[i] == ')' || x[i] == '}' || x[i] == ']'){
            if(isEmpty(S)){
                return false;
            }else if(!isMatching(popStack(S),x[i])){
                return false;
            }
        }
    }
    return isEmpty(S);
}

int main()
{
    string kata;
    cout << "Hello Faried" << endl;

    cout << "Masukan kata: ";
    cin >> kata;

    if(areBalanced(kata)){
        cout << "Balanced";
    }else{
        cout << "Not Balanced";
    }

    return 0;
}
