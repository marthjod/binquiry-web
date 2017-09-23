# binquiry-web

Simple JSON web server for [binquiry](https://github.com/marthjod/binquiry).

## Example

```bash
$ go run main.go
Serving on port 4242
```

```bash
$ curl -s 'localhost:4242/?q=kona' | jq '.[] | .gender'
"Feminine"
# etc.

$ curl -I 'localhost:4242/?foo=bar'
HTTP/1.1 400 Bad Request
```

