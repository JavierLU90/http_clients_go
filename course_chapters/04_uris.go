package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// URIs

/* Uniform Resource Identifiers
A URI, or Uniform Resource Identifier, 
is a unique character sequence that identifies a resource that is (almost always) accessed via the internet.

Just like Golang has syntax rules, so do URIs. 
These rules help ensure uniformity so that any program can interpret the meaning of the URI in the same way.

URIs come in two main types:
	URLs
	URNs

URLs have quite a few sections. Some are required, some are not.
*/

func newParsedURL(urlString string) ParsedURL {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return ParsedURL{}
	}

	password := ""
	if pw, hasPassword := parsedUrl.User.Password(); hasPassword {
		password = pw
	}

	return ParsedURL{
		protocol: parsedUrl.Scheme,
		username: parsedUrl.User.Username(),
		password: password,
		hostname: parsedUrl.Hostname(),
		port:     parsedUrl.Port(),
		pathname: parsedUrl.Path,
		search:   parsedUrl.RawQuery,
		hash:     parsedUrl.Fragment,
	}
}

/* There are 8 main parts to a URL, though not all the sections are always present. 
Each piece plays a role in helping clients locate the droids resources they're looking for.

PART		REQUIRED
Protocol	Yes
Username	No
Password	No
Domain		Yes
Port		No (defaults to 80 or 443)
Path		No (defaults to /)
Query		No
Fragment	No
*/

/* The Protocol
The "protocol" (also referred to as the "scheme") is the first component of a URL. 
It defines the rules by which the data being communicated is displayed, encoded or formatted.

Some examples of different URL protocols:
	http
	ftp
	mailto
	https
For example:

Not All Schemes Require a “//”
The "http" in a URL is always followed by ://. 
All URLs have the colon, but the // part is only included for schemes that have an authority component. 
As you can see above, the mailto scheme doesn't use an authority component, so it doesn't need the slashes.
*/
