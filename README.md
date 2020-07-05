# The Lazy DM's Workbook - The App

[The Lazy DM's Workbook](https://slyflourish.com/lazydmsworkbook/), by (Michael E. Shea (aka SlyFlourish))[https://slyflourish.com/start_here.html], is an outstanding publication that builds on his previous work to make DMing simpler, more fun, and overall more epic. I'm designing this app to dynamically generate content from it with the click of a button to make my life easier. All credit for the work goes to SlyFlourish. I hope that, one day, I'll be able to share this with him.

## Installation/Running

I'm using [Poetry](https://python-poetry.org/). Install that, make sure you're running Python 3.7+, then:

* `poetry install`
* `FLASK_APP="./lazy_dm/app.py" poetry run flask run`

Eventually this will be easier, and support something besides poetry, but for now this is how it works.

## Features

* Random Names (`/name`)
* Random Traps (`/traps`)
* Random Monuments (`/monuments`)
