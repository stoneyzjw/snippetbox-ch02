# snippetbox

## Chapter02 

Use go mod init command to create the module: github.com/stoneyzjw/snippetbox. the file go.mod is the
result.

## Web application basics 

Now that everything is set up correctly let's make the first iteration of our web application. We'll
begin with the three absolute essentials: 
1. The first thing we need is a handler. If you're coming from an MVC-background, you can think of
   handlers as being a bit like controllers. They're responsible for excuting your application logic
   and for writing HTTP response headers and bodies.

2. The second component is a router 9or servermux in Go terminology). This stores a mapping between the
   URL patterns for your application and the corresponding handlers. Usually you have one servemux for
   your application containing all your routes. 

3. The last thing we need is a web server. One of the great things about Go is that you can establish a
   web server and listen for incoming request as part of your application itself. You don't need an
   external third-party server like Nginx or Apache. 

Let's put these components together in the main.go file to make a working application. 
