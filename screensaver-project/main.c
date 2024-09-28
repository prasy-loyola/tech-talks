#include <stdbool.h>
#include <stdio.h>

#define SLOTS 40

void delay(int length) {
  for (int i = 0; i < 100000; i++) {
    for (int j = 0; j < 1000 * length; j++) {
        continue;
    }
  }
}

int main() {

  int k_position = 0; // assignment statement
  bool move_left = false;

  while (k_position < SLOTS && k_position >= 0) {

    // Print Pallanguzhi
    int count = 0;
    while (count < SLOTS) {
      if (count == k_position) {
        printf("#");
      } else {
        printf(".");
      }
      count = count + 1;
    }

    printf("\n\033[A");
    printf("\033[%dD", SLOTS);

    // At the right End
    if (k_position == SLOTS - 1) {
      move_left = true;
    }

    // At the left End
    if (k_position == 0) {
      move_left = false;
    }

    // Move k
    if (move_left) {
      k_position = k_position - 1;
    } else {
      k_position = k_position + 1;
    }
    delay(2);
  }
}
