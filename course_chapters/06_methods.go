package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Methods:

/* HTTP defines a set of methods. 
We must choose one to use each time we make an HTTP request. 
The most common ones include:
	GET
	POST
	PUT
	DELETE

The GET method is used to "get" a representation of a specified resource. 
It doesn't take (remove) the data from the server but rather gets a representation, or copy, of the resource in its current state. 
A GET request is considered a safe method to call multiple times because it shouldn't alter the state of the server.


Making a GET Request in Go:
There are two basic ways to make a Get request in Go.
	The simple but less powerful way: http.Get
	The verbose but more powerful way: http.Client, http.NewRequest, and http.Client.Do
*/
// If all you need to do is make a simple GET request to a URL, http.Get will work:
resp, err := http.Get("https://jsonplaceholder.typicode.com/users")

// If you need to customize things like headers, cookies, or timeouts, 
// you'll want to create a custom http.Client, and http.NewRequest, 
// then use the client's Do method to execute it.
client := &http.Client{
	Timeout: time.Second * 10,
}

req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/users", nil)
if err != nil {
	log.Fatal(err)
}

resp, err := client.Do(req)

/* POST Requests:
An HTTP POST request sends data to a server, typically to create a new resource.

Adding a Body:
The body of the request is the payload sent to the server. 
The special Content-Type header is used to tell the server the format of the body: application/json for JSON data in our case. 
POST requests are generally not safe methods to call multiple times because that would create duplicate records. 
For example, you wouldn't want to accidentally send a tweet twice.

Like http.Get, the standard library's http.Post function can be used to send simple POST requests. 
The trouble is, it's a bit limited. 
And because we need to add an X-API-Key header, the simple http.Post function won't work for us. 
Instead, we need to use http.NewRequest: 
*/
type Comment struct {
	Id      string `json:"id"`
	UserId  string `json:"user_id"`
	Comment string `json:"comment"`
}

func createComment(url, apiKey string, commentStruct Comment) (Comment, error) {
    // encode our comment as json
	jsonData, err := json.Marshal(commentStruct)
	if err != nil {
		return Comment{}, err
	}

    // create a new request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return Comment{}, err
	}

    // set request headers
	req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-API-Key", apiKey)

    // create a new client and make the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return Comment{}, err
	}
	defer res.Body.Close()

    // decode the json data from the response
	// into a new Comment struct
	var comment Comment
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&comment)
	if err != nil {
		return Comment{}, err
	}

	return comment, nil
}

// Status Code Property
// The http.Response struct has a .StatusCode property that contains the status code of the response.


// Delete:
// The DELETE method does exactly what you expect: it deletes a specified resource.
// This deletes the location with ID: 52fdfc07-2182-454f-963f-5f0f9a621d72
url := "https://api.boot.dev/v1/courses_rest_api/learn-http/locations/52fdfc07-2182-454f-963f-5f0f9a621d72"
req, err := http.NewRequest("DELETE", url, nil)
if err != nil {
	fmt.Println(err)
    return
}

client := &http.Client{}
res, err := client.Do(req)
if err != nil {
	fmt.Println(err)
    return
}
defer res.Body.Close()

if res.StatusCode > 299 {
    fmt.Println("request to delete location unsuccessful")
    return
}
fmt.Println("request to delete location successful")
