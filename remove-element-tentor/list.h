#ifndef LIST_H_INCLUDED
#define LIST_H_INCLUDED

#include <iostream>
using namespace std;

struct mahasiswa{
    string nama;
    int nim;
    string prodi;
};

typedef struct mahasiswa infoMhs;
typedef struct elm * adrMhs;

struct elm{
    adrMhs next;
    infoMhs infotypemhs;
};

struct List{
    adrMhs first;
};

adrMhs newElm(infoMhs x);
void insertLast(List &L,adrMhs p);
void createList(List &L);
void deleteLast(List &L,adrMhs &p);
void deleteFirst(List &L,adrMhs &p);
void deleteAfter(adrMhs prec,adrMhs &p);
adrMhs findElement(List L, infoMhs x);
void RemoveElm(List &L, infoMhs x);

void showList(List L);




#endif // LIST_H_INCLUDED
