# DSYSHandin4-Consensus
This program is an implementation of the Ring Solution, for Distributed Mutual Exclusion. 
To run the program, each node should be run on separate terminals. Inside the terminals, execute go run node.go, which is located inside the nodes folder. When the file gets executed, the program will request you to type in a port number into the terminal, which for the first node MUST BE 5050, for the program to run correctly. It will then ask you to insert a target port, the port on which the next node will be set up, which you can freely choose. 
When launching the next node, choose the previous node's target port, as the starting port, then select this nodes target port. Repeat this until you have the wished amount of nodes. There has to be at least 2 nodes, for the system to run.
For the last node, to complete the ring, set the target port to 5050. Now the main program will start and the token will be passed around to every single node. 

Examples of a sequence:

5050 -> 5051, 
5051 -> 5052, 
5053 -> 5050

or

5050 -> 5051, 
5051 -> 5221, 
5221 -> 5057, 
5057 -> 5050

