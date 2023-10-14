USE swapi_db;
SELECT * FROM tables;

show tables;

/* -------------------- */
/* CREATE TABLE PEOPLE */
/* -------------------- */
create table people (
	id INT(10) primary key auto_increment, 
    name VARCHAR(10), /**/
    mass INT(10),
    hair_color VARCHAR(10), /**/
    skin_color varchar(10), /**/
    eye_color varchar(10), /**/
    birth_year varchar(10), /**/
    gender enum('male','female') default NULL,
    created_at DATETIME /**/
);

ALTER TABLE people
	add column height INT(10) after name;
    
ALTER TABLE people
	MODIFY gender ENUM('male', 'female', 'no gender');

ALTER TABLE people
	MODIFY name VARCHAR(100);

ALTER TABLE people
	ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	MODIFY hair_color VARCHAR(100), /**/
    MODIFY skin_color varchar(100), /**/
    MODIFY eye_color varchar(100), /**/
    MODIFY birth_year varchar(100); /**/

select * from people;

insert into people(id, name, height, mass, hair_color, skin_color, eye_color, birth_year, gender)
values(1, 'Luke Skywalker', 172, 77, 'blond', 'fair', 'blue', '19BBY', 'male'),
	(2, 'C-390', 167, 75, 'n/a', 'gold', 'yellow', '112BBY', 'no gender'),
    (3, 'R2-D2', 96, 32, 'n/a', 'white and blue', 'red', '33BBY', 'no gender'),
    (4, 'Darth Vader', 202, 136, 'none', 'white', 'yellow', '41.9BBY', 'male'),
    (5, 'Leia Organa', 150, 49, 'brown', 'light', 'brown', '19BBY', 'female'),
    (6, 'Owen Lars', 178, 120, 'brown, grey', 'light', 'blue', '52BBY', 'male'),
    (7, 'Beru Whitesun lars', 165, 75, 'brown', 'light', 'blue', '47BBY', 'female'),
    (8, 'R5-D4', 97, 32, 'n/a', 'white, red', 'red', 'unknown', 'no gender'),
    (9, 'Biggs Darklighter', 183, 84, 'black', 'light', 'brown', '24BBY', 'male'),
    (10, 'Obi-Wan Kenobi', 182, 77, 'auburn, white', 'fair', 'blue-ray', '57BBY', 'male');

DESC user;

DROP TABLE peoples;

/* -------------------- */
/* CREATE TABLE PLANETS */
/* -------------------- */
create table planet(
	id INT(10) primary key auto_increment,
    rotation_period INT(10),
    orbital_period INT(10),
    diameter INT(10),
    climate VARCHAR(100),
    gravity VARCHAR(100),
    terrain VARCHAR(100),
    surface_water INT(10),
    population INT(100), /**/
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE planet
	ADD name VARCHAR(100) NOT NULL AFTER id;

ALTER TABLE planet
	MODIFY population BIGINT(100);

select * from planet;

insert into planet(id, name, rotation_period, orbital_period, diameter, climate, gravity, terrain, surface_water, population)
values(1, 'Tatooine', 23, 304, 10465, 'arid', '1 standard', 'desert', 1, '200000'),
	(2, 'Alderaan', 24, 364, 12500, 'temperate', '1 Standard', 'grasslands, mountains', 40, 2000000000),
    (3, 'Yavin IV', 24, 4818, 10200, 'temperate, tropical', '1 standard', 'jungle, rainforests', 8, 1000),
    (4, 'Hoth', 23, 549, 7200, 'Frozen', '1.1 Standard', 'tundra, ice caves, mountain ranges', 100, null),
    (5, 'Dagobah', 23, 341, 8900, 'Murky', 'n/a', 'Swamp, Jungles', 8, null),
    (6, 'Bespin', 12, 5110, 118000, 'temperate', '1.5 (surface), 1 standard (Cloud City)', 'gas giant', 0, 6000000),
    (7, 'Endor', 18, 402, 4900, 'temperate', '0.85 standard', 'forests, mountains, lakes', 8, 30000000),
    (8, 'Naboo', 26, 312, 12120, 'temperate', '1 Standards', 'grassy hills, swamps, forests, mountains', 12, 4500000000),
    (9, 'Coruscant', 24, 368, 12240, 'temparate', '1 standard', 'Cityscape, Mountains', null, 1000000000000),
    (10, 'Kamino', 27, 463, 19720, 'temperate', '1 standard', 'ocean', 100, 1000000000);
    
select * from planet;

/* -------------------- */
/* CREATE TABLE STARSHIPS */
/* -------------------- */
create table starships (
	id INT(10) PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
	model VARCHAR(100),
    manufacturer VARCHAR(100),
    cost_in_credits BIGINT(100),
    length FLOAT,
    max_atmosphering_speed INT(100),
    crew VARCHAR(100),
    passengers INT(100),
    cargo_capacity BIGINT(100),
	consumables VARCHAR(100),
    hyperdrive_rating FLOAT,
    MGLT INT(100),
    starship_class VARCHAR(100)
);

SELECT * FROM starships;

insert into starships(id, name, model, manufacturer, cost_in_credits, length, max_atmosphering_speed, crew, passengers, cargo_capacity, consumables, hyperdrive_rating, MGLT, starship_class)
values(1, 'CR90 Corvette', 'CR90 Corvette', 'Corellian Engineering Corporation', 3500000, 150, 950, '30-165', 600, 3000000, '1 year', 2.0, 60, 'corvette'),
	(2, 'Star Destroyer', 'Imperial I-class Star Destroyer', 'Kuat Drive Yards', 150000000, 1.600, 975, '47.060', null, 36000000, '2 years', 2.0, 60, 'Star Destroyer'),
    (3, 'Sentinel-class landing craft', 'Sentinel-class landing craft', 'Sienar Fleet Systems, Cyngus Spaceworks', 240000, 38, 1000, 5, 75, 180000, '1 month', 1.0, 70, 'landing craft'),
    (4, 'Death Star', 'DS-1 Orbital Battle Station', 'Imperial Department of Military Research, Sienar Fleet Systems', 1000000000000, 120000, null, '342.953', 843342, 1000000000000, '3 years', 4.0, 10, 'Deep Space Mobile Battlestation'),
    (5, 'Millennium Falcon', 'YT-1300 light freighter', 'Corellian Engineering Corporation', 100000, 34.37, 1050, '4', 6, 100000, '2 months', 0.5, 75, 'Light freighter');


    
    
    
    
    
    
    
    
