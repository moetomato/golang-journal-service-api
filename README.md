# Journal site APIs :bouquet:

## Overview :cake:
Simple APIs for posting and retrieving journal articles, and posting comments on them.

## Features :peach: 
- retrieve a journal post by ID 
- retrieve journal lists (limit 5,  requires offset)
- post journal posts
- like a journal
- post comments on a journal

## Authorization :doughnut:
OIDC (OAuth2.0)

To use the APIs, it is necessary to first obtain the user's authorization and then obtain an access token by logging in
[here](https://accounts.google.com/o/oauth2/v2/auth/oauthchooseaccount?client_id=315275688731-ctvfikjeqikaerjngmcsqu54s0frk4bk.apps.googleusercontent.com&response_type=id_token&scope=openid%20profile&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fcallback&state=xyz&nonce=abc).

Note that the access token issued will be required each time you access the API. It will be invalid after a certain period of time. 


## Endpoints :cherries:
|                                                | メソッド        | URI              | 
| ---------------------------------------------- | ---------------| ---------------- |
| retrieve a specific journal with comments      | GET            | /journal/{id:.*} |
| retrieve journal lists                         | GET            | /journal/list    |
| post a journal                                 | POST           | /journal         |
| like a journal                                 | POST           | /journal/nice    |
| comment on a journal                           | POST           | /comment         |

## Requirements :cupcake:
Paramiters required are the followings : 

- path
- queries
- headers
- request body (for POST) 

The parameters in the request body are accepted in JSON format.


## Response :custard:
On success, a 200 status code is returned, and in case of an error, a 4xx or 5xx status code (where x is a digit) is returned.


## Authour :parachute:
moetomato


