/*
 * Fav
 *
 * The perfect tool to come up with a name for that new
 * tool that you've been planning on making. Generate
 * a list of words using
 *
 *          fav [-r (optional)]
 *
 * and use the first decent sounding three-letter word
 * as the name of your new tool.
 */
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

void usage(const char *error)
{
  printf("fav: %s\n", error);
  exit(1);
}

void combine(void)
{
  for (char a = 'a'; a <= 'z'; ++a) {
    for (char b = 'a'; b <= 'z'; ++b) {
      for (char c = 'a'; c <= 'z'; ++c) {
  	printf("%1c%1c%1c\t", a, b, c);
      }
    }
  }
}

int main(int argc, char *argv[])
{
  int recurse = 0;
  if (argc == 2)
  {
    if (!strcmp(argv[1], "-r")) {
      recurse = 1;
    } else {
      usage("fav [-r (optional)]");
    }
  }

recurse:
  combine();
  if (recurse) {
    goto recurse;
  }

  printf("\n");
}
