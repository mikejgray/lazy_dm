# The Lazy DM's Workbook - The App

[The Lazy DM's Workbook](https://slyflourish.com/lazydmsworkbook/), by [Michael E. Shea (aka SlyFlourish)](https://slyflourish.com/start_here.html), is an outstanding publication that builds on his previous work to make DMing simpler, more fun, and overall more epic. I'm designing this app to dynamically generate content from it with the click of a button to make my life easier. All credit for the work goes to SlyFlourish. I hope that, one day, I'll be able to share this with him.

## Installation/Running

* To build the API as a binary:
```go
go build ./lazy_dm/api/api.go
```

* To run the API as a Docker container:
```bash
cd ./lazy_dm/api
docker build --pull --rm -t lazydm:latest
docker run -p 8080:8080 --rm -d lazydm:latest
```

* To run in GCP App Engine:
```bash
git clone https://github.com/mikejgray/lazy_dm
cd ./lazy_dm/api
gcloud app deploy
```

The site itself can be hosted in a cloud provider bucket, such as AWS S3 or GCP Storage, or run locally. Be sure to change the `Api` variable in the JavaScript file to match where you're actually hosting the API.

## API Endpoints

* Random Names (`/names`)
* Random Traps (`/traps`)
* Random Monuments (`/monuments`)
* Random Magic Items (`/items`)
* Random Town Events (`/events`)

### Endpoint Query Strings

Each endpoint accepts a `?type=` query string that conforms to the individual types for each endpoint:

`/names`:

* given
* surname

`/traps`:

* type
* flavor
* trigger

`/monuments`:

* condition
* origin
* type
* effect

`/events`:

* mundane
* weather
* sentiment
* fantastic

`/items`:

* origin
* condition
* weapon
* armor
* healing
* mundane
* spellEffect

**Example:**

`https://localhost:8080/v1/items?type=spellEffect` will return something to the effect of:

```json
{"origin":["None"],"condition":["None"],"weapon":["None"],"armor":["None"],"healing":["None"],"mundane":["None"],"spellEffect":["Insect plague"]}
```

## API TODO Items
* Clean up individual type requests so it only returns the necessary data
* All for a `?count` query parameter to return numerous items

## UI TODO Items
* Add radio buttons for endpoint types
* Add count slider for multiple responses
* Add checkboxes to combine multiple endpoint type responsese
* Create a magic puzzle box that can be manipulated to return random endpoint information
