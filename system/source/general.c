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

// OSX 
#ifdef __APPLE__

#include <sys/types.h>
#include <sys/sysctl.h>
#include <stdio.h>

long device_up_time(void);

int main(int argc, char** argv)
{

    return 0;
}

long device_up_time(void){

    struct timeval boottime;
    size_t len = sizeof(boottime);

    if (sysctlbyname("kern.boottime", &boottime, len, NULL, 0) == -1){
        perror("sysctl error");
        return 1;
    }
}

#endif

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


