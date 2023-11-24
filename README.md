# Handler

Bring the proxy up

```sh
go run main.go
```

Try various OpenAI endpoints
```sh
python ./samples/openai_turbo_new_api.py
python /samples/openai_basic.py
python /samples/openai_otherapis.py
python /samples/openai_streaming.py
python /samples/openai_turbo_new_api.py
```

This approach supports
1. Rapidly rolling out support for new API endpoints
2. Handling streaming
3. Request inspection
4. Response inspection
5. gzip format handling for streaming
6. It should be enhanced to inspect r.Path to only allow for allow-listed endpoints for the company