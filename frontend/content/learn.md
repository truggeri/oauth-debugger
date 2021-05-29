---
date: 2021-05-23T16:28:24-04:00
---

# Learn

The OAuth2 protocol can be confusing to work with. Let's break down the basic pieces of the flow.

## Authorization

The first part of the OAuth2 flow is authorization. This is where the end user gives permission for the client app (your app) to use some of the resources from the resource server. You've probably seen this step hundreds of times around the web - "App Foo requests access to the following resources."

If the end user grants permission, a `code` is returned to the client.

## Code for a Token

The second step of the process is for the client app to exchange that code, along with their `client_secret`, for an access token. Think of this as exchanging two pieces of authorization - the `code` signifies that the end user has given the client permission and the `client_secret` signifies that the client has established an "account" with the resource server.

The response here is a few things. The most important is a `token`, which is longer lived than a `code`. This will be used in the next step. We also provide a `refresh_token` which can be used to get a new `token` once the current one has expired.

## Protected Resources

TODO

{{< get_started >}}
