# sikritklab

A tag based anonymous message board for a free as in freedom free-speech platform

### Dependencies

* github.com/gin-gonic/gin
* github.com/oklog/ulid
* github.com/JesusIslam/lowger
* github.com/asdine/storm

### How to run

1. Check `run.sh` and edit it.
2. Run it `./run.sh`

Or you can just build it and run it, but remember to set all of the environment variables in `run.sh`

### APIs

* GET /thread/search

```
Queries
page = 0
per_page = 10
tags = "comma,separated,strings"
title = "a string"
```

* GET /thread/random

* GET /thread/id/:id

* POST /thread/new

```
JSON
{
    "title": "A string",
    "content": "A longer string",
    "tags": ["many", "tags"]
}
```

* POST /thread/id/:id

```
JSON
{
    "title": "A string",
    "Content": "A longer string"
}
```
