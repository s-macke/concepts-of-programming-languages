// pointer points to an integer at memory address 2048
char *data = (char*)2048;

void foo(char x) {
    // set memory address
    *data = x + 1;
}