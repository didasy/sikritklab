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
  Search thread, defaulted to 10 threads per page descended by `CreatedAt`

```
Queries
page = 0
per_page = 10
tags = "comma,separated,strings"
title = "a string"
```

Returns:

```
{
    "message": [
        {
            "thread": {
                "id": "01C1Q2XF42NX6F64CE22AMDS06",
                "created_at": "2017-12-19T16:50:44.866377215+07:00",
                "title": "test"
            },
            "posts": [
                {
                    "id": 7,
                    "created_at": "2017-12-19T16:50:44.866377215+07:00",
                    "thread_id": "01C1Q2XF42NX6F64CE22AMDS06",
                    "content": "just a test"
                }
            ],
            "tags": [
                {
                    "id": 3,
                    "created_at": "2017-12-19T16:50:44.866377215+07:00",
                    "thread_id": "01C1Q2XF42NX6F64CE22AMDS06",
                    "tag": "test"
                }, ...
            ]
        }
    ]
}
```

* GET /thread/random
  Get random thread, returns:

```
/thread/id/{{thread_id}}
```

* GET /thread/id/:id
  Get thread by id, returns:

```
{
    "message": {
        "thread": {
            "id": "01C1Q2XF42NX6F64CE22AMDS06",
            "created_at": "2017-12-19T16:50:44.866377215+07:00",
            "title": "test"
        },
        "posts": [
            {
                "id": 7,
                "created_at": "2017-12-19T16:50:44.866377215+07:00",
                "thread_id": "01C1Q2XF42NX6F64CE22AMDS06",
                "content": "just a test"
            }, ...
        ],
        "tags": [
            {
                "id": 3,
                "created_at": "2017-12-19T16:50:44.866377215+07:00",
                "thread_id": "01C1Q2XF42NX6F64CE22AMDS06",
                "tag": "test"
            }, ...
        ]
    }
}
```

* POST /thread/new
  Creating new thread

```
JSON
{
    "title": "A string",
    "content": "A longer string",
    "tags": ["many", "tags"],
    "image": "http://placehold.it/100x100"
}
```

Returns:

```
{
    "message": {
        "id": "01C1Q2XF42NX6F64CE22AMDS06",
        "created_at": "2017-12-19T16:50:44.866377215+07:00",
        "title": "test"
    }
}
```

* POST /thread/id/:id
  Replying to a thread.

```
JSON
{
    "title": "A string",
    "content": "A longer string",
    "image": "http://placehold.it/100x100"
}
```

Returns:

```
{
    "message": {
        "id": 9,
        "created_at": "2017-12-19T16:51:33.970133184+07:00",
        "thread_id": "01C1Q2XF42NX6F64CE22AMDS06",
        "content": "HOLY SHIT",
        "image": "http://placehold.it/100x100"
    }
}
```
