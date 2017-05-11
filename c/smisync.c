#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define BUFLEN 256
char buf[BUFLEN];
#define SYNCSTR    "<SYNC Start="

void usage(char *cmd)
{
  printf("USAGE: %s <start>+/-<range> < input.smi > output.smi\n", cmd);
  return;
}

int parse_option(char *optstr, int *start, int *range)
{
  char *rangestr;
  *start = strtol(optstr, &rangestr, 10);
  *range = strtol(rangestr, NULL, 10);
  return 0;
}

int main(int argc, char *argv[])
{
  int start, range, sync;
  char *eptr, *tptr;

  if(argc < 2) {
    usage(argv[0]);
    exit(0);
  }

  parse_option(argv[1], &start, &range);

  memset(buf, 0, BUFLEN);

  while(fgets(buf, BUFLEN, stdin) != NULL) {
    if(strncasecmp(buf, SYNCSTR, strlen(SYNCSTR)) == 0) {
      tptr = buf + strlen(SYNCSTR);
      sync = strtol(tptr, &eptr, 10);
      printf("%s%d%s", SYNCSTR, sync + ((sync >= start)? range: 0),
          eptr);
    }
    else
      printf("%s", buf);
  }

  return 0;
}
