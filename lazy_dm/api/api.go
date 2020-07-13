package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type nameStruct struct {
	Given   string `json:"given"`
	Surname string `json:"surname"`
}

func name() *nameStruct {
	rand.Seed(time.Now().Unix())
	givenNames := []string{"Sakib", "Jestan", "Natalie", "Lasus", "Teng", "Judithe", "Haina", "Rauffe", "Hassan", "Valentin", "Wulfhilda", "Pholos", "Gerbold", "Emeline", "Bidar", "Hamon", "Pan", "Caplan", "Pippa", "Elizabeta", "Reona", "Thoas", "Lauda", "Abel", "Kapys", "Hadgu", "Beneger", "Roysia", "Demidra", "Eleazar", "Dorothe", "Rickeman", "Emerick", "Eldred", "Celeste", "Leone", "Sophia", "Amelia", "Resul", "Myles", "Abigail", "Ismenia", "Jediah", "Latona", "Simond", "Elliot", "Lettice", "Velic", "Powle", "Syndony", "Navarre", "Kadelon", "Khellus", "Zhen", "Reothine", "Gryffen", "Echo", "Wauter", "Mydrede", "Cornell", "Teukros", "Wyatt", "Keira", "Enis", "Henrietta", "Metope", "Richarde", "Gavin", "Liao", "Elonwey", "Tanah", "Hippote", "Alora", "Levi", "Aedith", "Halla", "Elsebee", "Redmond", "Guthrie", "Archedios", "Linoa", "Kynos", "Logan", "Esperaunce", "Littlejohn", "Karic", "Jonath", "Cunovin", "Yun", "Fulke", "Braya", "Holadamos", "Ophellia", "Cybele", "Cecily", "Athilla", "Sydney", "Kaitlyn", "Timothy", "Ogden", "Musa", "Admiranda", "Hepaklos", "Pullo", "Lysa", "Ophis", "Alkeme", "Anselem", "Ezib", "Wystan", "Symon", "Gared", "Bernaith", "Felice", "Bakis", "Kampe", "Jordyn", "Kalan", "Fauna", "Esmond", "Alyne", "Eschiva", "Ethelred", "Stella", "Joyse", "Katrine", "Zephyrus", "Hafiza", "Edgar", "Sinope", "Nicholina", "Sabra", "Solece", "Callie", "Fahira", "Hildegard", "Bellamy", "Jordan", "Jolice", "Eugenia", "Phanes", "Mathys", "Paz", "Hylas", "Olyffe", "Hart", "Aelina", "Naida", "Somerchild", "Scarlett", "Galain", "Calloway", "Margeria", "Erebos", "Nora", "Orson", "Thearden", "Ierick", "Lorcan", "Tamsin", "Ivan", "Stewart", "Florens", "Aegipan", "Habiba", "Bukolos", "Freya", "Cashel", "Seada", "Muriel", "Clarimond", "Alexis", "Martine", "Arianna", "Bahta", "Lap", "Rhadine", "Cadya", "Galatia", "Minos", "Sydnee", "Landon", "Goddard", "Lambert", "Laurence", "Raoul", "Helvina", "Judye", "Adrian", "Pedias", "Galeos", "Norman", "Olivia", "Leukon", "Almir", "Carmen", "Halgan", "Randwulf", "Janbert", "Primeus", "Lily", "Brunhilde", "Merza", "Zacheus", "Morgan", "Avelyn", "Mestor", "Jeger", "Zineta", "Rebeccah", "Liliana", "Tarvorwen", "Cosmo", "Hemithea", "Thrax", "Symond", "Kurtz", "Pelias", "Lowell", "Persa", "Kaylee", "Iamos", "Madison", "Fridgia", "Ahmet", "Xamso", "Cain", "Myrkenai", "Peny", "Clive", "Josia", "Brice", "Swift", "Adonis", "Daimbert", "Mathild", "Rose", "Denston", "Jaane", "Ingham", "Griffith", "Kenrick", "Kamran", "Theda", "Grayson", "Jillian", "Isylde", "Neria", "Emery", "Germainne", "Helenor", "Bran", "Breccan", "Melusine", "Kain", "Linyeve", "Poter", "Katherine", "Dominic", "Atropos", "Kaivan", "Oghrym", "Lanos", "Drake", "Micah", "Molos", "Banderon", "Havynn", "Hudson", "Lake", "Sigmund", "Bertram", "Abraham", "Wilfrid", "Morgayne", "Anaexia", "Maronne", "Mordrid", "Ambrosia", "Edelinne", "Eugene", "Shui", "Hilda", "Dodonna", "Agando", "Piter", "Folke", "Daanesh", "Orthia", "Jenyfer", "Dagim", "Boyle", "Maddeline", "Eliana", "Oswyn", "Dekelos", "Blake", "Ischys", "Khesrow", "Gregory", "Epione", "Thrydwulf", "Minerva", "Kathe", "May", "Shizuko", "Karyl", "Elenor", "Arnette", "Cathie", "Tammera", "Margrett", "Josue", "Berihu", "Melesse"}
	surnames := []string{"Beastglove", "Glorystar", "Duskheart", "Goblincrippler", "Anvilhunter", "Mooncaster", "Hillmaster", "Jewelcaster", "Songchuckle", "Gravelchest", "Forgegiver", "Needlespur", "Ghostcloud", "Graysoother", "Ebonrazor", "Thornheart", "Duskdancer", "Broadhowler", "Spidertoes", "Felbelly", "Kingson", "Scalehair", "Starhowler", "Giantfang", "Bearsmile", "Bronzeknee", "Greenknocker", "Longhair", "Firespear", "Willowchewer", "Wyrmclaw", "Copperwind", "Gloomtoes", "Ebonhood", "Millblood", "Swordbane", "Gooseflinger", "Moonstar", "Gloomhammer", "Wisefang", "Spidersmasher", "Foebeard", "Bronzecloud", "Drumwillow", "Riddlemaker", "Halfspear", "Iceknee", "Brightwhisper", "Northbound", "Foerazor", "Icebright", "Wormknocker", "Swordstorm", "Foeborn", "Gloryblade", "Gentleblood", "Dustchaser", "Crowsoother", "Darktraveler", "Whitehood", "Eagleson", "Halfrazor", "Goldrain", "Frostgaze", "Nightrend", "Harpkiss", "Giantsoother", "Ebonforger", "Traildancer", "Hillfang", "Dawntraveler", "Gloomtrail", "Catstinger", "Wolfhouse", "Spirithand", "Mountaincleaver", "Coppersmasher", "Greenclaw", "Foxstorm", "Nightstinger", "Gemwing", "Glassforger", "Tallteeth", "Spelltouched", "Swordwing", "Darkgaze", "Wyrmglacier", "Tigerbeard", "Nightborn", "Gloryson", "Oxcutter", "Hawkborn", "Siegewalker", "Oakstalker", "Northrain", "Drakehunter", "Beastcaller", "Icecaller", "Willowsteel", "Tigersmasher", "Bluebeard", "Anvildancer", "Doommaker", "Gemblood", "Felheart", "Graytongue", "Shadowtooth", "Greenbottom", "Rainbrow", "Oakharp", "Flowerheart", "Emeraldforger", "Drakecloud", "Goosekick", "Songdancer", "Felhoof", "Moonsong", "Gentlebelly", "Earthglove", "Broadtooth", "Springwalker", "Needlebelly", "Ratchewer", "Strifetooth", "Dragonstinger", "Foxhowl", "Forgebane", "Halfheart", "Longwhisper", "Willowboot", "Doghunter", "Titanteeth", "Wolfsoul", "Shieldheart", "Greendazer", "Needletoes", "Hillbrow", "Whitefinger", "Iceglove", "Dawnwillow", "Redsong", "Lightwhisker", "Graverock", "Macefinger", "Drumsmasher", "Halfkin", "Gemviper", "Faeriewhisper", "Millstone", "Ghoultalker", "Spiritstar", "Wyrmflinger", "Springwalker", "Graysong", "Leafhouse", "Shieldstorm", "Firewind", "Goldseeker", "Titanchewer", "Northtongue", "Spinevalley", "Wyrmgazer", "Goosehood", "Thornstar", "Leafcaster", "Silvergust", "Forgehammer", "Sharpboot", "Shadowsoul", "Rainblade", "Tallheart", "Ironcrippler", "Frostslicer", "Kingstone", "Silvercutter", "Spellstinger", "Smilesteel", "Wolfwatcher", "Freechaser", "Ratbeard", "Foxsmile", "Trailwhisker", "Hillchaser", "Knifebright", "Leafhound", "Wormbeard", "Angelheart", "Lawknocker", "Lionhunter", "Dusthide", "Stormsong", "Gloomsinger", "Darkcry", "Eaglecry", "Lightboot", "Wormblade", "Mountainbrow", "Wormtoes", "Hawkrunner", "Brightharp", "Crowheart", "Whitewalker", "Bluesoul", "Frostcleaver", "Spinetalker", "Flamehelm", "Crowlover", "Millsong", "Dirtbound", "Emeraldwind", "Quickwhisker", "Hollyflinger", "Doomglacier", "Wormstalker", "Ironbelly", "Moonteeth", "Oakstar", "Ghosttalker", "Goblincloak", "Spiderslicer", "Graymaker", "Goldthumb", "Wisehound", "Ironstalker", "Forgecloak", "Firewhisper", "Glassbuckle", "Starboard", "Stoneknee", "Moonspur", "Gravebright", "Spiritcowl", "Stonesoul"}
	n := &nameStruct{Given: givenNames[rand.Intn(len(givenNames))], Surname: surnames[rand.Intn(len(surnames))]}
	return n
}

