#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_LINE_LENGTH 256

int main() {
  FILE *file = fopen("input.txt", "r");
  char line[MAX_LINE_LENGTH];

  int point_sum = 0;

  while (fgets(line, MAX_LINE_LENGTH, file)) {
    int winning_numbers[10] = {0};
    int given_numbers[25] = {0};
    int card_num;

    char *token = strtok(line, ":");
    sscanf(token, "Card %d", &card_num);

    token = strtok(NULL, ":");
    char *wnums = strtok(token, "|");
    char *gnums = strtok(NULL, "|");

    token = strtok(wnums, " ");
    for (int i = 0; token != NULL; i++) {
      winning_numbers[i] = atoi(token);
      token = strtok(NULL, " ");
    }

    token = strtok(gnums, " ");
    int count = 0;
    while (token != NULL) {
      int num = atoi(token);
      for (int j = 0; j < 10; j++) {
        if (winning_numbers[j] == num) {
          count++;
        }
      }
      token = strtok(NULL, " ");
    }
    if (count != 0)
      point_sum += pow(2, count - 1);
  }
  printf("Total sum of point is %d", point_sum);
}
