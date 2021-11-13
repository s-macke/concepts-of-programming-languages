# Exercise 7 - Distributed Computing

If you do not finish during the lecture period, please finish it as homework.

## Exercise 7.1 - TCP connection

Make a TCP connection with *towel.blinkenlights.nl* on port *23* and, read from the data stream and write the output onto the screen in an infinite loop.


## Exercise 7.2 - Remote procedure call

Implement in Go a net/rpc client TCP connection via rpc.Dial against the URL *simulationcorner.net:1234* . 
This interface offers two remote procedures calls.

A function, in which you receive a rune

    var dummy struct {}
    var output rune
    err := client.Call("Session.Read", &dummy, &output)

and a function in which you have to provide a rune

    var input rune
    err := client.Call("Session.Write", &input, &dummy)

Create for each routine its own goroutine and call these functions in an infinite loop. Write the output of *Session.Read* to the terminal. As input for *Session.Write* use os.Stdin to read a rune from keyboard via

    reader := bufio.NewReader(os.Stdin)
    input, _, err := reader.ReadRune()

## Exercise 7.3 - Call Rest API

Call the Github API with Go to receive the repositories with the most stars and the query "awesome"

https://api.github.com/search/repositories?sort=stars&order=desc&q=awesome

Use https://mholt.github.io/json-to-go/ to instantly create a Go structure from an arbitrary JSON.

## Exercise 7.4 - Microservice with API and embedded Web Site

The goal of this exercise is to create a REST API for a key value store.

```Go
    map[string]string
```

### Add an "add" entry API call

```
GET /api/add?key=mykey&value=myvalue
```

To test the rest interface you can use tools such as
- curl
- httpie
- Postman
- Your own Go Rest call code
- Web Browser

###  List entry API call

Add a Rest API GET call
```
GET /api/list 
```
call which returns the complete key value store as JSON

```JSON
{
    "mykey1": "myvalue1",
    "mykey2": "myvalue2",
    ...
}
```

if the response should be display in the browser you have to add the following Header
```HTTP
Content-Type: application/json
```

### Clear Store API call

Add a Rest API GET call
```
GET /api/clear
```
which removes the complete key value store

###  Include web site

Add web site as an embedded file system, which can interface the key value store

### Add a unit test

Use the httptest.NewRequest function to create a request and test the API

### POST call

Add a Rest API POST call /api/addAll call which takes a JSON as input

```Json
[{ 
    "key":  "MyKey",
    "value": "MyValue"
},{
    "key":  "NextKey",
    "value": "NextValue"
}]
```
