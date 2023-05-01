# responses <a href="https://github.com/gouniverse/responses" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

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

## Methods Available

- CSSResponse(w http.ResponseWriter, r *http.Request, body string)
- CSSResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string)
- HTMLResponse(w http.ResponseWriter, r *http.Request, body string)
- HTMLResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string)
- JSResponse(w http.ResponseWriter, r *http.Request, body string)
- JSResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string)
- JSONResponse(w http.ResponseWriter, r *http.Request, body string)
- JSONResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string)
- XMLResponse(w http.ResponseWriter, r *http.Request, body string)
- XMLResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string)

## GZip Compressed Methods Available
- GzipCSSResponse(w http.ResponseWriter, r *http.Request, body string)
- GzipCSSResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string)
- GzipHTMLResponse(w http.ResponseWriter, r *http.Request, body string)
- GzipHTMLResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string)
- GzipJSResponse(w http.ResponseWriter, r *http.Request, body string)
- GzipJSResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string)
- GzipJSONResponse(w http.ResponseWriter, r *http.Request, body string)
- GzipJSONResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string)
- GzipXMLResponse(w http.ResponseWriter, r *http.Request, body string)
- GzipXMLResponseF(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request) string)
