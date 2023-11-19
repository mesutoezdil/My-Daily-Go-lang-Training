//I want to print out a result as a response to a web request. We are going to register a function that's going to be called every time a request comes in on a certain URL. We are going to reach into the HTTP package within the standad library, we are goiung to call a function called HandleFunc. It accepts two parameters. The first parameter is the URL pattern that we are matching against. So I want to actually respond to every request. So my root pattern is going to be the /, and then going to match every request that comes into my app.

//Then I need to provide a func to Go that's going to be invoked every time a request comes in on a path that matches my pattern here. I am going to add an anonymous func here. Go does support function as first class citizens, and so I can create a func right here or any time that I need one. Now notice I've got two arguments that are being passed into this func. --> (w http.ResponseWriter, r *http.Request)... I've got sth called a Response Writer, which is what we use to write responses to our requests, and a Request object. This request object contains all of the info that Go knows about with the incoming request, so the HTTP method, the request body, any cookies that are going to come in are all going to be part of that Request object. But for now, I do not need that. I am just going to write the response back out.

//I am going to reach into that fmt package again, and I am going to use the Fprint func. That allows us to print to standart out. Fprint is a more general function. It allows to write our output to anything that is considered a writer, as you can see right here in the function signuture. Writers are interfaces that describe anything that can be written to. So we could write to standart out, we could write to a file, or we could write to the response of a network request, which is what we're going to do. So we'll pass in that ResponseWriter, that's that variable w, and then the message I want to print out is Web services are easy with Go! 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {fmt.Fprint(w, "Web services are easy with Golang!")})

//When I save that, my tooling will automatically import the package for me, so you see that net/http up here on line 5 and then it's time to start the server, so let'S go ahead and do that. I am goint to start that server by reaching, one again, into that HTTP package and then calling another function,ListenandServe. The first is going to be the port that I want to listen on, and then, we are going to provide a value for the second one. We are going to provide the equivalent of null in Go called nil.

//What this is going to do is tell the standart library that we want it to provide a default handler for our web requests, which is very commonly done. So the standart library comes with a default request handler that's going to handle the routing of our requests to the handlers like the one that we define on this line: http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

//We do not have any objects. That HTTP package is a package. We are calling functions like HandleFunc and ListenAndServe. When we want to print sth, we are calling the Fprint func from the fmt package, no object involved. We do have objects in Go. As a matter of fact, if I look at that Request object and ask for its properties, you can see that we have several properties here, as well as several methods. All of these are available on the Request object. fmt.Fprint(w, "Web services are easy with Golang!")

//So we do have objects with Go, but we are not beholding to everything being a object. If we need a func, we provide a func with Go.

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Web services are easy with Golang!")
	})

	// Let's extend this just a little bit:

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./home.html")
	})

	http.ListenAndServe(":3006", nil)
}

//I've got that home.html file.. I am going to start by adding another HandleFunc bcs I want to register another path that's going to handle requests. We'll call this /home, and then I need that function that's going to be called when a request comes in. Within this implementation, all I want to Go to do is serve up that home.html file. http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {... So once again, within the standart library, I have in the HTTP package, I have a func called ServeFile. And all it needs is the destination, so I need that ResponseWriter, and a request, those are provided to me, so I just pass those along. And then the name of the file, so where is Go going to find this on my file system? That's going to be in my current directory, and it's a file called home.html.-->http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {http.ServeFile(w, r, "./home.html")})

//So whether you are working with static content like this home.html file or dynamically generated content, like we might pass to the Fprint func, Go makes the creation of web services extremely easy, whether you are working with HTML or JSON.
