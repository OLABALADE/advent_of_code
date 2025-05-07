#include <ctype.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
  char *spelling;
  char value;
} Number;

int part_1(FILE *file) {
  char line[256];
  int sum = 0;

  while (fgets(line, sizeof(line), file)) {
    char fnum, lnum;
    bool is_first = true;

    for (int i = 0; i < strlen(line); i++) {
      if (isdigit(line[i])) {
        if (is_first) {
          fnum = line[i];
          is_first = false;
        }
        lnum = line[i];
      }
    }

    char num[] = {fnum, lnum, '\0'};
    int value = atoi(num);
    sum += value;
  }
  return sum;
}

int part_2(FILE *file) {
  char line[256];
  int sum = 0;

  Number nums[] = {{"one", '1'},   {"two", '2'},   {"three", '3'},
                   {"four", '4'},  {"five", '5'},  {"six", '6'},
                   {"seven", '7'}, {"eight", '8'}, {"nine", '9'}};

  int nums_len = sizeof(nums) / sizeof(nums[0]);

  while (fgets(line, sizeof(line), file)) {
    char fnum, lnum;
    bool is_first = 1;

    for (int i = 0; i < strlen(line); i++) {
      if (isdigit(line[i])) {
        if (is_first) {
          fnum = line[i];
          is_first = false;
        }
        lnum = line[i];
        continue;
      }

      for (int j = 0; j < nums_len; j++) {
        if (strncmp(&line[i], nums[j].spelling, strlen(nums[j].spelling)) ==
            0) {
          if (is_first) {
            fnum = nums[j].value;
            is_first = false;
          }
          lnum = nums[j].value;
        }
      }
    }

    char combined_num[] = {fnum, lnum, '\0'};

    sum += atoi(combined_num);
  }
  return sum;
}

int main() {
  FILE *file = fopen("input.txt", "r");
  if (file == NULL) {
    printf("Error opening file");
    return 1;
  }

  printf("Part 1:%d\n", part_1(file));
  printf("Part 2:%d", part_2(file));
  fclose(file);
  return 0;
}
