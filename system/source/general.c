/*
 * =====================================================================================
 *
 *       Filename:  general.c
 *
 *    Description: Retrieving basic information about the device
 *
 *        Version:  1.0
 *        Created:  05/08/2025 22:22:00
 *       Revision:  none
 *       Compiler:  clang 
 *
 *         Author:  nasr,
 *   Organization:  synf
 *
 * ===================================================================================== */

// GNU LINUX 

#ifdef __gnu_linux__

#include <stdio.h>
#include <stdlib.h>
#include <sys/sysinfo.h>

long device_up_time(void);

int main(int argc, char** argv)
{
    printf("The total uptime is (seconds): %lu", device_up_time());
    return 0;
}

long device_up_time(void)
{
    struct sysinfo info;
    if (sysinfo(&info) == -1)
        perror("sysinfo");

    return info.uptime;
}

#endif


