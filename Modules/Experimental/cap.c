#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <termios.h>
#include <unistd.h>

#define MAX_HISTORY 100
#define MAX_INPUT 100

// Function to set the terminal in raw mode
void enableRawMode() {
    struct termios raw;
    tcgetattr(STDIN_FILENO, &raw);
    raw.c_lflag &= ~(ECHO | ICANON);
    tcsetattr(STDIN_FILENO, TCSAFLUSH, &raw);
}

// Function to restore the terminal's original settings
void disableRawMode() {
    struct termios raw;
    tcgetattr(STDIN_FILENO, &raw);
    raw.c_lflag |= (ECHO | ICANON);
    tcsetattr(STDIN_FILENO, TCSAFLUSH, &raw);
}

int main() {
    enableRawMode();

    char history[MAX_HISTORY][MAX_INPUT];
    int history_count = 0;
    int history_index = 0;
    char input[MAX_INPUT] = "";
    int input_length = 0;
    int cursor_pos = 0;

    printf("Type your commands. Press 'Ctrl + D' to quit.\n");

    while (1) {
        char c = getchar();

        if (c == EOF || c == 4) { // Ctrl + D (End of File)
            break; // Exit the loop and quit
        } else if (c == 27) {
            // If the character is the escape sequence for arrow keys
            char arrow[3];
            if (scanf("[%1s", arrow) != 1)
                continue;

            if (arrow[0] == '1') { // Home or End key
                scanf("%1s", arrow + 1);
                if (arrow[1] == ';') {
                    scanf("%1s", arrow + 2);
                }
            }

            switch (arrow[0]) {
                case 'A': // Up arrow
                    if (history_index > 0) {
                        history_index--;
                        strncpy(input, history[history_index], sizeof(input));
                        input_length = strlen(input);
                        cursor_pos = input_length;
                    }
                    break;
                case 'B': // Down arrow
                    if (history_index < history_count - 1) {
                        history_index++;
                        strncpy(input, history[history_index], sizeof(input));
                        input_length = strlen(input);
                        cursor_pos = input_length;
                    } else {
                        input[0] = '\0';
                        input_length = 0;
                        cursor_pos = 0;
                    }
                    break;
                case 'C': // Right arrow
                    if (cursor_pos < input_length) {
                        cursor_pos++;
                    }
                    break;
                case 'D': // Left arrow
                    if (cursor_pos > 0) {
                        cursor_pos--;
                    }
                    break;
                case '1': // Home key
                    cursor_pos = 0;
                    break;
                case '4': // End key
                    cursor_pos = input_length;
                    break;
            }
        } else if (c == '\n') {
            // Enter key: execute the command and add it to history
            if (input_length > 0) {
                printf("Executing: %s\n", input);

                if (history_count < MAX_HISTORY) {
                    strncpy(history[history_count], input, sizeof(input));
                    history_count++;
                } else {
                    for (int i = 1; i < MAX_HISTORY; i++) {
                        strncpy(history[i - 1], history[i], sizeof(input));
                    }
                    strncpy(history[MAX_HISTORY - 1], input, sizeof(input));
                }

                history_index = history_count;
                input[0] = '\0';
                input_length = 0;
                cursor_pos = 0;
            }
        } else if (c == 127) {
            // Backspace (127 is ASCII code for backspace)
            if (cursor_pos > 0) {
                for (int i = cursor_pos - 1; i < input_length; i++) {
                    input[i] = input[i + 1];
                }
                cursor_pos--;
                input_length--;
            }
        } else if (input_length < MAX_INPUT - 1) {
            // Regular character input
            for (int i = input_length; i > cursor_pos; i--) {
                input[i] = input[i - 1];
            }
            input[cursor_pos] = c;
            cursor_pos++;
            input_length++;
        }

        // Move the cursor to the beginning of the line and print the current input
        printf("\r\033[K%s", input);
        // Move the cursor to the current position
        printf("\r\033[%dC", cursor_pos + 1);
        fflush(stdout);
    }

    disableRawMode();
    return 0;
}
