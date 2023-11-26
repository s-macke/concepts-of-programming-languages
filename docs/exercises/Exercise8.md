# Exercise 8 - Network Programming

If you do not finish during the lecture period, please finish it as homework.

## Exercise 8.1 - TCP connection

Open a TCP connection to *simulationcorner.net* on port *54321*, read from the data stream and write the output onto the screen in an infinite loop.

## Exercise 8.2 - TCP two way communication

Open a TCP connection to *simulationcorner.net* on port *54322* and,
read from the data stream and write the output onto the screen in an infinite loop.
In addition, read from the keyboard and send the input to the server in an infinite loop.

Create for each routine its own goroutine.
As input use os.Stdin to read a rune from keyboard via
```Go
    reader := bufio.NewReader(os.Stdin)
    ....
    input, err := reader.ReadString('\n')
```

## Exercise 8.3 - Call Rest API

Call the Github API to receive the repositories with the most stars and the query "awesome"

https://api.github.com/search/repositories?sort=stars&order=desc&q=awesome

Use [https://mholt.github.io/json-to-go/](json-to-go) to instantly create a Go structure from an arbitrary JSON.

## Exercise 8.4 Modules and GIN

Use the teacher-bot linked in the learning campus to learn how to use GIN.

## Exercise 8.5 - Microservice with API and embedded Web Site

Create a REST API for a simple in-memory key value store for strings.

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

If the response should be displayed in the browser you have to add the following Header
```HTTP
Content-Type: application/json
```

You should be able to see the JSON in the browser when you try to access http://localhost:8080/api/list

### Clear Store API call

Add a Rest API GET call
```
GET /api/clear
```
which removes the complete key value store

### POST call

Currently we are only using the GET call of HTTP. GET does not have any payload. E. g. the data must be in the URL 
But the advantage is, that we can easily test the API with a browser.

Add a Rest API POST call /api/addAll call which takes a JSON as input.

```Json
{ 
    "MyKey": "MyValue",
    "NextKey": "NextValue"
}
```

You can test via httpie

```
http -v POST "http://localhost:8080/api/addAll" mykey1=myvalue1 mykey2=myvalue2
```

### Add a unit test

Use the httptest.NewRequest function to create a request Object and test the API
