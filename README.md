# run script to Deno

- Example `scripts.json`

```json
{
    "start": "deno run --allow-net webserver.ts",
    "build": "deno compile webserver.ts"
}
```

## Compile `rundeno`

`go build ./rundeno.go`

## Use task

`./rundeno build` or `./rundeno start`
