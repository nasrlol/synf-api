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
#include <IOKit/IOKitLib.h>
#include <CoreFoundation/CoreFoundation.h>

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


void get_disk()
{
    CFMutableDictionaryRef match_dict = IOServiceMatching("IOMedia");
    if (!match_dict) {
        printf("failed to create match directory");
    }

    CFDictionarySetValue(match_dict, CFSTR("Whole"), kCFBooleanTrue);

    io_iterator_t iter;
    kern_return_t kr = IOServiceGetMatchingServices(kIOMainPortDefault, match_dict, &iter);
    if (kr != KERN_SUCCESS) {
        printf("Error matching services");
    }

    io_object_t service;
    while ((service = IOIteratorNext(iter))) {
        CFStringRef bsdName = IORegistryEntryCreateCFProperty(service, CFSTR("BSD Name"), kCFAllocatorDefault, 0);
        if (bsdName) {
            char name[1024];
            CFStringGetCString(bsdName, name, sizeof(name), kCFStringEncodingUTF8);
            printf("Disk: /dev/%s\n", name);
            CFRelease(bsdName);
        }

        CFNumberRef sizeRef = IORegistryEntryCreateCFProperty(service, CFSTR("Size"), kCFAllocatorDefault, 0);
        if (sizeRef) {
            long long size = 0;
            CFNumberGetValue(sizeRef, kCFNumberLongLongType, &size);
            printf("  Size: %lld bytes\n", size);
            CFRelease(sizeRef);
        }

        IOObjectRelease(service);
    }

    IOObjectRelease(iter);
}


int main(int argc, char** argv)
{
    return 0;
}

