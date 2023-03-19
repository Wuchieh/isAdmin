#include <stdio.h>
#include <stdlib.h>
#include "isAdmin.h"

// build gcc example.c -o exampleC.exe -L. -lisAdmin

int main() {
    if (isAdmin()){
        printf("Hi Administrator\n");
    }else{
        printf("Hi Friend\n");
    }
    system("PAUSE");
}