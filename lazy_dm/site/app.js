var Api = 'http://localhost:2020'

function getName() {
    namePlace = document.getElementById('name');
    fetch(`${Api}/v1/names`)
        .then(response => response.json())
        .then(data => namePlace.textContent = `${data.given} ${data.surname}`);
};

function getEvent() {
    eventPlace = document.getElementById('event');
    fetch(`${Api}/v1/events`)
        .then(response => response.json())
        .then(data => eventPlace.innerHTML = (
            `<p>Mundane Event: ${data.mundane}<p>Town Sentiment: ${data.sentiment}<p>Notable Weather: ${data.weather}<p>Fantastic Event: ${data.fantastic}`
        ));
};

function getMonument() {
    namePlace = document.getElementById('monument');
    fetch(`${Api}/v1/monuments`)
        .then(response => response.json())
        .then(data => namePlace.innerHTML =
            `${data.condition} ${data.origin} ${data.type} with a ${data.effect} effect`);
};

function getTrap() {
    eventPlace = document.getElementById('trap');
    fetch(`${Api}/v1/traps`)
        .then(response => response.json())
        .then(data => eventPlace.innerHTML = (
            `<p>Type: ${data.type}<p>Flavor: ${data.flavor}<p>Trigger: ${data.trigger}`
        ));
};

function getItem() {
    eventPlace = document.getElementById('item');
    fetch(`${Api}/v1/items`)
        .then(response => response.json())
        .then(data => eventPlace.innerHTML = (
            `<p>Origin: ${data.origin}<p>Condition: ${data.condition}<p>Weapon: ${data.weapon}<p>Armor: ${data.armor}<p>Healing: ${data.healing}<p>Mundane: ${data.mundane}<p>Spell Effect: ${data.spell_effect}`
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