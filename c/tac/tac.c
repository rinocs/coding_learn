#include <stdio.h>
#include <stdlib.h>
#include <string.h>


struct line{
    char *s;
    struct line *next;
};

int main(int argc, char **argv){ 
    if (argc != 2) {
        printf("Missing file operand\n");
        return 1;
    }

    FILE *file = fopen(argv[1], "r");
    if (!file) {
        printf("Error opening file");
        return 1;
    }

    struct line *head = NULL;

    char buffer[1024];
    while (fgets(buffer, sizeof(buffer), file) != NULL) {
        struct line *new_line = malloc(sizeof(struct line));
        size_t len = strlen(buffer);
        new_line->s = malloc(len + 1);
        for (size_t i = 0; i <= len ; i++) {
            new_line->s[i] = buffer[i];
        }
        new_line->next = head;
        head = new_line;
    }

    struct line *current = head;
    while (current != NULL) {
        printf("%s", current->s);
        struct line *temp = current;
        current = current->next;
        free(temp->s);
        free(temp);
    }

    fclose(file);
    return 0;
}