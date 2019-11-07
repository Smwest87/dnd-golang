# dnd-golang


DND

CREATE SCHEMA IF NOT EXISTS dnd;

CREATE TABLE IF NOT EXISTS dnd.dnd_characters
(
id SERIAL PRIMARY KEY,
name VARCHAR NOT NULL,
class VARCHAR NOT NULL,
level int NOT NULL,
hitpointmaximum int NOT NULL,
strength int NOT NULL,
dexterity int NOT NULL,
constitution int NOT NULL,
wisdom int NOT NULL,
intelligence int NOT NULL,
charisma int NOT NULL,
initiative int NOT NULL,
modifiers int[] NOT NULL
);

CREATE TABLE IF NOT EXISTS dnd.races
(
id serial Primary key NOT NULL,
name text NOT NULL,
strength int NOT NULL,
dexterity int NOT NULL,
constitution int NOT NULL,
wisdom int NOT NULL,
intelligence int NOT NULL,
charisma int NOT NULL,
size text NOT NULL,
speed int NOT NULL,
languages TEXT [],
traits TEXT []
);

INSERT INTO dnd.races
(name, strength, dexterity, constitution, wisdom, intelligence, charisma, size, speed,languages,traits)
VALUES
('Gray Dwarf',1,0,2,0,0,0, 'Medium', 25, ARRAY ['Common','Dwarvish'], ARRAY ['120 Darkvision', 'Dwarven Combat Training', 'Stone-cunning', 'Adv. on poison saves', 'Poison Res.', 'Smiths/Brewers/Masons tools prof.', 'Speed not reduced by heavy armor', 'Sunlight Sensitivity','Duergar Magic','Illusion,charm, paralysis resistance']);

SELECT * FROM dnd.races;

DELETE FROM dnd.races 
WHERE id is not NULL;

DROP TABLE dnd.races;


SELECT * FROM dnd.dnd_characters;

DELETE FROM dnd.dnd_characters;

DROP TABLE dnd.dnd_characters;