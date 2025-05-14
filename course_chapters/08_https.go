package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// HTTPS:
/* Hypertext Transfer Protocol Secure or HTTPS is an extension of the HTTP protocol. 
HTTPS secures the data transfer between client and server by encrypting all of the communication.

HTTPS allows a client to safely share sensitive information with the server through an HTTP request, 
such as credit card information, passwords, or bank account numbers.*/

/* Security and Encryption:
HTTPS requires that the client use SSL or TLS to protect requests and traffic by encrypting the information in the request. 
You don’t need to manually handle encryption or decryption—it's all done for you by the protocol. 
HTTPS is simply HTTP with built-in security!

HTTPS != Privacy
We won't cover how encryption works in this course, 
but it's important to understand what it does and doesn't do, at least in regards to HTTPS.
	HTTPS will encrypt the data in a request (like passwords, headers, body information, query parameters, etc). 
		You and the recipient are the only ones who can read that information.
	HTTPS will not hide who you are or that you're communicating with the given server. 
		If you don't want your wife who knows how to sniff network traffic to know you're using Boot.dev, 
		HTTPS won't help.

HTTPS Verifies the Server:
In addition to encrypting the information within a request, 
HTTPS uses digital signatures to prove that you're communicating with the server that you think you are. 
If a hacker were to intercept an HTTPS request by tapping into a network cable, 
they wouldn't be able to successfully pretend they are your bank's web server.*/
