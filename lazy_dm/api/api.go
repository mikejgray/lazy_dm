package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type tables struct {
	Names     nameStruct     `json:"names"`
	Traps     trapStruct     `json:"traps"`
	Monuments monumentStruct `json:"monuments"`
	Events    eventStruct    `json:"events"`
	Items     items          `json:"items"`
}

type nameStruct struct {
	Given   []string `json:"given"`
	Surname []string `json:"surname"`
}

type trapStruct struct {
	Types    []string `json:"type"`
	Flavors  []string `json:"flavor"`
	Triggers []string `json:"trigger"`
}

type monumentStruct struct {
	Condition []string `json:"condition"`
	Origin    []string `json:"origin"`
	Type      []string `json:"type"`
	Effect    []string `json:"effect"`
}

type eventStruct struct {
	Mundane   []string `json:"mundane"`
	Weather   []string `json:"weather"`
	Sentiment []string `json:"sentiment"`
	Fantastic []string `json:"fantastic"`
}

type items struct {
	Origin      []string `json:"origin"`
	Condition   []string `json:"condition"`
	Weapon      []string `json:"weapon"`
	Armor       []string `json:"armor"`
	Healing     []string `json:"healing"`
	Mundane     []string `json:"mundane"`
	SpellEffect []string `json:"spellEffect"`
}

var allTables tables

func name(key ...string) *nameStruct {
	rand.Seed(time.Now().Unix())

	switch query := key[0]; query {
	case "given":
		return &nameStruct{
			Given:   []string{allTables.Names.Given[rand.Intn(len(allTables.Names.Given))]},
			Surname: []string{"None"},
		}
	case "surname":
		return &nameStruct{
			Given:   []string{"None"},
			Surname: []string{allTables.Names.Surname[rand.Intn(len(allTables.Names.Surname))]},
		}
	default:
		return &nameStruct{
			Given:   []string{allTables.Names.Given[rand.Intn(len(allTables.Names.Given))]},
			Surname: []string{allTables.Names.Surname[rand.Intn(len(allTables.Names.Surname))]},
		}
	}
}

func trap(key ...string) *trapStruct {
	rand.Seed(time.Now().Unix())

	switch query := key[0]; query {
	case "type":
		return &trapStruct{
			Types:    []string{allTables.Traps.Types[rand.Intn(len(allTables.Traps.Types))]},
			Flavors:  []string{"None"},
			Triggers: []string{"None"},
		}
	case "flavor":
		return &trapStruct{
			Types:    []string{"None"},
			Flavors:  []string{allTables.Traps.Flavors[rand.Intn(len(allTables.Traps.Flavors))]},
			Triggers: []string{"None"},
		}
	case "trigger":
		return &trapStruct{
			Types:    []string{"None"},
			Flavors:  []string{"None"},
			Triggers: []string{allTables.Traps.Triggers[rand.Intn(len(allTables.Traps.Triggers))]},
		}
	default:
		return &trapStruct{
			Types:    []string{allTables.Traps.Types[rand.Intn(len(allTables.Traps.Types))]},
			Flavors:  []string{allTables.Traps.Flavors[rand.Intn(len(allTables.Traps.Flavors))]},
			Triggers: []string{allTables.Traps.Triggers[rand.Intn(len(allTables.Traps.Triggers))]},
		}
	}
}

func monument(key ...string) *monumentStruct {
	rand.Seed(time.Now().Unix())

	switch query := key[0]; query {
	case "condition":
		return &monumentStruct{
			Condition: []string{allTables.Monuments.Condition[rand.Intn(len(allTables.Monuments.Condition))]},
			Origin:    []string{"None"},
			Type:      []string{"None"},
			Effect:    []string{"None"},
		}
	case "origin":
		return &monumentStruct{
			Condition: []string{"None"},
			Origin:    []string{allTables.Monuments.Origin[rand.Intn(len(allTables.Monuments.Origin))]},
			Type:      []string{"None"},
			Effect:    []string{"None"},
		}
	case "type":
		return &monumentStruct{
			Condition: []string{"None"},
			Origin:    []string{"None"},
			Type:      []string{allTables.Monuments.Type[rand.Intn(len(allTables.Monuments.Type))]},
			Effect:    []string{"None"},
		}
	case "effect":
		return &monumentStruct{
			Condition: []string{"None"},
			Origin:    []string{"None"},
			Type:      []string{"None"},
			Effect:    []string{allTables.Monuments.Effect[rand.Intn(len(allTables.Monuments.Effect))]},
		}
	default:
		return &monumentStruct{
			Condition: []string{allTables.Monuments.Condition[rand.Intn(len(allTables.Monuments.Condition))]},
			Origin:    []string{allTables.Monuments.Origin[rand.Intn(len(allTables.Monuments.Origin))]},
			Type:      []string{allTables.Monuments.Type[rand.Intn(len(allTables.Monuments.Type))]},
			Effect:    []string{allTables.Monuments.Effect[rand.Intn(len(allTables.Monuments.Effect))]},
		}
	}
}

