from random import choice, seed
from typing import Optional

from flask import Flask, render_template, url_for

app = Flask(__name__)
app.config.from_object('config')

# Routes
@app.route("/", methods=["GET"])
def index():
    # https://stackoverflow.com/questions/13317536/get-list-of-all-routes-defined-in-the-flask-app
    links = []
    for rule in app.url_map.iter_rules():
        # Filter out rules we can't navigate to in a browser
        # and rules that require parameters
        if "GET" in rule.methods \
        and has_no_empty_params(rule) and \
        url_for(rule.endpoint, **(rule.defaults or {})) != "/":
            url = url_for(rule.endpoint, **(rule.defaults or {}))
            links.append((url, rule.endpoint))
    # links is now a list of url, endpoint tuples
    return render_template("index.html", links=links)

@app.route("/names", methods=["GET"])
def names():
    return render_template("name.html", name=generate_name())

@app.route("/traps", methods=["GET"])
def traps():
    return render_template("traps.html", trap=generate_trap())

@app.route("/monuments", methods=["GET"])
def monuments():
    return render_template("monuments.html", monument=generate_monument())

@app.route("/magic_items", methods=["GET"])
def magic_items():
    return render_template("items.html",
                            weapon=generate_item("weapon"),
                            armor=generate_item("armor"),
                            healing=generate_item("healing"))


# Helper functions
def has_no_empty_params(rule) -> bool:
    defaults = rule.defaults if rule.defaults is not None else ()
    arguments = rule.arguments if rule.arguments is not None else ()
    return len(defaults) >= len(arguments)

def generate_name() -> str:
    """Generate a random name from a list of given names and surnames."""
    seed()
    return f"{choice(app.config['GIVEN_NAMES'])} {choice(app.config['SURNAMES'])}"


def generate_trap() -> dict:
    """Generate a trap with type, flavor, and trigger."""
    seed()
    traps = app.config["TRAPS"]
    return {"type": choice(traps["type"]),
            "flavor": choice(traps["flavor"]),
            "trigger": choice(traps["trigger"])}


def generate_monument() -> dict:
    """Generate a monument with condition, origin, physical type, and effect."""
    seed()
    monument = app.config["MONUMENTS"]
    return {"condition": choice(monument["condition"]),
            "origin": choice(monument["origin"]),
            "type": choice(monument["type"]),
            "effect": choice(monument["effect"])}


def generate_item(item_type: str) -> Optional[dict]:
    """Generate a random magical item with condition, origin. Item types are healing, weapon, or armor."""
    if item_type not in ("healing", "weapon", "armor"):
        return None
    seed()
    items = app.config["ITEMS"]
    if item_type == "healing":
        return choice(items["healing"])
    return {"condition": choice(items["condition"]),
            "origin": choice(items["origin"]),
            "type": choice(items[item_type])
    }
