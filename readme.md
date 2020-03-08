# Simple Mockserver

## Why

To enable frontend developers to applying backend response contract without relying too much to actual backend response.

## How to Use

Take a look on `rules.example.json`, simply copy the example to `rules.json` then adjust according to your needs. For example, it will look like this...

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

Every adding the rules you must restarting the mock server.

## Building

Simply run this script

```
./build.sh
```

Then enter directory `build/mockserver`, add rule file on it then run `./mockserver`