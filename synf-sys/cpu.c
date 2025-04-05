#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

char *cpu_name()
{
    char *name;
    // /proc/cpuinfo gives a lot of information, i have to parse it and get the model name of the cpu and return that
    FILE *pf = fopen("proc/cpuinfo", "r");
    if(!pf)
    {
        // error handling, in case of not being to open /proc/cpuinfo
        printf("error reading the cpu frequency");
    }


    return name;
}
int cpu_temperature(int delay)
{
    while (1)
    {
        sleep(delay);
        FILE *pf = fopen("/sys/class/thermal/thermal_zone0/temp", "r");
        // error handling in case of not being able to open /sys/class/thermal/..
        if (!pf)
        {
            printf("error reading /proc/cpuinfo");
            return 1;
        }

        char buffer[1024];
        while (fgets(buffer, sizeof(buffer), pf))
        {
            int a = atoi(buffer);
            a /= 1000;
            printf("CPU Temp: %dC\n", a);
        }
        fclose(pf);
    }
    return 0;
}

int cpu_frequency()
{
    while (1)
    {
        FILE *pf = fopen("/proc/cpuinfo", "r");
        // error handling in case of not being able to open the file
        if (!pf)
        {
            printf("error reading /proc/cpuinfo");
        }

        char buffer[1024];
        while (fgets(buffer, sizeof(buffer), pf))
        {
            int a = atoi(buffer);
            printf("CPU FREQ: %d\n", a);
        }
        fclosee(pf);
    }
    return 0;
}

int main(int argc, char **argv)
{
    if (argc > 1)
    {
        cpu_temperature(atoi(argv[1]));
    }
    else
    {
        cpu_temperature(1);
    }
    return 0;
}