type trapStruct struct {
	Types    string `json:"type"`
	Flavors  string `json:"flavor"`
	Triggers string `json:"trigger"`
}

func trap() *trapStruct {
	rand.Seed(time.Now().Unix())
	trapTypes := []string{"Bolts", "Spears", "Scythes", "Bolos", "Spiked chains", "Pit", "Rolling ball", "Crushing pillars", "Darts", "Glyphs", "Swords", "Axes", "Tendrils", "Whips", "Nets", "Bear traps", "Cages", "Beams", "Hammers", "Shurikens"}
	trapFlavors := []string{"Fiery", "Freezing", "Necrotic", "Poisonous", "Acidic", "Thunderous", "Lightning", "Forceful", "Diseased", "Stunning", "Blinding", "Deafening", "Weakening", "Draining", "Sleep-inducing", "Binding", "Dominating", "Psychic", "Maddening", "Confusing"}
	trapTriggers := []string{"Door", "Floor plate", "Tripwire", "Throne", "Corpse", "Chest", "Old book", "Child’s toy", "Jeweled skull", "Beams of light", "Golden angelic statue", "Crystal goblet on pedestal", "Onyx demonic skull", "Jeweled pillar", "Steep stair", "Jeweled crown", "Gilded sarcophagus", "Bound prisoner", "Weapon on an altar", "Idol on pedestal"}
	t := &trapStruct{Types: trapTypes[rand.Intn(len(trapTypes))], Flavors: trapFlavors[rand.Intn(len(trapFlavors))], Triggers: trapTriggers[rand.Intn(len(trapTriggers))]}
	return t
}

