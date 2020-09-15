# DistSysMP1

# To Run
To set up each process run
> go run . -s [id]
where [id] is the process id specified in config.txt

To send a message simply type
> send [destination] [message]
where destination is the id of the destination node and [message] is the message you want to send

# Code Decisions and Process
I split my code into four files, main.go, utils.go, server.go, and client.go. Main is for initializing stuff like the server and getting command line commands. Server.go is for server stuff like starting the server and receiving input. This is where unicast_receive lives. Client.go is for client stuff like connecting to the destination server and sending the message. This is where unicast_send lives. Finally, utils.go is for utility functions like reading files, checking errors, and holding structs. Aight let's get to the real code.

# Main
First of all I use argparse to grab the command line argument, which is the process id. The reason I used argparse is because I was required to use it. And because it's better than grabbing the argument normally. I pass the argparsed id to a goroutine of startServer in server.go. This starts a listen server with the correct id port. Meanwhile in main we're loopin around reading command line stuff. If I read that the user typed "send " then I start paying attention and grab where the message is going and what the message is. I call goroutine unicast_send and keep dilligently reading the command line.

# Server
Over here in server.go I just got startServer called which means I gotta boot up a TCP server with the correct port. I find the ID on the config.txt file and set my port with a listen server. When I get someone sending me something I gotta wait a bit to simulate network delay, then I decode the message and print it.

# Client
Im chillin here in client where the only func is unicast_send. When unicast_send gets called I get the destination ip and port using a handy utils.go function. Then I connect and send the message "down the pipeline" as they say.

# Utils
Utils is pretty straightforward. Basically if a function doesn't belong in server.go or client.go I slap it in utils.go and call it a day. More seriously, the functions in utils.go are functions that, if this was a bigger project, I would be using multiple times for various reasons. It's good to compartmentalize these small getters and readers in case the project is expanded on in the future.

# Final Thoughts
I tried my best for this MP but it was hella confusing. I think I got it in the end but I guess that's for you to decide. Let me know how this writeup was cuz it took me a long ass time and idk if you wanted all the info. I think I got points off for not writing enough last time so I wanted to make sure you got enough info about the project and why it is the way it is.
