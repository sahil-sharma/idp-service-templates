# ${{ values.name }}

${{ values.description }}

Created from the Backstage **Python service** golden path.

```bash
pip install -r requirements.txt
uvicorn app.main:app --port 8080
curl localhost:8080/ localhost:8080/healthz localhost:8080/version
```
