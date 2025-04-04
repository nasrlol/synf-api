#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int cpu_temperature(int delay)
{
    while (1) {

        sleep(delay);
        FILE *pf = fopen("/sys/class/thermal/thermal_zone0/temp", "r");
        if (!pf)
        {
            printf("error reading /proc/cpuinfo");
            return 1;
        }

        char buffer[1024];
        while (fgets(buffer, sizeof(buffer),pf))
        {
            int a = atoi(buffer);
            a /= 1000;
            printf("CPU Temp: %dC\n", a);
        }
        fclose(pf);
    }
    return 0;
}
int main(int argc, char** argv)
{
    if (argc > 1)
    {
        cpu_temperature(atoi(argv[1]));
    } else 
    {
        cpu_temperature(1);
    }
    return 0;
}
