## countdown-api

HTTP server for managing the tasks 

* Server is written in go and starts on port 8080 by default 
* Uses a postgres database for storing all the tasks infomation 
---

* List all the tasks 

```
request : HTTP GET /tasks 

response : 
            [
                {
                    "ID": 5,
                    "CreatedAt": "2020-10-31T07:57:53.509273Z",
                    "UpdatedAt": "2020-10-31T07:57:53.509273Z",
                    "DeletedAt": null,
                    "task": "expired task",
                    "timestamp": 1603527508,
                    "deadline": 0
                },
                {
                    "ID": 6,
                    "CreatedAt": "2020-10-31T07:58:27.994489Z",
                    "UpdatedAt": "2020-10-31T07:58:27.994489Z",
                    "DeletedAt": null,
                    "task": "task 1",
                    "timestamp": 1604217478,
                    "deadline": 14637
                },
                {
                    "ID": 7,
                    "CreatedAt": "2020-10-31T07:58:45.330195Z",
                    "UpdatedAt": "2020-10-31T07:58:45.330195Z",
                    "DeletedAt": null,
                    "task": "task 2",
                    "timestamp": 1606204678,
                    "deadline": 2001837
                },
                {
                    "ID": 8,
                    "CreatedAt": "2020-10-31T07:58:59.575999Z",
                    "UpdatedAt": "2020-10-31T07:58:59.575999Z",
                    "DeletedAt": null,
                    "task": "task 3",
                    "timestamp": 1608883078,
                    "deadline": 4680237
                }
            ]
```

* Create a new task
    - timestamp is the unix timestamp of the deadline 

```
request : HTTP POST /task 

body : 
            {
                "task" : "task 3",
                "timestamp" : 1608883078
            }
        
response : 
            {
                "message" : "task created successfully" 
            }

```

* Delete a task 
    - `task_id` can be retrived from `/tasks`
```
request : HTTP DELETE /task 

body : 
        {
            "task_id": 4
        }

response : 
        {
            "message" : "task deleted successfully"
        }

