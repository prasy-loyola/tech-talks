#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

struct Node {
  char value;
  struct Node *next;
};

typedef struct Node Node;

void printLinkedList(Node *node) {
  if (node == NULL) {
    return;
  }
  printf("Node: value: %c, next: %p\n", node->value, node->next);

  printLinkedList(node->next);
}

Node *head = NULL;
Node *tail = NULL;

#define LOG_DEBUG false

void readUserInput() {
  char value = 0;

  scanf("%c", &value);
  while (value != '\n') {
    LOG_DEBUG &&printf("value: %c\n", value);
    Node *node = malloc(sizeof(Node));
    node->value = value;
    LOG_DEBUG &&printf("head: %p, tail: %p\n", head, tail);
    // head, tail - nil
    if (head == NULL && tail == NULL) {
      head = node;
      tail = node;
    } else {
      // head - value , tail - value
      tail->next = node; // Add element to the next of tail
      tail = node;       // make new element the tail
    }
    scanf("%c", &value);
  }
}

bool findCharacterInLinkedList(Node *head, char searchChar) {
  if (head == NULL) {
    return false;
  }
  if (head->value == searchChar) {
    return true;
  }
  return findCharacterInLinkedList(head->next, searchChar);
}

int main() {
  printf("Enter a text: ");
  readUserInput();
  char searchChar;
  printf("Enter a character to search: ");
  if (LOG_DEBUG) {
    printLinkedList(head);
  }
  scanf("%c", &searchChar);
  printf("Searching for '%c'\n", searchChar);

  bool found = findCharacterInLinkedList(head, searchChar);
  if (found) {
    printf("Found '%c' in the text\n", searchChar);
  } else {
    printf("Did not find '%c' in the text\n", searchChar);
  }
}
