package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// JSON:

/* JSON Syntax
JSON (JavaScript Object Notation), is a standard for representing structured data based on JavaScript's object syntax. 
It is commonly used to transmit data in web apps via HTTP. 

JSON supports the following primitive data types:
	Strings, e.g. "Hello, World!"
	Numbers, e.g. 42 or 3.14
	Booleans, e.g. true
	Null, e.g. null
And the following collection types:
	Arrays, e.g. [1, 2, 3]
	Object literals, e.g. {"key": "value"}

JSON is similar to JavaScript objects and Python dictionaries. 
Keys are always strings, and the values can be any data type, including other objects.

The following is valid JSON data:
*/
{
    "movies": [
        {
            "id": 1,
            "title": "Iron Man",
            "director": "Jon Favreau",
            "favorite": true
        },
        {
            "id": 2,
            "title": "The Avengers",
            "director": "Joss Whedon",
            "favorite": false
        }
    ]
}

/* Decoding JSON:
When we receive JSON data in the body of an HTTP response, it comes as a stream of bytes. 
As we saw before, we can just convert the bytes to a string... but in Go there's a better way. 
It's typically best to decode the JSON data into a struct.

Take this example JSON data:
*/
[
  {
    "id": "001-a",
    "title": "Unspaghettify code",
    "estimate": 9001
  }
]

// To decode this JSON into a slice of Issue structs, 
// we need to know the JSON fields and their types. 
// The standard encoding/json package uses tags to map JSON fields to struct fields.
type Issue struct { //Struct fields must be exported (capitalized) to decode JSON.
	Id       string `json:"id"`
	Title    string `json:"title"`
	Estimate int    `json:"estimate"`
}

// After receiving the response, 
// we can decode it into a slice of Issues with the "address of" operator &:
var issues []Issue
decoder := json.NewDecoder(res.Body) // res is a successful `http.Response`
if err := decoder.Decode(&issues); err != nil {
    fmt.Println("error decoding response body")
    return
}

// If no error occurs, we can use the slice of items in our program.
for _, issue := range issues {
    fmt.Printf("Issue - id: %v, title: %v, estimate: %v\n", issue.Id, issue.Title, issue.Estimate)
    // Issue – id: 001-a, title: Unspaghettify code, estimate: 9001
}


// Unmarshal JSON:
/* We can decode JSON bytes (or strings) into a Go struct using json.Unmarshal or a json.Decoder.

The Decode method of json.Decoder streams data from an io.Reader into a Go struct, 
while json.Unmarshal works with data that's already in []byte format. 
Using a json.Decoder can be more memory-efficient because it doesn't load all the data into memory at once. 
json.Unmarshal is ideal for small JSON data you already have in memory. 
When dealing with HTTP requests and responses, 
you will likely use json.Decoder since it works directly with an io.Reader.*/

// res is an http.Response
defer res.Body.Close()

data, err := io.ReadAll(res.Body)
if err != nil {
	return nil, err
}

var issues []Issue
if err = json.Unmarshal(data, &issues); err != nil {
    return nil, err
}


// Marshal JSON:
// If there is a way to unmarshal JSON data, there must be a way to marshal it as well. 
// The json.Marshal function converts a Go struct into a slice of bytes representing JSON data.
type Board struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	TeamId   int    `json:"team"`
	TeamName string `json:"team_name"`
}

board := Board{
	Id:       1,
	Name:     "API",
	TeamId:   9001,
	TeamName: "Backend",
}

data, err := json.Marshal(board)
if err != nil {
	log.Fatal(err)
}
fmt.Println(string(data))
// {"id":1,"name":"API","team":9001,"team_name":"Backend"}


// XML
/* We can't talk about JSON without mentioning XML. 
XML, or "Extensible Markup Language" is a text-based format for representing structured information, 
like JSON - but different (and a bit more verbose).

XML Syntax:
XML is a markup language like HTML, but it's more generalized in that it does not use predefined tags. 
Just like how in a JSON object keys can be called anything, XML tags can also have any name.

XML representing a movie: 
<root>
  <id>1</id>
  <genre>Action</genre>
  <title>Iron Man</title>
  <director>Jon Favreau</director>
</root>

The same data in JSON: */
{
  "id": "1",
  "genre": "Action",
  "title": "Iron Man",
  "director": "Jon Favreau"
}


// map[string]interface{}
// Sometimes you have to deal with JSON data of unknown or varying structure in Go. 
// In those instances map[string]interface{} offers a flexible way to handle it without predefined structs.

// Think about it: a JSON object is a key-value pair, 
// where the key is a string and the value can be any JSON type. 
// map[string]interface{} is a map where the key is a string and the value can be any Go type.

var data map[string]interface{}
jsonString := `{"name": "Alice", "age": 30, "address": {"city": "Wonderland"}}`
json.Unmarshal([]byte(jsonString), &data)
fmt.Println(data["name"])  // Output: Alice
fmt.Println(data["address"].(map[string]interface{})["city"])  // Output: Wonderland

// any is an alias for interface{}
