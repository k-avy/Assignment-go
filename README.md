# Assignment-Go

It is a command line tool that solves the prefix and postfix expression, it also converts the expression into desired form.

## Prerequisite

Please [install and set-up Golang](https://go.dev/doc/install) on your system in advance.

## How to run this project?

1. Clone this Project and Navigate to the folder.

```bash
https://github.com/k-avy/Assignment-go.git
cd Assignment
```

2. Build the project using following command.

```bash
go build ./cmd/assign
```

3. Run the executable in your vscode terminal.

```bash
./assign
```

4. You can directly run it by the following command.

```bash 
go run ./cmd/assign
```


## Features

1. You can Create and Resolve a Dispute.


```bash 
# to create Dispute
./assign --username admin --password secret create --id d123 --txn t456 --merchant m789

# for Resolve
touch evidence.txt

./assign --username admin --password secret resolve --id d123 --evidence evidence.txt 
```

2. A json file will be formed to store all the despute and resolved by the agent and other schema.

3. Used mutex to update once resolved.

4. Working on the Dashboard to create the better understanding and analysis.
```bash

# to see the dashboard
 ./assign --username admin --password secret dashboard  
 ```
 

