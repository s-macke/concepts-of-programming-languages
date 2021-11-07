# Exercise 7 - Distributed Computing

If you do not finish during the lecture period, please finish it as homework.

## Exercise 7.1 - Microservice with API and embedded Web Site

The goal of this exercise is to create a simple REST API for a key value store

```Go
    map[string]string
```

### Add an "add" entry API call

```
/api/add?key=mykey&value=myvalue
```

To test the rest interface you can use tools such as
- curl
- httpie
- Postman
- Your own Go Rest call code

###  List entry API call

Add a Rest API GET call /api/list call which returns the complete key value store as JSON list

###  Include web site

Add web site as an embedded file system, which can interface the key value store

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
