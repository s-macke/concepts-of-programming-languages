#include <stdio.h>
#include <stdlib.h>
#include <sqlite3.h>

// compile with gcc sqlite.c -lsqlite3

int main() {
    sqlite3 *db;
    int rc = sqlite3_open("test.db", &db);
    if (rc != SQLITE_OK) {
        printf("sqlite3_open returned code %i", rc);
        return 1;
    }

    sqlite3_close(db);
    return 0;
}