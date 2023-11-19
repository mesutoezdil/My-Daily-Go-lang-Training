/*
How can we create a command-line app with Go. So to show you that, I've got this log.txt file, which is just a sample of what you might see in a common logging situation where we've got a timestamp, a logging level, and a message. So let's say that we want to create an app  that's going to filter through this log and print out log messages that have a certain level. We are going to start by allowing that log level to be provided to us. We're going to use the command-line parameter here. Notice we do not receive the command-line parameters in the signature for the main function, so we need a way to access those. The way that we're going to do that is we're going to use another part of the standart library called the flag package. This allows strongly typed command-line parameters to be passed into our app as flags. So we need to tell it what type of flag we're going to receive, so we're going to receive a sting, and then we need to provide three parameters. What's the name of the parameter? So this is going to be called level, the default value, so we are going to say, look for CRITICAL, and then a help message, so log level to filter for. So when we run our app, what we're going to do is we're going to be able to look to see what command-line parameters are available, and Go is going to use this info to derive a help document for us.


							package main

							import "flag"

							func main() {
								level := flag.String("level", "CRITICAL", "log level to filter for")
							}


So the next thing that I need to do, you do notice I have some errors here. One error is bcs I do not have the flag package imported, so I can fix that simply by saving the file, and my tooling is going to import that. But notice I've got another error here, and that is that level declared but not used. One of the rules that Go puts in place in order to keep our app as clean as possible and as well as possible is we can't have a local variable that does not have a use. This is actually a compiler error with Go.And so we do have to provide a use for that, but we're not quite ready for that yet. The next thing that I do need to do though is I need to call the parse func from the flag package. This tells Go to look at those command-line parameters and actually populate the variables like we see here on line--> level := flag.String("level", "CRITICAL", "log level to filter for")


						package main

						import "flag"

						func main() {
							level := flag.String("level", "CRITICAL", "log level to filter for")
							flag.Parse()
						}
The next thing that I need to do is I'am going to open up my log, so I am going to do that using Open function from the OS package. And we see another critical pattern starting to develop here with Go. Notice that I get 2 variables back. I get f and err. F is going to be my file handle. That's how I am going to be able to read the contents of my log file. Err is going to be a variable that's going to be populated if sth goes wrong opening that file. Now this is sth else you're going to very commonly see in Go. We promote error handling, and we handle errors first. The intention here is to increase the production stability of our app and the ability to understand how our app responds to errors. Error maanagement is very important in Go, and so we elevate that to a very primary position, and that position is right here--> if err...

						package main

						import "flag"

						func main() {
							level := flag.String("level", "CRITICAL", "log level to filter for")
							flag.Parse()

							f, err := os.Open("./log.txt")
							if err != nil {
								log.Fatal(err)
							}
						}
We are going to check to see if an error was returned. If one was, then for now, just to keep things simple for this demo, I am just going to crash the app using this Fatal func from the log package, printing the error out. Basically what this means is if it fails to open that log file for any reason, we're just going to crash the app out.

And then, what's the next thing that I need to do? Well, the file handle, this f variable on here (f, err := os.Open("./log.txt")), is a resource that I actually requested from the OS. One of the most important things we need to do with those resources is release them in a timely manner so that the OS is able to run as efficiently as possible. In Go, the way that we do that is we have a defer keyword that we can add in front of a function invocation. So we can see here, we want to close that file, but we're going to defer the execution. When do we defer to? We are going to defer the execution if this Close func until after the main func exits. Why do we have this with Go? Bcs if we look at lines "f, err := os.Open("./log.txt")" through "defer f.Close()", we have the attempt to acquire a resource, checking to see if that resource acquisition was successful, managing it if it was not, and then releasing that resource. So we have this really tight construct to make sure that any time we acquire a resource, we are properly managing it. This pattern is commonly used in Go and is one of the reasons why Go programs tend to be stable in production. we elevate error management to be one of the most important things that our program manages.

						package main

						import "flag"

						func main() {
							level := flag.String("level", "CRITICAL", "log level to filter for")
							flag.Parse()

							f, err := os.Open("./log.txt")
							if err != nil {
								log.Fatal(err)
							}
							defer f.Close()
						}

The next thing that I need to do is the file itself is a little bit too primitive for what I want to do. I want to actually be able to iterate through my log file one line at a time. So I am going to wrap this file with a buffered reader. So there is another package in the standard library called bufio, so buffered input/output, and I an going to create a reader that's going to wrap the existing reader. So this is a decarator pattern where we're talking a file and we're adding additional functionally to it by wrapping this buffered reader around it.   

						package main

						import "flag"

						func main() {
							level := flag.String("level", "CRITICAL", "log level to filter for")
							flag.Parse()

							f, err := os.Open("./log.txt")
							if err != nil {
								log.Fatal(err)
							}
							defer f.Close()

							bufReader := bufio.NewReader(f)
						}

Now what I can do is I can loop through the lines of the file. So every loop in Go is a for loop, so I do not have to worry about how am I going to loop through this file? I am just going to start with a for loop. I am going to retrieve an individual line and potentially an error bcs retrieving a line from a file can fail. So we will get that, and then we will ask for the next string from the reader that ends in a new line character. And then as long as the error is nil, we're going to continue. So we are going to actually go through this loop until we end up with an error. Bcs one of the errors that we're going to get is we've run out of lines. So we're just going to keep looping as long as we do not have an error. Actually, I am going to add a post clause as well. It is going to be the same as the initializer. What it's going to do is this first clause is going to initialize our loop. It's going to be our first line read. We're then going to have our test, making sure that we do not have an error (err == nil). And then every time we go through the loop, we're going to execute that initializer statement again to pull the next time "line, err = bufReader.ReadString('\n')".

						package main

						import "flag"

						func main() {
							level := flag.String("level", "CRITICAL", "log level to filter for")
							flag.Parse()

							f, err := os.Open("./log.txt")
							if err != nil {
								log.Fatal(err)
							}
							defer f.Close()

							bufReader := bufio.NewReader(f)

							for line, err := bufReader.ReadString('\n'); err == nil; line, err = bufReader.ReadString('\n') {

							}
						}

And then we'll check to see if the current the line contains that log level, and if it does, then we'll print it. This contains func comes from the standard library yet again. It comes from the strings package. 

						package main

						import "flag"

						func main() {
							level := flag.String("level", "CRITICAL", "log level to filter for")
							flag.Parse()

							f, err := os.Open("./log.txt")
							if err != nil {
								log.Fatal(err)
							}
							defer f.Close()

							bufReader := bufio.NewReader(f)

							for line, err := bufReader.ReadString('\n'); err == nil; line, err = bufReader.ReadString('\n') {
								if strings.Contains(line, *level) {
									fmt.Println(line)
								}
							}
						}

And then when I run this program, you'll see since I've got a default log level of CRITICAL, I should see these two lines right here: 

						1970-01-01 00:00:00 CRITICAL gRPC client failed to connect to server
						1970-01-01 00:00:00 CRITICAL Failed to find user ID '42'
			
Let's go ahead and run that to see if that happens. I do see those 2 lines printed out. If I come back here, maybe I asked for the debug information, so I'l try and get this line out here: "1970-01-01 00:00:00 DEBUG This is an debug message", it will provide that flag level, and we'll say this is going to be DEBUG (i.e. on the terminal: go run . -level DEBUG) and now I get that one debug message printed out. 

Now one last thing, I do want to talk about these asteriks right here like we see: (line, *level) {... This is a pointer. So if you've worked with a language that uses pointers in the past, then that's exactly what we're looking at here. Talking about when to use pointers and when not is a little bit more advanced than what we're going to talk about here.




						

