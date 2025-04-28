/*
 * =====================================================================================
 *
 *       Filename:  ram.c
 *
 *    Description: retrieve ram information from the device 
 *
 *        Version:  1.0
 *        Created:  04/08/2025 01:34:33
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  nasr 
 *   Organization:  synf 
 *
 * =====================================================================================
 */
#include <stdlib.h>
#include <stdio.h>


struct cpu {
    char* memName;
    int memSize;
};

int getFreeMem(void);
int getTotalMem(void);
int getMemFreq(void);

int main()
{
    return 0;
}

int getFreeMem(){
    printf("hello world");
    return 0;
}