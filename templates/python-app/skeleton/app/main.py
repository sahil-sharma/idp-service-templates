import os

from fastapi import FastAPI

NAME = "${{ values.name }}"
VERSION = os.getenv("APP_VERSION", "dev")
MESSAGE = os.getenv("APP_MESSAGE", "Hello from " + NAME)

app = FastAPI(title=NAME, version=VERSION)


@app.get("/")
def root():
    return {"service": NAME, "message": MESSAGE}


@app.get("/healthz")
def healthz():
    return {"status": "ok"}


@app.get("/version")
def version():
    return {"version": VERSION}