type monumentStruct struct {
	Condition string `json:"condition"`
	Origin    string `json:"origin"`
	Type      string `json:"type"`
	Effect    string `json:"effect"`
}

func monument() *monumentStruct {
	rand.Seed(time.Now().Unix())
	condition := []string{"Crumbling", "Sunken", "Pristine", "Excavated", "Vine-covered", "Ruined", "Cracked", "Shattered", "Buried", "Gore-covered", "Bloody", "Glyph-marked", "Rune-scribed", "Obsidian", "Metallic", "Ornate", "Desecrated", "Ancient", "Decorated", "Floating"}
	origin := []string{"Draconic", "Dwarven", "Elven", "Primeval", "Divine", "Unholy", "Abyssal", "Otherworldly", "Orcish", "Undead", "Goblinoid", "Ghoulish", "Vampiric", "Dark elven", "Astral", "Ethereal", "Hellish", "Demonic", "Elemental", "Gnomish"}
	mtype := []string{"Obelisk", "Pillar", "Tomb", "Monolith", "Ruin", "Mosaic", "Ship", "Altar", "Shrine", "Tree", "Statue", "Stone circle", "Throne", "Podium", "Rock", "Fossil", "Fountain", "Mausoleum", "Gravestone", "Cairn", "Geode", "Skull", "Barrow", "Well", "Meteorite", "Archway", "Battlefield", "Charnel pit", "Slaughter field", "Siege engine", "Tower", "Lectern", "Pool", "Orb", "Sarcophagus", "Banner", "Standing stone", "Machine", "Construct", "Keep", "Sundial", "Mirror", "Spire", "Bridge", "Sinkhole", "Effigy", "Gallows", "Ziggurat", "Crystal", "Idol"}
	effect := []string{"Undeath", "Fire", "Madness", "Water", "Radiance", "Arcane", "Poison", "Acid", "Disease", "Psionics", "Frost", "Lightning", "Antimagic", "Ooze", "Charming", "Fear", "Domination", "Sleep", "Thunder", "Tentacles"}
	m := &monumentStruct{Condition: condition[rand.Intn(len(condition))], Origin: origin[rand.Intn(len(origin))], Type: mtype[rand.Intn(len(mtype))], Effect: effect[rand.Intn(len(effect))]}
	return m
}

