rss-to-kindle
===

This is a utility for converting rss into .mobi periodical, and sending them to your kindle email.

Currently the script is kind of rough, but it will get the job done (even images!), for instance:
* no support for multiple section yet (everything is under the "Main" section)
* minor formatting issue (padding, margin, etc ...)

## Requirement

The only dependency is a golang compiler and kindlegen. Please install [golang](https://golang.org/dl/) and [kindlegen](https://www.amazon.com/gp/feature.html?docId=1000765211) first and install it in your $PATH.

## Usage

To use the script, create a `.env` file in the same path as the script, containing:
```
SERVER=mail.gmx.com
PORT=587
FROM_EMAIL=<your-email>@gmx.com
PASSWORD=<you-email-password>
TO_EMAIL=<your-kindle-email>@kindle.com
LINKS=http://feeds.initium.news/theinitium?format=xml,https://www.thestandnews.com/rss/,http://arstechnica.com/feed/
```

* SERVER: URL of your email server ([gmx.com](http://gmx.com/) is a nice choice)
* PORT: Port of your email server
* FROM_EMAIL: Your email address (will be used to send the email)
* PASSWORD: Password of your email account
* TO_EMAIL: Your kindle email address
* LINKS: links of all the rss urls (seperated by commas)

To kickstart everything, run:
```sh
go run main.go
```