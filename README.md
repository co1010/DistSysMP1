# DistSysMP1

# To Run
To set up each process run
> go run . -s [id]

where [id] is the process id specified in config.txt

To send a message simply type
> send [destination] [message]

where destination is the id of the destination node and [message] is the message you want to send

# Code Decisions and Process
I split my code into four files, main.go, utils.go, server.go, and client.go. Main is for calling functions in server.go and client.go as well as getting command line commands. Server.go starts the server and receives incoming messages. This is where unicast_receive lives. Client.go connects to the destination server and sends messages. This is where unicast_send lives. Finally, utils.go is for utility functions like reading files, checking errors, and holding structs. Utils.go is compartmentalized into it's own directory and subpackage because it is used by all other files. 

# High-level Description
I use argparse to grab the command line argument, which is the process id. I used argparse is because I was required to use it and because it's better than grabbing the argument normally. I pass the argparsed id to a goroutine of startServer in server.go. This starts a listen server with the correct id port. Meanwhile in main there is a loop reading the command line. If the user uses the send command, the reader grabs the message destination and message content. Unicast_send is called as a goroutine to avoid a bottleneck.

In server.go startServer() boots up a TCP server with the correct port. The config file is referenced to find the port for the listen server. The server then accepts the client's connection. To simulate network delay without causing a bottleneck, the decodeMessage function is called using a goroutine. This function sleeps the goroutine for a random time within the parameters then decodes and calls unicast_receive which prints the message.

In client.go the only function is unicast_send. When unicast_send gets called the destination ip and port are grabbed using the utils.go function GetNodeDetails. The ip and port are used to connect to the TCP server. If there's a TCP server available the message is encoded using gob and sent to the server.

The functions in utils.go are functions that, if this was a bigger project, I would be using multiple times for various reasons. It's good to compartmentalize these small getters and readers in case the project is expanded on in the future. utils.go is compartmentalized into it's own subpackage because it's used in lots of files.

# Delay Implementation
Delay is implemented as to not cause a bottleneck on a server that might receive multiple messages in a row. For every accepted connection, a goroutine is called which implements delay within that function. This way, if a server receives multiple messages at once, a goroutine with separate delays will be called for each connection. This eliminates the bottleneck while perserving the delay.

# Message Struct
Messages are sent using Message structs which contain only a content field (string). Messages are formatted as structs because gob cannot encode strings and structs are an easy way to send data through a TCP connection using gob encoding. One limitation in the future is if messages can't be contained within a struct due to extra fields or properties of those messages.

# Limitations and Shortcomings
No code is perfect, and every implementation has some sort of flaw. One drawback to my implementation is that only 1 message can be sent to a server before it closes. This is due to defer ln.Close() on line 25 in server.go. This closes the connection after receiving a single message. One possible fix would be to put line 26-30 in an infinite for loop. This brings up a different problem which is then that the server never closes. An implementation using channels to eventually close the server when given a command is probably possible but out of the scope of this project.