type eventStruct struct {
	Mundane   string `json:"mundane"`
	Weather   string `json:"weather"`
	Sentiment string `json:"sentiment"`
	Fantastic string `json:"fantastic"`
}

func randomEvent() *eventStruct {
	rand.Seed(time.Now().Unix())
	mundane := []string{"Wedding", "Funeral", "Preparing for war", "Seasonal celebration", "Burning of an effigy", "Death of a noble lord", "Day of drunkenness", "Celebration of lovers", "Great feast", "Execution", "Market day", "Parade of vanquished foes", "Celebration of the dead", "Religious holiday", "Wild boar hat festival", "Robbery", "Brawl", "Visit by the circus", "Wrangling of rampaging beasts", "Festival of kites"}
	weather := []string{"Fog", "Heavy mist", "New moon", "Full moon", "Hot day", "Chilly day", "Light rain", "Moderate rain", "Heavy rain", "Windstorm", "Hailstorm", "Ice storm", "Cloudy day", "Sunny day", "Humid day", "Dry day", "Windy day", "Light snowfall", "Moderate snowfall", "Snowstorm"}
	sentiment := []string{"Happy", "Elated", "Uncaring", "Joyful", "Optimistic", "Pessimistic", "Downtrodden", "Frightened", "Horrified", "Concerned", "Unconcerned", "Harried", "Sleep-deprived", "Dazed", "Hyperactive", "Purposeful", "Lazy", "Melancholy", "Busy", "Suspicious"}
	fantastic := []string{"The stars have disappeared from the sky", "An unexpected solar eclipse", "The blood moon rises", "Swarms of stinging insects descend", "Acidic fog rolls in", "A second sun appears in the sky", "A storm of arcane energy", "The arrival of a servant of a god", "Meteor shower", "A cyclopean behemoth rises", "Swarms of mischievous devils", "Tentacles appear in the sky", "The dancing dead come to life", "Volcanic eruption", "Collapsing sinkhole reveals ancient ruins below", "The sun does not rise", "A great floating tower appears", "The lord's castle disappears", "The border to the fey realm grows thin", "The world of shadow bleeds over into the material realm"}
	e := &eventStruct{Mundane: mundane[rand.Intn(len(mundane))], Weather: weather[rand.Intn(len(weather))], Sentiment: sentiment[rand.Intn(len(sentiment))], Fantastic: fantastic[rand.Intn(len(fantastic))]}
	return e
}

type items struct {
	Origin      string `json:"origin"`
	Condition   string `json:"condition"`
	Weapon      string `json:"weapon"`
	Armor       string `json:"armor"`
	Healing     string `json:"healing"`
	Mundane     string `json:"mundane"`
	SpellEffect string `json:"spell_effect"`
}

