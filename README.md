# Silver-telegram

The program exposes an HTTP interface with two endpoints:

1. `POST /learn` expecting a body of text, sent as the POST body with `Content-Type: text/plain`.
2. `GET /generate` returns randomly-generated text based on all the trigrams that have been learned since starting the program.

Note: to chnage the complexity of what is returned in the second endpoint, a query of complexity=10 can be added

```
$ curl --data-binary @pride-prejudice.txt localhost:8080/learn
$ curl localhost:8080/generate
```

## EXPECTED RESPONSES

```
{
    "message": "Trigram collected Successfully",
    "success": true
}

{
    "message": "Empty data set",
    "success": false
}

{
    "message": "to ask them she said herself; and Jane�s offences felicity followed, and not know how",
    "success": true
}
```

## MAKE COMMANDS

```
Run Go Test
make test

Run Go Vet
make vet

Run Go Lint
make lint

Run Go Mock to generate mock files
make gen-mocks
```
