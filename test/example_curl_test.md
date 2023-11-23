curl -L -i -X GET 'http://localhost:8000/example/api/v1/echoargs/echo_query?id=1&f32=3.2&f64=6.4&email=email&si=100&si=1000&sf32=3.21&sf32=3.22&sf64=6.41&sf64=6.42&ss=str1&ss=str2&tm=2009-02-14T07:31:30.123Z&v=1000000'

curl -L -i -X POST 'http://localhost:8000/example/api/v1/echoargs/echo_json' -H 'Content-Type: applicatiton/json' -d '{"id": 1, "f32": 3.2, "f64": 6.4, "email": "email", "si": [100, 1000], "sf32": [3.21, 3.22], "sf64": [6.41, 6.42], "ss": ["str1", "str2"], "tm": "2009-02-14T07:31:30.123Z"}'

curl -L -i -X POST 'http://localhost:8000/example/api/v1/echoargs/echo_form' -H 'Content-Type: application/x-www-form-urlencoded' -d 'id=1&f32=3.2&f64=6.4&email=email&si=100&si=1000&sf32=3.21&sf32=3.22&sf64=6.41&sf64=6.42&ss=str1&ss=str2&tm=2009-02-14T07:31:30.123Z&v=1000000'

curl -L -i -X POST 'http://localhost:8000/example/api/v1/echoargs/echo_multipart_form' -F 'id=1' -F 'f32=3.2' -F 'f64=6.4' -F 'email=email' -F 'si=100' -F 'si=1000' -F 'sf32=3.21' -F 'sf32=3.22' -F 'sf64=6.41' -F 'sf64=6.42' -F 'ss=str1' -F 'ss=str2' -F 'tm=2009-02-14T07:31:30.123Z' -F 'v=1000000'

curl -L -i -X GET 'http://localhost:8000/example/api/v2/echoargs/echo_query?id=1&f32=3.2&f64=6.4&email=email&si=100&si=1000&sf32=3.21&sf32=3.22&sf64=6.41&sf64=6.42&ss=str1&ss=str2&tm=2009-02-14T07:31:30.123Z&v=1000000'

curl -L -i -X POST 'http://localhost:8000/example/api/v2/echoargs/echo_json' -H 'Content-Type: applicatiton/json' -d '{"id": 1, "f32": 3.2, "f64": 6.4, "email": "email", "si": [100, 1000], "sf32": [3.21, 3.22], "sf64": [6.41, 6.42], "ss": ["str1", "str2"], "tm": "2009-02-14T07:31:30.123Z"}'

curl -L -i -X POST 'http://localhost:8000/example/api/v2/echoargs/echo_form' -H 'Content-Type: application/x-www-form-urlencoded' -d 'id=1&f32=3.2&f64=6.4&email=email&si=100&si=1000&sf32=3.21&sf32=3.22&sf64=6.41&sf64=6.42&ss=str1&ss=str2&tm=2009-02-14T07:31:30.123Z&v=1000000'

curl -L -i -X POST 'http://localhost:8000/example/api/v2/echoargs/echo_multipart_form' -F 'id=1' -F 'f32=3.2' -F 'f64=6.4' -F 'email=email' -F 'si=100' -F 'si=1000' -F 'sf32=3.21' -F 'sf32=3.22' -F 'sf64=6.41' -F 'sf64=6.42' -F 'ss=str1' -F 'ss=str2' -F 'tm=2009-02-14T07:31:30.123Z' -F 'v=1000000'