func randomItem() *items {
	rand.Seed(time.Now().Unix())
	origin := []string{"Draconic", "Dwarven", "Elven", "Primeval", "Divine", "Unholy", "Abyssal", "Otherworldly", "Orcish", "Undead", "Goblinoid", "Ghoulish", "Vampiric", "Dark elven", "Astral", "Ethereal", "Hellish", "Demonic", "Elemental", "Gnomish"}
	condition := []string{"Grimy", "Chipped", "Rough", "Smooth", "Ancient", "Crumbling", "Pristine", "Cool", "Ornate", "Plain", "Rune-scribed", "Carved", "Decorated", "Delicate", "Burned", "Oily", "Pulsing", "Glowing", "Shining", "Smoldering"}
	weapon := []string{"Dagger", "Mace", "Quarterstaff", "Spear", "Light crossbow", "Shortbow", "Battleaxe", "Flail", "Glaive", "Greataxe", "Greatsword", "Longsword", "Maul", "Morningstar", "Rapier", "Scimitar", "Shortsword", "Warhammer", "Heavy crossbow", "Longbow"}
	armor := []string{"Leather", "Studded leather", "Hide", "Chain shirt", "Scale mail", "Breastplate", "Half plate", "Ring mail", "Chain mail", "Splint", "Plate", "Shield"}
	healing := []string{"Healing (Common) 2d4+2", "Greater healing (Uncommon) 4d4+4", "Superior healing (Rare) 8d4+8", "Supreme healing (Very rare) 10d4+20"}
	mundane := []string{"Amulet", "Arrowhead", "Bell", "Bird skull", "Bone", "Bowl", "Box", "Bracelet", "Brooch", "Buckle", "Candle", "Coin", "Crown", "Cup", "Dagger", "Disc", "Earring", "Figurine", "Finger bone", "Flute", "Forked rod", "Gemstone", "Glove", "Goblet", "Hammer", "Idol", "Jewelry box", "Key", "Lamp", "Mask", "Medallion", "Mirror", "Necklace", "Opal", "Orb", "Pipe", "Quill", "Ring", "Rod", "Skull", "Sphere", "Spike", "Statue", "Stone", "String of beads", "Symbol", "Tiara", "Tooth", "Vial", "Wand"}
	spellEffect := []string{"Light", "Bane", "Bless", "Cure wounds", "Detect evil and good", "Detect magic", "Guiding bolt", "Inflict wounds", "Shield of faith", "Blindness/deafness", "Silence", "Bestow curse", "Dispel magic", "Flame strike", "Insect plague", "Acid splash", "Shocking grasp", "True strike", "Burning hands", "Charm person", "Color spray", "Comprehend languages", "Detect magic", "Fog cloud", "Jump", "Sleep", "Thunderwave", "Acid arrow", "Invisibility", "Misty step", "Ray of enfeeblement", "Scorching ray", "Shatter", "Web", "Fear", "Fly", "Gaseous form", "Haste", "Lightning bolt", "Slow", "Stinking cloud", "Banishment", "Black tentacles", "Blight", "Fire shield", "Ice storm", "Stoneskin", "Cloudkill", "Cone of cold", "Disintegrate"}
	i := &items{Origin: origin[rand.Intn(len(origin))], Condition: condition[rand.Intn(len(condition))], Weapon: weapon[rand.Intn(len(weapon))], Armor: armor[rand.Intn(len(armor))], Healing: healing[rand.Intn(len(healing))], Mundane: mundane[rand.Intn(len(mundane))], SpellEffect: spellEffect[rand.Intn(len(spellEffect))]}
	return i
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Original content courtesy of slyflourish.com, try an endpoint")
}

func getName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(name())
}

func getTrap(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(trap())
}

func getMonument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(monument())
}

func getEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(randomEvent())
}

func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(randomItem())
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
	log.Fatal(http.ListenAndServe(":2020", c.Handler(myRouter)))
}

func main() {
	fmt.Println("Lazy DM Guide Workbook API v1.0")
	handleRequests()
}
