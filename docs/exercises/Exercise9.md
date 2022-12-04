# Exercise 9 - Systems Programming

If you do not finish during the lecture period, please finish it as homework.

## Exercise 9.1 - Cgo

Write a Cgo program that reads the hostname from the libc function `gethostname` and prints it to the console.
The corresponding C function is defined in `unistd.h` and can be used as follows:

```C
#include<unistd.h>
#include<stdio.h>

int main() {
    char buffer[256];
    gethostname(buffer, 256);
    printf("%s\n", buffer);
    return 0;
}
```

Write a function `GetHostname` that calls the C function and returns the hostname as a Go string.

```Go
func getHostname() string {
// TODO implement
}
```

Use C.char for the type char and C.GoString to convert the resulting C string to a Go string.

## Exercise 9.2 - Cgo

In the same way as in the previous exercise write a Cgo program that
creates a new SQLITE database file and executes the following statements:

```sqlite
create table if not exists students
(
    id   integer not null primary key autoincrement,
    name text
);
insert into students (name)
values ('Your Name');
insert into students (name)
values ('Another Name');
```

Write your own Cgo wrapper. Verify the content of the database
using any SQLITE inspector (e.g. IntelliJ or https://sqlitebrowser.org/).

Windows users can use the Windows Subsystem for Linux and install the following packages:

```shell script
sudo apt install gcc golang sqlite3 libsqlite3-dev
```

Here is an example code of how to use the libs

```C
// compile with gcc sqlite.c -lsqlite3
#include <stdio.h>
#include <stdlib.h>
#include <sqlite3.h>

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
```

Hints:

* You can define your own database object with

```Go
var db *C.sqlite3
```

* Use `C.CString` to convert a Go string to a C string

## Exercise 9.3 - FUSE Filesystem

Choose either the FUSE Filesystem Exercise or the Containerer Exercise.

### Try out the test filesystem

Experiment with the test filesystem.

https://github.com/s-macke/concepts-of-programming-languages/blob/master/src/system/fuse/main.go

#### Write a RAM file system

- Try to make the INode of each file unique or limit to one file first.
- Implement "write" to file capabilities. Store the content in RAM.
- Allow to create new files inside the directory.

## Exercise 9.4 - Containerization

### Try out chroot

Download the alpine distribution
at https://dl-cdn.alpinelinux.org/alpine/v3.16/releases/x86_64/alpine-minirootfs-3.16.3-x86_64.tar.gz
Then create a new directory and extract the tarball into it and run chroot.

```shell script
# prepare chroot
mkdir -p jail
cd jail
tar -xvzf PATH/TO/alpine-minirootfs-3.16.3-x86_64.tar.gz
cd ..
# run chroot and start the shell via /bin/sh
sudo chroot jail /bin/sh
# mount the process information pseudo-file system
mount -t proc none /proc
```
Look at the directory structure of the chroot environment and check, that you are inside of the jail.
Return to the host system with `exit`

### 

**Questions:**

- What is your user id inside and outside of the container? (Hint `id`)
- Can you see the processes outside of the chroot environment? (Hint: `ps aux`)
- What is the pid (process id) inside and outside of the container of your shell `/bin/sh` ? (Hint `pstree -p` or `ps aux`)

### Building your own containerization system

#### 1. Start with a simple main file:

```go
func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("help")
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
```

#### 2. Create your own non-isolated container command

```go
func run() {
	log.Printf("Running %v \n", os.Args[1:])

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(cmd.Run())
}
```

#### 3. Extend your program by setting some sysflags

```go
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWUSER | syscall.CLONE_NEWUTS,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},
	}

```

**Questions:**

- Try changing the hostname again. What is the hostname if you exit the container?
- Call `ps`. What do you see?

#### 4. Try to become more isolated using exec/fork

- Copy the run function and name the copy child().
- Remove cmd.SysProcAttr in the child() function.
- Modify the run() function to execute the following instead:

```go
cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
```

**Questions:**

- Call `ps`. What do you see?

#### 5. Use a custom root filesystem

Use the alpine mini root as in the first part of the exercise.

In the child() function, add the following code:

```go
must(syscall.Chroot("/jail")) // local fs
must(os.Chdir("/"))
must(syscall.Mount("proc", "proc", "proc", 0, ""))

must(cmd.Run())

must(syscall.Unmount("proc", 0))
```

**Questions:**

- Call `ps`. What do you see?