func randomEvent(key ...string) *eventStruct {
	rand.Seed(time.Now().Unix())

	switch query := key[0]; query {
	case "mundane":
		return &eventStruct{
			Mundane:   []string{allTables.Events.Mundane[rand.Intn(len(allTables.Events.Mundane))]},
			Weather:   []string{"None"},
			Sentiment: []string{"None"},
			Fantastic: []string{"None"},
		}
	case "weather":
		return &eventStruct{
			Mundane:   []string{"None"},
			Weather:   []string{allTables.Events.Weather[rand.Intn(len(allTables.Events.Weather))]},
			Sentiment: []string{"None"},
			Fantastic: []string{"None"},
		}
	case "sentiment":
		return &eventStruct{
			Mundane:   []string{"None"},
			Weather:   []string{"None"},
			Sentiment: []string{allTables.Events.Sentiment[rand.Intn(len(allTables.Events.Sentiment))]},
			Fantastic: []string{"None"},
		}
	case "fantastic":
		return &eventStruct{
			Mundane:   []string{"None"},
			Weather:   []string{"None"},
			Sentiment: []string{"None"},
			Fantastic: []string{allTables.Events.Fantastic[rand.Intn(len(allTables.Events.Fantastic))]},
		}
	default:
		return &eventStruct{
			Mundane:   []string{allTables.Events.Mundane[rand.Intn(len(allTables.Events.Mundane))]},
			Weather:   []string{allTables.Events.Weather[rand.Intn(len(allTables.Events.Weather))]},
			Sentiment: []string{allTables.Events.Sentiment[rand.Intn(len(allTables.Events.Sentiment))]},
			Fantastic: []string{allTables.Events.Fantastic[rand.Intn(len(allTables.Events.Fantastic))]},
		}
	}
}

