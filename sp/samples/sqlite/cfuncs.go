package main

/*
extern int callback(void*, int, char**, char**);
int callback_cgo(void* unused, int columnCount, char** columnTexts, char** columnsNames) {
   return callback(unused, columnCount, columnTexts, columnsNames);
}
*/
import "C"
