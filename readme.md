# Simple Mockserver

## Why

To enable frontend developers to applying backend response contract without relying too much to actual backend response.

## Building

Simply run this script

```
./build.sh
```

You can find build result at `build/mockserver`.

## How to Use

Make sure your current workding directory at `build/mockserver`. Take a look on `rules.example.json`, simply copy the example to the current working directory and rename as `rules.json`. Then adjust according to your needs. For example, it will look like this...

```json
[
    {
        "path": "/hello",
        "response": {
            "type": "JSON",
            "value": "./json_sample/hello.json",
            "status_code": 200
        }
    },
    {
        "path": "/foobar",
        "response": {
            "type": "JSON",
            "value": "./json_sample/foobar.json",
            "status_code": 200
        }
    }
]
```

After adding some rules, you can start your server using...

```
./mockserver
```

It will run on default port `8181`, if you desire to use another port, you can execute command above as follows

```
./mockserver --port <desired_port>
```

Every adding the rules you must restart the mock server.