func randomItem(key ...string) *items {
	rand.Seed(time.Now().Unix())

	switch query := key[0]; query {
	case "origin":
		return &items{
			Origin:      []string{allTables.Items.Origin[rand.Intn(len(allTables.Items.Origin))]},
			Condition:   []string{"None"},
			Weapon:      []string{"None"},
			Armor:       []string{"None"},
			Healing:     []string{"None"},
			Mundane:     []string{"None"},
			SpellEffect: []string{"None"},
		}
	case "condition":
		return &items{
			Origin:      []string{"None"},
			Condition:   []string{allTables.Items.Condition[rand.Intn(len(allTables.Items.Condition))]},
			Weapon:      []string{"None"},
			Armor:       []string{"None"},
			Healing:     []string{"None"},
			Mundane:     []string{"None"},
			SpellEffect: []string{"None"},
		}
	case "weapon":
		return &items{
			Origin:      []string{"None"},
			Condition:   []string{"None"},
			Weapon:      []string{allTables.Items.Weapon[rand.Intn(len(allTables.Items.Weapon))]},
			Armor:       []string{"None"},
			Healing:     []string{"None"},
			Mundane:     []string{"None"},
			SpellEffect: []string{"None"},
		}
	case "armor":
		return &items{
			Origin:      []string{"None"},
			Condition:   []string{"None"},
			Weapon:      []string{"None"},
			Armor:       []string{allTables.Items.Armor[rand.Intn(len(allTables.Items.Armor))]},
			Healing:     []string{"None"},
			Mundane:     []string{"None"},
			SpellEffect: []string{"None"},
		}
	case "healing":
		return &items{
			Origin:      []string{"None"},
			Condition:   []string{"None"},
			Weapon:      []string{"None"},
			Armor:       []string{"None"},
			Healing:     []string{allTables.Items.Healing[rand.Intn(len(allTables.Items.Healing))]},
			Mundane:     []string{"None"},
			SpellEffect: []string{"None"},
		}
	case "mundane":
		return &items{
			Origin:      []string{"None"},
			Condition:   []string{"None"},
			Weapon:      []string{"None"},
			Armor:       []string{"None"},
			Healing:     []string{"None"},
			Mundane:     []string{allTables.Items.Mundane[rand.Intn(len(allTables.Items.Mundane))]},
			SpellEffect: []string{"None"},
		}
	case "spellEffect":
		return &items{
			Origin:      []string{"None"},
			Condition:   []string{"None"},
			Weapon:      []string{"None"},
			Armor:       []string{"None"},
			Healing:     []string{"None"},
			Mundane:     []string{"None"},
			SpellEffect: []string{allTables.Items.SpellEffect[rand.Intn(len(allTables.Items.SpellEffect))]},
		}
	default:
		return &items{
			Origin:      []string{allTables.Items.Origin[rand.Intn(len(allTables.Items.Origin))]},
			Condition:   []string{allTables.Items.Condition[rand.Intn(len(allTables.Items.Condition))]},
			Weapon:      []string{allTables.Items.Weapon[rand.Intn(len(allTables.Items.Weapon))]},
			Armor:       []string{allTables.Items.Armor[rand.Intn(len(allTables.Items.Armor))]},
			Healing:     []string{allTables.Items.Healing[rand.Intn(len(allTables.Items.Healing))]},
			Mundane:     []string{allTables.Items.Mundane[rand.Intn(len(allTables.Items.Mundane))]},
			SpellEffect: []string{allTables.Items.SpellEffect[rand.Intn(len(allTables.Items.SpellEffect))]},
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Original content courtesy of slyflourish.com, try an endpoint on /v1")
}

func getName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	keyType, ok := r.URL.Query()["type"]
	if !ok || len(keyType[0]) < 1 {
		json.NewEncoder(w).Encode(&nameStruct{
			Given:   []string{allTables.Names.Given[rand.Intn(len(allTables.Names.Given))]},
			Surname: []string{allTables.Names.Surname[rand.Intn(len(allTables.Names.Surname))]},
		})
	} else {
		key := keyType[0]
		json.NewEncoder(w).Encode(name(key))
	}
}

func getTrap(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	keyType, ok := r.URL.Query()["type"]
	if !ok || len(keyType[0]) < 1 {
		json.NewEncoder(w).Encode(&trapStruct{
			Types:    []string{allTables.Traps.Types[rand.Intn(len(allTables.Traps.Types))]},
			Flavors:  []string{allTables.Traps.Flavors[rand.Intn(len(allTables.Traps.Flavors))]},
			Triggers: []string{allTables.Traps.Triggers[rand.Intn(len(allTables.Traps.Triggers))]},
		})
	} else {
		key := keyType[0]
		json.NewEncoder(w).Encode(trap(key))
	}
}

func getMonument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	keyType, ok := r.URL.Query()["type"]
	if !ok || len(keyType[0]) < 1 {
		json.NewEncoder(w).Encode(&monumentStruct{
			Condition: []string{allTables.Monuments.Condition[rand.Intn(len(allTables.Monuments.Condition))]},
			Origin:    []string{allTables.Monuments.Origin[rand.Intn(len(allTables.Monuments.Origin))]},
			Type:      []string{allTables.Monuments.Type[rand.Intn(len(allTables.Monuments.Type))]},
			Effect:    []string{allTables.Monuments.Effect[rand.Intn(len(allTables.Monuments.Effect))]},
		})
	} else {
		key := keyType[0]
		json.NewEncoder(w).Encode(monument(key))
	}
}

func getEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	keyType, ok := r.URL.Query()["type"]
	if !ok || len(keyType[0]) < 1 {
		json.NewEncoder(w).Encode(&eventStruct{
			Mundane:   []string{allTables.Events.Mundane[rand.Intn(len(allTables.Events.Mundane))]},
			Weather:   []string{allTables.Events.Weather[rand.Intn(len(allTables.Events.Weather))]},
			Sentiment: []string{allTables.Events.Sentiment[rand.Intn(len(allTables.Events.Sentiment))]},
			Fantastic: []string{allTables.Events.Fantastic[rand.Intn(len(allTables.Events.Fantastic))]},
		})
	} else {
		key := keyType[0]
		json.NewEncoder(w).Encode(randomEvent(key))
	}
}

func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	keyType, ok := r.URL.Query()["type"]
	if !ok || len(keyType[0]) < 1 {
		json.NewEncoder(w).Encode(randomItem())
	} else {
		key := keyType[0]
		json.NewEncoder(w).Encode(randomItem(key))
	}
}

func importTables(filename string) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &allTables)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},   // All origins
		AllowedMethods: []string{"GET"}, // Allowing only get
	})

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/v1/names", getName)
	myRouter.HandleFunc("/v1/traps", getTrap)
	myRouter.HandleFunc("/v1/monuments", getMonument)
	myRouter.HandleFunc("/v1/events", getEvent)
	myRouter.HandleFunc("/v1/items", getItem)
	log.Fatal(http.ListenAndServe(":8080", c.Handler(myRouter)))
}

func main() {
	data, err := os.LookupEnv("DATA")
	if err == true || len(data) == 0 {
		data = filepath.FromSlash("./tables.json")
	}
	importTables(data)
	fmt.Println("Lazy DM Guide Workbook API v1.2")
	handleRequests()
}
