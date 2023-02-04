# responses

This is a very useful utility package to ease web development by removing the mindane task of handling web responses.

## Installation

```golang
go get github.com/gouniverse/responses
```

## Usage

### Direct Usage

```golang
router.Get("/hello-world", func(w http.ResponseWriter, r *http.Request) {
  responses.HTMLResponse(w, r, "Hello world")
})
```

## Functional Usage

```
router.Get("/hello-world", func(w http.ResponseWriter, r *http.Request) {
  responses.HTMLResponseF(w, r, MyHelloWorldController)
})

func MyHelloWorldController(w http.ResponseWriter, r *http.Request) string {
    return "Hello World"
}
```

## Methods available
