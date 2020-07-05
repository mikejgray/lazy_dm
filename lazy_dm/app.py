from random import choice, seed

from flask import Flask, render_template

app = Flask(__name__)
app.config.from_object('config')


@app.route("/", methods=["GET"])
def index():
    return render_template("index.html")

@app.route("/name", methods=["GET"])
def name():
    return render_template("name.html", name=generate_name())

@app.route("/traps", methods=["GET"])
def trap():
    return render_template("traps.html", trap=generate_trap())


def generate_name():
    """Generate a random name from a list of given names and surnames."""
    seed()
    return f"{choice(app.config['GIVEN_NAMES'])} {choice(app.config['SURNAMES'])}"


def generate_trap():
    """Generate a trap with type, flavor, and trigger."""
    traps = app.config["TRAPS"]
    return {"type": choice(traps["type"]),
            "flavor": choice(traps["flavor"]),
            "trigger": choice(traps["trigger"])}
