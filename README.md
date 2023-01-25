# gRPC_Example
This is gRPC application using Golang

# Prerequisites 

1. Go Runtime: any three latest major releases. If you haven't already installed Go, you can manually install go.<br /> 

2. Protocol Buffer compiler, protoc, version 3. Refer protoc-installation for more info. <br />

3. Go plugins for the protocol buffer compiler. <br />

Install plugins by issuing the below commands in the terminal <br />

`` $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`` <br />


# How to Run this Programm
Step 1: First Clone the repository and go to the grpc_example folder <br />
Step 2: Run `` $ go mod tidy `` <br />
Step 3: `` $ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative practice.proto `` <br />
    After Run this command two file are creating <br />
    ### a) practice_grpc.pb.go ( file contains code for populating, serializing and retrieving NewExample and Example message types)<br />
    ### b) practice.pb.go (contains client and server code)<br />
 Step 4: Run the client and server main programme sepearte terminals<br />
    like : In Server Folders - `` $ go run main.go ``<br />
           In Client Folders - `` $ go run main.go ``<br />
           
    
