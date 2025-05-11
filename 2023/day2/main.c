#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main() {
  FILE *file = fopen("input.txt", "r");
  if (file == NULL) {
    printf("Error opening file");
    return 1;
  }

  char *saveptr1, *saveptr2;
  char *token;
  bool posssible;

  char line[256];
  int game_id;
  int sum = 0;
  int sum_of_power = 0;

  while (fgets(line, sizeof(line), file)) {
    posssible = true;
    int max_red = 0, max_green = 0, max_blue = 0;
    token = strtok_r(line, ":", &saveptr1);
    sscanf(token, "Game %d", &game_id);

    token = strtok_r(NULL, ":", &saveptr1);
    char *game_sets = token;
    token = strtok_r(game_sets, ";", &saveptr2);

    while (token != NULL) {
      int num_red = 0;
      int num_green = 0;
      int num_blue = 0;

      char *cubes = token;
      char *cube_token = strtok_r(token, ",", &saveptr1);

      while (cube_token != NULL) {
        char color[10];
        int count;

        sscanf(cube_token, "%d %s", &count, color);
        if (strcmp(color, "red") == 0) {
          num_red = count;
        } else if (strcmp(color, "green") == 0) {
          num_green = count;
        } else if (strcmp(color, "blue") == 0) {
          num_blue = count;
        }
        cube_token = strtok_r(NULL, ",", &saveptr1);
      }

      if (num_red > 12 || num_green > 13 || num_blue > 14) {
        posssible = false;
      }

      max_red = (max_red < num_red) ? num_red : max_red;
      max_green = (max_green < num_green) ? num_green : max_green;
      max_blue = (max_blue < num_blue) ? num_blue : max_blue;

      token = strtok_r(NULL, ";", &saveptr2);
    }

    if (posssible) {
      sum += game_id;
    }
    sum_of_power += max_red * max_green * max_blue;
  }
  printf("Sum is %d\n", sum);
  printf("Sum of power is %d\n", sum_of_power);
  fclose(file);
}
