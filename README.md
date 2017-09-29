# binquiry-web

Simple JSON web server for [binquiry](https://github.com/marthjod/binquiry).

Heroku [demo](https://young-caverns-16582.herokuapp.com/).

## Example

```bash
$ go run main.go
Serving on port 4242
```

### GUI 

http://localhost:4242/

### CLI

```bash
$ curl -s 'localhost:4242/?q=kona' | jq '.[] | .gender'
"Feminine"
# etc.
```

For more examples, see [binquiry](https://github.com/marthjod/binquiry).

