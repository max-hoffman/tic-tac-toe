# tic-tac-toe
This is a command-line game of basic tic-tac-toe, written in go.

# Use w/o docker
1. Setup system to work with golang: https://golang.org/doc/code.html#GOPATH
2. Clone the repo
3. Navigate to the proper directory
4. ```go install``` to add to src (link above describes other ways)
5. Either run ```tic-tac-toe``` using the $PATH/Go literal, or from current
directory: ```./tic-tac-toe```

# Use w/ Docker
1. Build locally (I did not push to my registry)
```
docker build -t tic-tac-toe .
```
** do not miss the "." at the end, indiciating the local ```Dockerfile```

2. Run
```
docker run -it -p 8080:8080 tic-tac-toe
```
** you'll need interactive mode to use this app with the command line