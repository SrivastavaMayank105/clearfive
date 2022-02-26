# clearfive
# clearfive
i have created a basic logic if user entered the No in the msg in the reques body it will give you error response code and if given Yes in the msg then it will respond 
with the requested ID .Please let me know if you want to make change in the logic of response or any change in the request or reponse struct.

To run the code run the following : go rum cmd/main.go

sample request :
Msg:Yes
curl --location --request POST 'http://localhost:8080/getUserData' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"mayank",
    "age":28,
    "mobNo":9876543210,
    "msg":"Yes"
}'

sample response:
{
    "reqId": 5577006791947779410,
    "reqStatus": {
        "status": "Success",
        "error": "Code:00",
        "msg": "no error"
    }
}

msg:NO
Request:
curl --location --request POST 'http://localhost:8080/getUserData' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"mayank",
    "age":28,
    "mobNo":9876543210,
    "msg":"no"
}'
response
{
    "reqId": -1,
    "reqStatus": {
        "status": "Failed",
        "error": "Code:6998",
        "msg": "no such user info present"
    }
}

when mobile number not equal to 10
curl --location --request POST 'http://localhost:8080/getUserData' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"mayank",
    "age":28,
    "mobNo":987997797900770,
    "msg":"no"
}'
rsponse:
"invalid mobile number"

Age is less than 18
curl --location --request POST 'http://localhost:8080/getUserData' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"mayank",
    "age":17,
    "mobNo":9876543210,
    "msg":"no"
}'

response:
"invalid age"
