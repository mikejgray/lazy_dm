from os import path, getcwd
from json import loads
from random import randint

from fastapi import FastAPI, HTTPException
from mangum import Mangum
from pydantic import BaseSettings


class Settings(BaseSettings):
    GENERATOR_TABLES_LOC = path.join(getcwd(), "data", "tables.json")


def generate_random(base: str, subtype: str):
    return GENERATORS[base][subtype][randint(0, len(GENERATORS[base][subtype]) - 1)]


__version__ = "1.3"
app = FastAPI()
s = Settings()

with open(s.GENERATOR_TABLES_LOC, "r") as f:
    GENERATORS = loads(f.read())


@app.get("/")
async def root():
    return {"message": f"Please see /docs for valid endpoints to use. Lazy DM API {__version__}"}


@app.get("/v1/names")
async def get_names():
    return {
        "given": generate_random("names", "given"),
        "surname": generate_random("names", "surname"),
    }


@app.get("/v1/names/{kind}")
async def get_name(kind: str):
    types = ("given", "surname")
    if kind in types:
        return {kind: generate_random("names", kind)}
    else:
        raise HTTPException(status_code=404, detail=f"Name type not found. Valid types are {', '.join(types)}.")


@app.get("/v1/traps")
async def get_traps():
    return {
        "type": generate_random("traps", "type"),
        "flavor": generate_random("traps", "flavor"),
        "trigger": generate_random("traps", "trigger"),
    }


@app.get("/v1/traps/{kind}")
async def get_trap(kind: str):
    types = ("type", "flavor", "trigger")
    if kind in types:
        return {kind: generate_random("traps", kind)}
    else:
        raise HTTPException(status_code=404, detail=f"Trap type not found. Valid types are {', '.join(types)}.")


@app.get("/v1/monuments")
async def get_monuments():
    return {
        "condition": generate_random("monuments", "condition"),
        "origin": generate_random("monuments", "origin"),
        "type": generate_random("monuments", "type"),
        "effect": generate_random("monuments", "effect"),
    }


@app.get("/v1/monuments/{kind}")
async def get_monument(kind: str):
    types = ("condition", "origin", "type", "effect")
    if kind in types:
        return {kind: generate_random("monuments", kind)}
    else:
        raise HTTPException(status_code=404, detail=f"Monument type not found. Valid types are {', '.join(types)}.")


@app.get("/v1/events")
async def get_events():
    return {
        "mundane": generate_random("events", "mundane"),
        "weather": generate_random("events", "weather"),
        "sentiment": generate_random("events", "sentiment"),
        "fantastic": generate_random("events", "fantastic"),
    }


@app.get("/v1/events/{kind}")
async def get_event(kind: str):
    types = ("mundane", "weather", "sentiment", "fantastic")
    if kind in types:
        return {kind: generate_random("events", kind)}
    else:
        raise HTTPException(status_code=404, detail=f"Event type not found. Valid types are {', '.join(types)}.")


@app.get("/v1/items")
async def get_items():
    return {
        "origin": generate_random("items", "origin"),
        "condition": generate_random("items", "condition"),
        "weapon": generate_random("items", "weapon"),
        "armor": generate_random("items", "armor"),
        "healing": generate_random("items", "healing"),
        "mundane": generate_random("items", "mundane"),
        "spellEffect": generate_random("items", "spellEffect"),
    }


@app.get("/v1/items/{kind}")
async def get_item(kind: str):
    types = ("origin", "condition", "weapon", "armor", "healing", "mundane", "spellEffect")
    if kind in types:
        return {kind: generate_random("items", kind)}
    else:
        raise HTTPException(status_code=404, detail=f"Item type not found. Valid types are {', '.join(types)}.")

handler = Mangum(app=app)
