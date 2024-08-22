#include "list.h"

int main()
{
    List L;
    adrMhs p;
    infoMhs x;
    cout << "Hello world!" << endl;

    createList(L);
    x.nama = "Muhammad Faried Gunawan";
    x.nim = 1301223229;
    x.prodi = "FIF";
    insertLast(L,newElm(x));
    x.nama = "Ghani Naja";
    x.nim = 1301223119;
    x.prodi = "FIF";
    insertLast(L,newElm(x));
    x.nama = "Daffa Prima Putra";
    x.nim = 1301223339;
    x.prodi = "FIF";
    insertLast(L,newElm(x));
    showList(L);
    cout << "Delete" << endl;
    x.nama = "Ghani Naja";
    x.nim = 1301223119;
    x.prodi = "FIF";
    RemoveElm(L,x);
    showList(L);
    return 0;
}
