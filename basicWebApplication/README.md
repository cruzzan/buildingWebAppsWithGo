Simple app that takes requests to `/` and `/markdown` on port 8181.

`/`: Root, a form for writing raw markdown

`/markdown`: Endpoint that receives the raw markdown and returns html generated from it.

Run the app by executing `make build-and-run` and then going to `http://localhost:8181`
