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
#include <math.h>

#define CONVERT_BYTES_TO_GIGABYTES 107374182   

typedef struct {
  unsigned long mem_size;
} ram;

unsigned long get_total(void);
unsigned long get_usage(void);


unsigned long get_usage(void) 
{

  struct rusage usage;
  if(0 == getrusage(RUSAGE_SELF, &usage))
    return usage.ru_maxrss / CONVERT_BYTES_TO_GIGABYTES ; 
  else
    return 0;
}

unsigned long get_mem(void){

  int mib[2];
  size_t size;
  uint64_t ram_size;

  mib[0] = CTL_HW;
  mib[1] = HW_MEMSIZE;

  size = sizeof(ram_size);
  if (sysctl(mib, 2, &ram_size, &size, NULL, 0) == -1) {
    perror("sysctl");
  }

  return ram_size;

}

int main()
{
  printf("%lu", get_mem());
  return 0;
}


#endif

#ifdef __gnu_linux__

#include <stdio.h>
#include <stdlib.h>

typedef struct 
{
  unsigned long mem_size;
} ram;

unsigned long get_total(void);
unsigned long get_usage(void);


int main()
{

  return 0;
}

unsigned long get_total(void)
{

}

unsigned long get_usage(void) 
{

}



#endif

