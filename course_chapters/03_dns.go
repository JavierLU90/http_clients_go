package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// DNS

// The net/url Package:
// The net/url package is part of Go's standard library. 
// You can instantiate a URL struct using url.Parse:
parsedURL, err := url.Parse("https://homestarrunner.com/toons")
if err != nil {
	fmt.Println("error parsing url:", err)
	return
}
//And then you can extract just the hostname:
parsedURL.Hostname()
// homestarrunner.com

/* What Is the Domain Name System?
So we've talked about domain names, but we haven't talked about the system that makes them work.

DNS, or the "Domain Name System", is the phonebook of the internet. 
Humans type easy-to-read domain names like Boot.dev. 
DNS "resolves" those domain names to their associated IP addresses so that web clients can find the server they're looking for.

Domain names are for humans, IP addresses are for computers.*/

/* How Does DNS Work?
We'll go into more detail on DNS in a future course, but to give you a simplified idea, let's just introduce ICANN. 
ICANN is a not-for-profit organization that manages DNS for the entire internet.

Whenever your computer attempts to resolve a domain name, 
it contacts one of ICANN's "root nameservers" whose address is included in your computer's networking configuration. 
From there, that nameserver can gather the domain records for a specific domain name from their distributed DNS database.

If you think of DNS as a phonebook, ICANN is the publisher that keeps the phonebook in print and available.*/
