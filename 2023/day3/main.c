#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_LINES 140
#define MAX_LINE_LENGTH 142

int extract_number(char *p, int *offset) {
  char *start = p;
  char *end = p;

  while (start > p - strlen(p) && isdigit(*(start - 1))) {
    start--;
  }

  while (isdigit(*end)) {
    end++;
  }

  int len = end - start;
  *offset = p - start;

  char num[20];
  strncpy(num, start, len);
  num[len] = '\0';

  return atoi(num);
}

int main() {
  FILE *file = fopen("input.txt", "r");
  char grid[MAX_LINES][MAX_LINE_LENGTH];
  int seen[MAX_LINES][MAX_LINE_LENGTH] = {0};
  int parts_sum = 0;
  int gear_ratio_sum = 0;
  int dirs[8][2] = {{0, -1}, {1, -1}, {1, 0},  {1, 1},
                    {0, 1},  {-1, 1}, {-1, 0}, {-1, -1}};

  int i = 0;
  // Read File into grid array
  while (fgets(grid[i], MAX_LINE_LENGTH, file) != NULL && i < MAX_LINES) {
    size_t len = strcspn(grid[i], "\n");
    grid[i][len] = '\0';
    i++;
  }

  for (int j = 0; j < MAX_LINES; j++) {
    for (int k = 0; k < MAX_LINE_LENGTH; k++) {

      if (grid[j][k] != '.' && grid[j][k] != '\0' && !isdigit(grid[j][k])) {
        int gear_ratio = 1, count = 0;

        for (int l = 0; l < 8; l++) {
          int dx = j + dirs[l][0];
          int dy = k + dirs[l][1];

          if (dx >= 0 && dx < MAX_LINES && dy >= 0 && dy < MAX_LINE_LENGTH) {
            char *c = &grid[dx][dy];
            int offset;

            if (isdigit(*c)) {
              int num = extract_number(c, &offset);

              if (!seen[dx][dy - offset]) {
                if (grid[j][k] == '*') {
                  gear_ratio *= num;
                  count++;
                }
                parts_sum += num;
                seen[dx][dy - offset] = 1;
              }
            }
          }
        }
        if (count == 2 && grid[j][k] == '*')
          gear_ratio_sum += gear_ratio;
      }
    }
  }
  printf("Part 1: Sum of parts is %d\n", parts_sum);
  printf("Part 1: Sum of gear ratios is %d", gear_ratio_sum);
  fclose(file);
}
