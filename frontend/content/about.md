---
date: 2021-05-29T08:28:24-04:00
---

# About

This site is made for web developers that are trying to work with OAuth2. Born out of frustration of trying
to test client code, this site aims to provide a "dummy" setup that can be used to debug your code _before_
going live.

## How It Works

The aim is for simplicity. First, create a client from our setup page. Once you have your credentials,
place them in your application (the client), and attempt to use the OAuth2 flow. If your code is working,
you'll see our login page which offers three different users. We do this so now real account information
is exchanged in the process.

Once you log in, the process should complete passing you back a token allowing for you to get information about
your selected user from our API - just like from a real resource server. If anything goes wrong along the way,
helpful error messages will inform you of what expectation was not met from the server helping you to debug
your code.

## Technology

The underlying code is written using [Go](https://golang.org/) functions running on
[Google Cloud Platform](https://cloud.google.com/functions). We store your client
information and tokens using [Firestore](https://firebase.google.com/products/firestore),
a NoSQL cloud database. The front-end, which you are most likely
looking at now, is hosted using [Firebase](https://firebase.google.com/) hosting and written with a
combination of [Hugo](https://gohugo.io/) and [Svelte](https://svelte.dev/).

All of the code is available under the MIT license and hosted on Github.

## Support

If you have any questions or think you've found a bug, please file an issue on our Github page.
We also have an [api reference page](/docs) to help.
Please remember that this is not a sponsored project, but I will do my best to make this a
productive application for everyone.
