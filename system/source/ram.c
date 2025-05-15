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
  char* memName;
  int memSize;
} ram;

long getTotalMem(void);
int getMemFreq(void);
long getMemoryUsage(void);


long getMemoryUsage(void) 
{

  struct rusage usage;
  if(0 == getrusage(RUSAGE_SELF, &usage))
    return usage.ru_maxrss / CONVERT_BYTES_TO_GIGABYTES ; 
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

<<<<<<< HEAD
#include <stdio.h> 
#include <string.h>
#include <stdlib.h>
#include <sys/sysinfo.h>

typedef struct {
  char* mem_name;
  long mem_size;
  long av_mem_size;
  double mem_freq;
} ram;

unsigned long tot_mem_size(void);
unsigned long av_mem_size(void);
float mem_freq(void);

int main(int argc, char** argv)
=======
int main()
>>>>>>> refs/remotes/origin/main
{

  return 0;
}
#endif
