#include "list.h"

adrMhs newElm(infoMhs x){
    adrMhs p;
    p = new elm;
    p->infotypemhs = x;
    p->next = nullptr;
    return p;
}
void insertLast(List &L,adrMhs p){
    adrMhs q;
    if(L.first == nullptr){
        L.first = p;
    }else{
        q = L.first;
        while(q->next != nullptr){
            q = q->next;
        }
        q->next = p;
    }
}
void createList(List &L){
    L.first = nullptr;
}
void deleteLast(List &L,adrMhs &p){
    adrMhs q;
    if(L.first == nullptr){
        cout << "List Kosong";
    }else if(L.first->next == nullptr){
        p = L.first;
        L.first = nullptr;
    }else{
        p = L.first;
        while(p->next != nullptr){
            q = p;
            p = p->next;
        }
        q->next = nullptr;
    }
}
void deleteFirst(List &L,adrMhs &p){
    if(L.first == nullptr){
        cout << "List Kosong";
    }else if(L.first->next == nullptr){
        p = L.first;
        L.first = nullptr;
    }else{
        p = L.first;
        L.first = p->next;
        p->next = nullptr;
    }
}
void deleteAfter(adrMhs prec,adrMhs &p){
    if(prec == nullptr){
        cout << "Prec Kosong";
    }else{
        p = prec->next;
        prec->next = p->next;
        p->next = nullptr;
    }
}

adrMhs findElement(List L, infoMhs x) {
    adrMhs p = L.first;
    while (p != nullptr) {
        if (p->infotypemhs.nama == x.nama && p->infotypemhs.nim == x.nim) {
            return p;
        }
        p = p->next;
    }
    return nullptr;
}

void RemoveElm(List &L, infoMhs x) {
    adrMhs p = findElement(L, x);
    if (p == nullptr) {
        cout << "Tidak ditemukan" << endl;
        return;
    }

    if (p == L.first) {
        deleteFirst(L, p);
    } else if (p->next == nullptr) {
        deleteLast(L, p);
    } else {
        adrMhs q = L.first;
        while (q->next != p) {
            q = q->next;
        }
        deleteAfter(q, p);
    }
}

void showList(List L){
    adrMhs p;
    if(L.first == nullptr){
        cout << "List kosong";
    }else{
        p = L.first;
        while(p != nullptr){
            cout << "Nama : " << p->infotypemhs.nama << endl;
            cout << "Nim : " << p->infotypemhs.nim << endl;
            cout << "Prodi : " << p->infotypemhs.prodi << endl;
            cout << "=============================" << endl;
            p = p->next;
        }
    }
    cout <<endl;
}
