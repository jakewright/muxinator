# Muxinator

Muxinator is a simple wrapper around the [`gorilla/mux`](https://github.com/gorilla/mux) and [`urfave/negroni`](https://github.com/urfave/negroni) packages to make it easier to build an HTTP router with middleware.

## Example

```go
router := muxinator.NewRouter()
router.AddMiddleware(globalMiddleware)
router.Get("/path", middleware1, middleware2)
router.Patch("/path", middleware2, middleware3)
http.ListenAndServe(":80", router.BuildHandler())
```
