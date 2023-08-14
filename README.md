# server_go

get,post,put,delete

steps
******
    1 Import http package
    2 define the port
    3 define the port number
    4 define your routes
    5 create the server instance ,attach the port to the server that we create


get routes
***********
get will take the data via url
any get or post call once reached to the server will be divided into two parts
    • incoming (request)
    • output(response)  in the form of http 

we have to create http handlers

url
****
url is of two parts 
1. search -myntra.com
1. start with ? is called url params



error
******
404-page not found(wrong url)

500- internal server error(correct path but internally server gets crashed)A generic error message, given when no more specific message is suitable.

403-forbidden(route is correct but to access that route there is no permisiion)The request was a legal request, but the server is refusing to respond to it. Unlike a 401 Unauthorized response, authenticating will make no difference.

post and get
**************

post have request body
get and delete doesn't have body


array with structures
*********************

Gin framework
***************
MVC-design pattern
helps to manage project in effective managable manner
m -stands for model
v -stands for view
c -stands for controller

controller acting as a layer to communicate with business


