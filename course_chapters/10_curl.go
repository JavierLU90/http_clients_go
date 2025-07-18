package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// cURL:

/* cURL is a command-line tool for transferring data using various protocols, like HTTP and HTTPS. 
It’s a favorite among developers. 
Among other things, developers use cURL for:
	Quick testing: make HTTP requests with a single command.
	Automation: can seamlessly be integrated into scripts and automation workflows.
	Debugging: easily see requests and responses.*/

/* cURL Basics
The simplest cURL command makes a GET request to the provided URL, for example:

	curl https://jsonplaceholder.typicode.com/users/1

The response body is written to stdout.
Because the output is written to stdout you can combine it with other commands, e.g. redirecting the output to a file.

	curl https://jsonplaceholder.typicode.com/users/1 > user1.json

cURL supports a number of options and commands. 
Check their tutorial page or the man page for more information. We will only be covering the basics.

	man curl
*/

/* cURL POST Request
Making a POST request with cURL is almost as simple as a GET request. 
Use the -X POST option to specify the request method and the -d option to pass data.

Here's how to make a POST request:
	curl -X POST http://example.com/resource -d "param1=value1&param2=value2"

This command makes an HTTP POST request to http://example.com/resource with the data param1=value1&param2=value2. 
The server's response body is written to stdout.

If you need to send JSON data, use the -H option to set the Content-Type header 
and the -d option to pass the JSON data:
	curl -X POST http://example.com/resource -H "Content-Type: application/json" -d '{"key1":"value1","key2":"value2"}'
*/

/* jq is a powerful command-line tool for processing JSON data. 
It’s a favorite among developers for working with JSON because it can:

	Parse JSON: easily read and extract data from JSON responses.
	Manipulate JSON: modify JSON data on the fly.
	Filter JSON: find exactly what you're looking for within large JSON structures.
*/

/* jq Basics
Once installed, you can use jq to parse and manipulate JSON data. 
Here's a simple example to get you started. Suppose you have a JSON file named user.json:*/
{
  "name": "John",
  "age": 30,
  "city": "New York"
}
/* To extract the name field, you would use:
	jq '.name' user.json
	# "John"

jq with cURL:
jq is frequently used with cURL to parse JSON responses directly from HTTP requests. 
For example, to fetch JSON data from an API and extract a field we can use the object identifier index:
	curl https://jsonplaceholder.typicode.com/users/1 | jq .username
	# "Bret"

To get a field from each element in an array you can use the array/object value iterator .[], 
which can in turn be combined with the identifier index:
	curl https://jsonplaceholder.typicode.com/users | jq '.[].username'
	# "Bret"
	# "Antonette"
	# "Samantha"
	# "Karianne"
	# "Kamren"
	# "Leopoldo_Corkery"
	# "Elwyn.Skiles"
	# "Maxime_Nienow"
	# "Delphine"
	# "Moriah.Stanton"

Multiple Fields:
	curl https://jsonplaceholder.typicode.com/users/1 | jq '.name, .email'
	# "Leanne Graham"
	# "Sincere@april.biz"

jq is extremely powerful. We won't cover every feature in this course, 
but it's a tool you should know about for the future.*/
