# mocktcpserver
Mocked TCP server for testing using net.Pipes instead of sockets
I needed to do some mock tests for a non-http TCP server without starting a network up.

I found it a bit difficult to find a working example of how to do this. I googled and found pipes mentioned but didn't find a full example so I have put this up here to, hopefully, help others. The server does not function like a real echo server as it serves only one client and only reads and replies once; I did this to avoid overcomplicating the example. I hope this is fairly straightforward.

run like this:
>davids-MacBook-Pro:playground david$ go run t.go 
>client received from server:  Hi client! Server here...
>client said to server: hello Server this is my reply.
