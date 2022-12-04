#include<unistd.h>
#include<stdio.h>

int main() {
    char buffer[256];
    gethostname(buffer, 256);
    printf("%s\n", buffer);
    return 0;
}