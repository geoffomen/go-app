curl -L -X GET -v 'localhost:8785/exam/hello'

curl -L -X POST -v 'localhost:8785/exam/echo' -H 'Content-Type: application/json' --data-raw '{
    "strVal": "string",
    "intVal": 100,
    "intPtrVal": 1,
    "structVal": {
        "id": 2
    },
    "sliceVal": [
        2,
        3
    ]
}'