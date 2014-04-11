jsonpify
========

A go service to provide a jsonp representation of a json service endpoints data

Only supports get requests (since that is all jsonp supports). 

Accepts a url for an endpoint, and a callback for wrapping the content in. 

## Example (jquery)

```
$.getScript('http://jsonpify.service.tld?callback=test&url=http://echo.jsontest.com/key/value/one/two')
```

Will call a method called test with the data from the http://echo.jsontest.com/key/value/one/two service. 

```
test({
"one": "two",
"key": "value"
})
```

