#include <stdio.h>
#include <stdlib.h>
#include <sqlite3.h>
// compile with gcc sqlite.c -lsqlite3

void Exec(sqlite3 *db, char *sql) {
    int rc = sqlite3_exec(db, sql, NULL, NULL, NULL);
    if (rc != SQLITE_OK) {
        printf("sqlite3_open returned code %i", rc);
        exit(1);
    }
}
int main() {
    sqlite3 *db;
    int rc = sqlite3_open("test.db", &db);
    if (rc != SQLITE_OK) {
        printf("sqlite3_open returned code %i", rc);
        return 1;
    }
    Exec(db, "create table if not exists students (id integer not null primary key autoincrement, name text);");
    Exec(db, "insert into students (name) values ('Your Name');");
    Exec(db, "insert into students (name) values ('Another Name');");
    sqlite3_close(db);
    return 0;
}