package main

// WHY HTTP???

/* 
HTTP Requests and Responses:

At the heart of HTTP is a simple request-response system. 
The "requesting" computer, also known as the "client", asks another computer for some information. 
That computer, "the server" sends back a response with the information that was requested.

	Request: "What issues are on Jello?"
	Response: ["Fix bug", "Improve auth flow"]


HTTP Powers Websites:

HTTP, or Hypertext Transfer Protocol, is a protocol designed to transfer information between computers.
There are other protocols for communicating over the internet, 
but HTTP is the most popular and is particularly great for websites and web applications. 
Each time you visit a website, your browser is making an HTTP request to that website's server. 
The server responds with all the text, images, 
and styling information that your browser needs to render its pretty website!


HTTP URLs:

A URL, or Uniform Resource Locator, is the address of another computer, or "server" on the internet. 
Part of the URL specifies where to reach the server, 
and part of it tells the server what information we want.

Put simply, a URL represents a piece of information on some computer somewhere. 
We can get access to it by making a request, and reading the response that the server replies with.


Using URLs in HTTP:

The http:// at the beginning of a website URL specifies that the http protocol 
will be used for communication.

Other communication protocols use URLs as well, (hence "Uniform Resource Locator"). 
That's why we need to be specific when we're making HTTP requests by prefixing the URL with http://
*/


// Making a Request:
import (
	"fmt"
	"io"
	"net/http"
)

func getProjects() ([]byte, error) {
	res, err := http.Get("https://api.jello.com/projects")
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	return data, nil
}
/* 
- 	http.Get uses the http.DefaultClient to make a request to the given url
- 	res is the HTTP response that comes back from the server
- 	defer res.Body.Close() ensures that the response body is properly closed after reading. 
	Not doing so can cause memory issues.
- 	io.ReadAll reads the response body into a slice of bytes []byte called data */

/*
Web Servers:

Up to this point, most of the data you have worked with in your code has simply 
been generated and stored locally in variables.

While you'll always use variables to store and manipulate data while your program is running, 
most websites and apps use a web server to store, sort, 
and serve that data so that it sticks around for longer than a single session, 
and can be accessed by multiple devices.

Listening and Serving Data:

Similar to how a server at a restaurant brings your food to the table, 
a web server serves web resources, such as web pages, images, and other data. 
The server is turned on and "listening" for inbound requests constantly 
so that the second it receives a new request, it can send an appropriate response.

The Server Is the Back-End:

While the "front-end" of a website or web application is the device the user interacts with, 
the "back-end" is the server that keeps all the data housed in a central location.

A Server Is Just a Computer:

"Server" is just the name we give to a computer that is taking on the role 
of serving data across a network connection. 
A good server is turned on and available 24 hours a day, 7 days a week. 
While your laptop can be used as a server, 
it makes more sense to use a computer in a data center that's designed to be up and running constantly.*/