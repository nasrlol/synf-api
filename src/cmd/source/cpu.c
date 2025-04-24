/*
 * =====================================================================================
 *
 *       Filename:  cpuc.c
 *
 *    Description: Retrieving cpu information from device
 *
 *        Version:  1.0
 *        Created:  04/08/2025 01:00:21
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  nasr,
 *   Organization:  synf
 *
 * ===================================================================================== */


typedef 
// OSX
#ifdef __APPLE__

#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>

#include <stdio.h>
#include <stdint.h>
#include <sys/types.h>
#include <sys/sysctl.h>

uint64_t get_cpu_freq(void)
{
    uint64_t freq = 0;
    size_t size = sizeof(freq);

    if (sysctlbyname("hw.cpufrequency", &freq, &size, NULL, 0) < 0)
    {
        perror("sysctl");
    }
    return freq;
}

char* get_cpu_name(void)
{
    char* cpu_name;
    size_t size = sizeof(cpu_name);
    if (sysctlbyname("ker.hostname", &cpu_name, &size, NULL, 0) < 0)
        perror("sysctl"); 

    return cpu_name;  
}

void get_cpu_temperature(void)
{
    uint16_t cpu_temperature;
    size_t size = sizeof(cpu_temperature);

    if(())
}

int main(int argc, char **argv)
{
    if (argc > 1)
    {
        if (strcmp(argv[1], "frequency") == 0)
            printf("argument received");
    }
    printf("%llu", get_cpu_freq());
    return 0;
}

#endif

// GNU/Linux
#ifdef __gnu_linux__

#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>

#include <stdio.h>
#include <stdint.h>
#include <sys/types.h>
// #include <sys/sysctl.h>
// macos file for getting system information

#define MAXC 1024
#define MAXC_CHAR 256

void cpu_name(void);
void cpu_temperature(unsigned short delay);
void cpu_frequency(unsigned short delay);

int main(int argc, char **argv)
{
    if (argc > 1)
    {
        if (strcmp(argv[1], "frequency") == 0)
        {
            cpu_frequency(1);
            printf("starting the process of getting the CPU frequency\n");
        }

        else if (strcmp(argv[1], "name") == 0)
        {
            printf("starting the process of getting the CPU name\n");
            cpu_name();
        }
        else if (strcmp(argv[1], "temperature") == 0)
        {
            printf("CPU temperature:\n");
            cpu_temperature(1);
        }
    }
    else
        printf("no arguments passed, try again with : frequency, temperature or name");
    return 0;
}

void cpu_name(void)
{
    int buffer_size = 256;
    char cpu_name[buffer_size];

    FILE *fp = fopen("/proc/cpuinfo", "r");
    if (!fp)
        printf("can't open /proc/cpuinfo");

    char line[buffer_size];
    while (fgets(line, sizeof(line), fp))
    {
        if (strncmp(line, "model name", 10) == 0)
        {
            char *colon = strchr(line, ':');
            if (colon)
            {
                snprintf(cpu_name, buffer_size, "%s", colon + 2);
                cpu_name[strcspn(cpu_name, "\n")] = 0;
                int err = fclose(fp);
                if (err != 0)
                    printf("error closing /proc/cpuinfo");

            }
        }
    }

    printf("%s", cpu_name);
    snprintf(cpu_name, buffer_size, "%s", cpu_name);
    // dont know what the snprintf is doing here but removing it gives a segmentation fault
    // so im keeping it here :)
}

void cpu_temperature(unsigned short delay)
{
    while (1)
    {
        sleep(delay);
        FILE *pf = fopen("/sys/class/thermal/thermal_zone0/temp", "r");
        // error handling in case of not being able to open /sys/class/thermal/..
        if (!pf)
            printf("error reading /proc/cpuinfo");

        char buffer[MAXC];
        while (fgets(buffer, sizeof(buffer), pf))
        {
            int a = atoi(buffer);
            a /= 1000;
            printf("%dC\n", a);
            fflush(stdout);
        }
        fclose(pf);
    }
}

void cpu_frequency(unsigned short delay)
{
    while(1) 
    {
        sleep(delay);
        char buffer[MAXC_CHAR];
        FILE *fp = fopen("/proc/cpuinfo", "r");
        if (!fp)
            printf("can't open /proc/cpuinfo");

        while (fgets(buffer, sizeof(buffer), fp))
        {

            if (strncmp(buffer, "cpu MHz", 10) == 0)
            {
                char *colon = strchr(buffer, ':');
                if (colon)
                {
                    snprintf(buffer, MAXC_CHAR, "%s", colon + 2);
                    buffer[strcspn(buffer, "\n")] = 0;
                    int err = fclose(fp);
                    if (err != 0)
                        printf("error closing /proc/cpuinfo");

                }
            }
        }

        fflush(stdout);
        printf("%s", buffer);
        snprintf(buffer, MAXC_CHAR, "%s", buffer);
        // dont know what the snprintf is doing here but removing it gives a segmentation fault
        // so im keeping it here :)
    }
}



#endif
