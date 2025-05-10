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
 * =====================================================================================*/

// OSX
#ifdef __APPLE__

#include <stdlib.h>
#include <stdio.h>
#include <sys/sysctl.h>
#include <sys/types.h>

typedef struct {
  unsigned long mem_size;
  unsigned long av_mem_size;
} ram;

long getTotalMem(void);
int getMemFreq(void);
long getMemoryUsage(void);


long getMemoryUsage(void) 
{

  struct ramusage usage;
  if(0 == getRamUsage(RAM_USAGE_SELF, &usage))
    return usage.ru_maxrss; // returns the value in bytes
                            // convert it to gigabytes in the front end :) 
  else
    return 0;
}

long getTotalMem(void){

  int i, mib[4];
  size_t len;

  len = 4;


  return 0;
}

int main()
{
  printf("%lu", getMemoryUsage());
  return 0;
}


#endif

#ifdef __gnu_linux__

#include <stdio.h> 
#include <string.h>
#include <stdlib.h>
#include <sys/sysinfo.h>


typedef struct {
  unsigned long mem_size;
  unsigned long av_mem_size;
} ram;

unsigned long tot_mem_size(void);
unsigned long av_mem_size(void);

int main(int argc, char** argv)
{
  if (argc < 1)
  {
    if (strcmp(*argv, "size")) {
      printf("%lu", tot_mem_size());
    } else if (strcmp(*argv, "available"))
    {
      printf("%lu", av_mem_size());
    } else if (strcmp(*argv, "frequency")){
      printf("%f", mem_freq());
    }
  }
  return 0;
}

unsigned long tot_mem_size(void) 
{
  struct sysinfo info;

  return info.totalram;
}

unsigned long av_mem_size(void)
{
  struct sysinfo info;

  return info.freeram;
}

#endif
