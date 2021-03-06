# Exercise 9 - Systems Programming

If you do not finish during the lecture period, please finish it as homework.

## Exercise 9.1 - Cgo

Write a Cgo program that creates a new SQLITE database file and executes the following statements:
```sqlite
create table if not exists students (id integer not null primary key autoincrement, name text);
insert into students (name) values ('Your Name');
insert into students (name) values ('Another Name');
```
Do not use any GitHub libraries, but write your own Cgo wrapper.
Verify the content of the database using any SQLITE inspector (e.g. IntelliJ or https://sqlitebrowser.org/).

Windows users can use the Windows Subsystem for Linux and install the following packages:
```shell script
sudo apt install gcc golang sqlite3 libsqlite3-dev
```

## Exercise 9.2 - Containerization

### Try out chroot
```shell script
# prepare chroot
mkdir -p jail/{bin,lib,lib64}
cp -v /bin/{bash,ls,touch,rm} jail/bin
for i in $(ldd /bin/bash | egrep -o '/lib.*.\.[0-9]'); do cp -v --parents "$i" jail; done
for i in $(ldd /bin/ls | egrep -o '/lib.*.\.[0-9]'); do cp -v --parents "$i" jail; done
for i in $(ldd /bin/touch | egrep -o '/lib.*.\.[0-9]'); do cp -v --parents "$i" jail; done
for i in $(ldd /bin/rm | egrep -o '/lib.*.\.[0-9]'); do cp -v --parents "$i" jail; done

# run chroot
sudo chroot jail
```

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

**Questions:**
- Try changing the hostname. What is the hostname if you exit the container?
- What is the pid?


#### 3. Extend your program by setting some sysflags
```go
cmd.SysProcAttr = &syscall.SysProcAttr{
    Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
    Unshareflags: syscall.CLONE_NEWNS,
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
Create a complete chroot environment. For Debian based linux systems, you might use e.g.
```shell script
sudo debootstrap buster /jail http://deb.debian.org/debian
```

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

#### 6. Resource limit

Use cgroups (`/sys/fs/cgroup/`) to limit the container resource consumption.
Test your program!
