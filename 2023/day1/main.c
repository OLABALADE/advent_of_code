#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main() {
  FILE *file = fopen("input.txt", "r");
  if (file == NULL) {
    printf("Error opening file");
    return 1;
  }

  char line[256];
  int sum = 0;
  while (fgets(line, sizeof(line), file)) {
    char fnum, lnum;
    int is_first = 1;
    for (int i = 0; i < strlen(line); i++) {
      if (isdigit(line[i])) {
        if (is_first) {
          fnum = line[i];
          is_first = 0;
        }
        lnum = line[i];
      }
    }
    char num[] = {fnum, lnum, '\0'};
    int value = atoi(num);
    sum += value;
  }

  printf("Part 1:%d", sum);

  fclose(file);
  return 0;
}
