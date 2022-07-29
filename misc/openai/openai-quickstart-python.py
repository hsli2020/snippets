## app.py

import os

import openai
from flask import Flask, redirect, render_template, request, url_for

app = Flask(__name__)
openai.api_key = os.getenv("OPENAI_API_KEY")


@app.route("/", methods=("GET", "POST"))
def index():
    if request.method == "POST":
        animal = request.form["animal"]
        response = openai.Completion.create(
            model="text-davinci-003",
            prompt=generate_prompt(animal),
            temperature=0.6,
        )
        return redirect(url_for("index", result=response.choices[0].text))

    result = request.args.get("result")
    return render_template("index.html", result=result)


def generate_prompt(animal):
    return """Suggest three names for an animal that is a superhero.

Animal: Cat
Names: Captain Sharpclaw, Agent Fluffball, The Incredible Feline
Animal: Dog
Names: Ruff the Protector, Wonder Canine, Sir Barks-a-Lot
Animal: {}
Names:""".format(
        animal.capitalize()
    )

## templates/index.html

<!DOCTYPE html>
<head>
  <title>OpenAI Quickstart</title>
  <link
    rel="shortcut icon"
    href="{{ url_for('static', filename='dog.png') }}"
  />
  <link rel="stylesheet" href="{{ url_for('static', filename='main.css') }}" />
</head>

<body>
  <img src="{{ url_for('static', filename='dog.png') }}" class="icon" />
  <h3>Name my pet</h3>
  <form action="/" method="post">
    <input type="text" name="animal" placeholder="Enter an animal" required />
    <input type="submit" value="Generate names" />
  </form>
  {% if result %}
  <div class="result">{{ result }}</div>
  {% endif %}
</body>

## env

FLASK_APP=app
FLASK_ENV=development

# The API key should remain private.
OPENAI_API_KEY=

# requirements.txt

autopep8==1.6.0
certifi==2021.10.8
charset-normalizer==2.0.7
click==8.0.3
et-xmlfile==1.1.0
Flask==2.0.2
idna==3.3
itsdangerous==2.0.1
Jinja2==3.0.2
MarkupSafe==2.0.1
numpy==1.21.3
openai==0.19.0
openpyxl==3.0.9
pandas==1.3.4
pandas-stubs==1.2.0.35
pycodestyle==2.8.0
python-dateutil==2.8.2
python-dotenv==0.19.2
pytz==2021.3
requests==2.26.0
six==1.16.0
toml==0.10.2
tqdm==4.62.3
urllib3==1.26.7
Werkzeug==2.0.2

# README.md

This is an example pet name generator app used in the OpenAI API quickstart tutorial. 
It uses the Flask web framework. Check out the tutorial or follow the instructions 
below to get set up.

1.  If you don’t have Python installed, install it from here
2.  Clone this repository

3.  Navigate into the project directory
    $ cd openai-quickstart-python

4.  Create a new virtual environment
    $ python -m venv venv
    $ . venv/bin/activate

5.  Install the requirements
    $ pip install -r requirements.txt

6.  Make a copy of the example environment variables file
    $ cp .env.example .env

7.  Add your API key to the newly created .env file
8.  Run the app
    $ flask run

You should now be able to access the app at http://localhost:5000! 
For the full context behind this example app, check out the tutorial. 

https://beta.openai.com/docs/quickstart