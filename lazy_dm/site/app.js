var Api = 'https://dm-api.graywind.org'

function getName() {
    namePlace = document.getElementById('name');
    fetch(`${Api}/v1/names`)
        .then(response => response.json())
        .then(data => namePlace.textContent = `${data.given[0]} ${data.surname[0]}`);
};

function getEvent() {
    eventPlace = document.getElementById('event');
    fetch(`${Api}/v1/events`)
        .then(response => response.json())
        .then(data => eventPlace.innerHTML = (
            `<p>Mundane Event: ${data.mundane[0]}<p>Town Sentiment: ${data.sentiment[0]}<p>Notable Weather: ${data.weather[0]}<p>Fantastic Event: ${data.fantastic[0]}`
        ));
};

function getMonument() {
    namePlace = document.getElementById('monument');
    fetch(`${Api}/v1/monuments`)
        .then(response => response.json())
        .then(data => namePlace.innerHTML =
            `${data.condition[0]} ${data.origin[0]} ${data.type[0]} with a ${data.effect[0]} effect`);
};

function getTrap() {
    eventPlace = document.getElementById('trap');
    fetch(`${Api}/v1/traps`)
        .then(response => response.json())
        .then(data => eventPlace.innerHTML = (
            `<p>Type: ${data.type[0]}<p>Flavor: ${data.flavor[0]}<p>Trigger: ${data.trigger[0]}`
        ));
};

function getItem() {
    eventPlace = document.getElementById('item');
    fetch(`${Api}/v1/items`)
        .then(response => response.json())
        .then(data => eventPlace.innerHTML = (
            `<p>Origin: ${data.origin[0]}<p>Condition: ${data.condition[0]}<p>Weapon: ${data.weapon[0]}<p>Armor: ${data.armor[0]}<p>Healing: ${data.healing[0]}<p>Mundane: ${data.mundane[0]}<p>Spell Effect: ${data.spell_effect[0]}`
        ));
};

document.addEventListener('keyup', function (e) {
    if (e.key === "Escape") {
        const modals = document.querySelectorAll('.modal-overlay');
        for (const modal of modals) {
            modal.click();
        }
    }
});

$("button#name-btn").click(getName);
$("button#regen-name").click(getName);

$("button#event-btn").click(getEvent);
$("button#regen-event").click(getEvent);

$("button#monument-btn").click(getMonument);
$("button#regen-monument").click(getMonument);

$("button#trap-btn").click(getTrap);
$("button#regen-trap").click(getTrap);

$("button#item-btn").click(getItem);
$("button#regen-item").click(getItem);