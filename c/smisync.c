#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define BUFLEN 256
char buf[BUFLEN];
#define SYNCSTR		"<SYNC Start="

void usage(char *cmd)
{
	printf("USAGE: %s <start>+/-<range> < input.smi > output.smi\n", cmd);
	return;
}

int main(int argc, char *argv[])
{
	int start, range, sync;
	char *sptr, *rptr, *eptr;
	char *tptr;

	if(argc < 2) {
		usage(argv[0]);
		exit(0);
	}

	sptr = argv[1];
	start = strtol(sptr, &rptr, 10);
	range = strtol(rptr, &eptr, 10);

	memset(buf, 0, BUFLEN);

	while(fgets(buf, BUFLEN, stdin) != NULL) {
		if(strncasecmp(buf, SYNCSTR, strlen(SYNCSTR)) == 0) {
			printf("%s", SYNCSTR);

			tptr = buf + strlen(SYNCSTR);
			sync = strtol(tptr, &eptr, 10);
			if(sync >= start)
				printf("%d", sync+range);
			else
				printf("%d", sync);

			printf("%s", eptr);
		}
		else
			printf("%s", buf);
	}

	return 0;
}
