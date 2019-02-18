Simple application that returns the url path in a json.

To run:

    $ docker build -t echo .
    $ docker run --name echo -d -p 8888:8888 echo
    $ curl -s http://localhost:8888/hello | jq
    {
      "Message": "hello"
    }
