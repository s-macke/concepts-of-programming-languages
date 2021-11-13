# Exercise 7 - Distributed Computing

If you do not finish during the lecture period, please finish it as homework.

## Exercise 7.1 - Microservice with API and embedded Web Site

Call the Github API with Go to receive the repositories with the most stars and the query "awesome"

https://api.github.com/search/repositories?sort=stars&order=desc&q=awesome

Use https://mholt.github.io/json-to-go/ to instantly create a Go structure from an arbitrary JSON.




## Exercise 7.2 - Microservice with API and embedded Web Site

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
