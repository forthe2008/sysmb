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