# Exercise 8 - Network Programming

If you do not finish during the lecture period, please finish it as homework.

## Exercise 8.1 - TCP connection

Make a TCP connection with *simulationcorner.net* on port *54321*, read from the data stream and write the output onto the screen in an infinite loop.

## Exercise 8.2 - Remote procedure call

Implement in Go a net/rpc client TCP connection via rpc.Dial against the URL *simulationcorner.net:1234* . 
This interface offers two remote procedures calls.

A function, in which you receive a rune

    var dummy struct {}
    var output rune
    err := client.Call("Session.Read", &dummy, &output)

and a function in which you have to provide a rune

    var input rune
    err := client.Call("Session.Write", &input, &dummy)

For each function create its own goroutine and call these functions in an infinite loop. Write the output of *Session.Read* to the terminal. As input for *Session.Write* use os.Stdin to read a rune from keyboard via

    reader := bufio.NewReader(os.Stdin)
    input, _, err := reader.ReadRune()

## Exercise 8.3 - Call Rest API

Call the Github API to receive the repositories with the most stars and the query "awesome"

https://api.github.com/search/repositories?sort=stars&order=desc&q=awesome

Use https://mholt.github.io/json-to-go/ to instantly create a Go structure from an arbitrary JSON.

## Exercise 7.4 - Microservice with API and embedded Web Site

Create a REST API for a key value store.

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

###  Include static web site

Add an index.html file, which can interface the key value store. E. g. following HTML File shows the data from the key value store.

```HTML
<!DOCTYPE html>
<html>
<body>

<h3>Key Value List</h3>
<pre id="store"></pre>

<script>
    async function ListStore() {
        let response = await fetch("/api/list")
        let store = await response.json()
        let pre = document.getElementById("store")
        for (let key in store) {
            pre.innerText += key + ": " + store[key] + "\n"
        }
    }
    ListStore()
</script>

</body>
</html>
```

Add a http.Fileserver to your REST service to access this file.

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
