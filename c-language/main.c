#include "arithmetic2.h"
#include <inttypes.h>
#include <stdbool.h>
#include <stdio.h>

#define NUMBER 101

int main() {

    char* name1 = "Prasanna\0";
    char* name2 = "Sriram\0";

    printf("Name1: %s Addr: %p\n", name1, &name1);
    printf("Name2: %s Addr: %p\n", name2, &name2);

}
