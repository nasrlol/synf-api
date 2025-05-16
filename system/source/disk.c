/*
 * =====================================================================================
 *
 *       Filename:  disk.c
 *
 *    Description:  retrieving disk information from the device 
 *
 *        Version:  1.0
 *        Created:  04/08/2025 01:33:30
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

void disk_size();
char* disk_partitions();

typedef struct {

    char* partition_name;
    unsigned long parition_size;

} partition;

typedef struct {
    partition* paritions;
    long total_disk_size; 
} disk;


int main()
{
    return 0;
}

