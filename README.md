sysmb
=========
Simple System Message Broker by UDP.
sysmb supports simple publish a  message and subscribe a message, don't care which destination send to or receive from. 

## Installation
Make sure you have a working go enviroment. 

To install sysmb, run :
	go get github.com/forthe2008/sysmb
	
To compile it from source
	git clone https://github.com/forthe2008/sysmb
	go install && go test
	cd sysmb-server && go install && sysmb-server

## Example
	see https://github.com/forthe2008/sysmb/example
	
	Compile the sysmb-server, run sysmb-server as a daemon in the system. 
	Appilcation use the API privated in sysmb.go to publish/subscribe message
	
	