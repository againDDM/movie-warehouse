
BEGIN;

CREATE TABLE films (
    id SERIAL,
    name VARCHAR(64) NOT NULL,
    description TEXT NULL DEFAULT NULL,
    PRIMARY KEY (id)
);
INSERT INTO films (name, description) VALUES
    ('Pulp Fiction', 'Pulp Fiction is a 1994 American crime film written and directed by Quentin Tarantino, who conceived it with Roger Avary.[5] Starring John Travolta, Samuel L. Jackson, Bruce Willis, Tim Roth, Ving Rhames, and Uma Thurman, it tells several stories of criminal Los Angeles. The title refers to the pulp magazines and hardboiled crime novels popular during the mid-20th century, known for their graphic violence and punchy dialogue.'),
    ('Full metal jacket', 'Full Metal Jacket is a 1987 war film directed, co-written, and produced by Stanley Kubrick and starring Matthew Modine, R. Lee Ermey, Vincent D`Onofrio and Adam Baldwin. The screenplay by Kubrick, Michael Herr, and Gustav Hasford was based on Hasford`s novel The Short-Timers (1979). The storyline follows a platoon of U.S. Marines through their boot camp training in Marine Corps Recruit Depot Parris Island, South Carolina, primarily focusing on two privates, Joker and Pyle, who struggle under their abusive drill instructor, Gunnery Sergeant Hartman, and the experiences of two of the platoon`s Marines in Vietnamese cities of Da Nang and Huế during the Tet Offensive of the Vietnam War.'),
    ('RocknRolla', 'RocknRolla is a 2008 black comedy crime film written and directed by Guy Ritchie, and starring Gerard Butler, Tom Wilkinson, Thandie Newton, Mark Strong, Idris Elba, Tom Hardy, Gemma Arterton and Toby Kebbell. It was released on 5 September 2008 in the United Kingdom, hitting number one in the UK box office in its first week of release'),
    ('From Dusk till Dawn', 'The bar employees reveal themselves as vampires and kill most of the patrons. Richie is bitten by a stripper and dies; only Seth, Jacob, Kate, Scott, a biker named Sex Machine, and Frost, a Vietnam veteran, survive. The others are reborn as vampires, including Richie, forcing the survivors to kill them all. When an army of vampires, in bat form, assembles outside, the survivors lock themselves in, but Sex Machine is bitten, becomes a vampire, and bites Frost and Jacob. Frost throws Sex Machine through the door, allowing the vampires to enter while Frost turns into a vampire.'),
    ('Snatch', 'Snatch is a 2000 British crime comedy film written and directed by Guy Ritchie, featuring an ensemble cast. Set in the London criminal underworld, the film contains two intertwined plots: one dealing with the search for a stolen diamond, the other with a small-time boxing promoter (Jason Statham) who finds himself under the thumb of a ruthless gangster (Alan Ford) who is ready and willing to have his subordinates carry out severe and sadistic acts of violence.'),
    ('Sin City', 'Much of the film is based on the first, third, and fourth books in Miller`s original comic series. The Hard Goodbye is about a man who embarks on a brutal rampage in search of his one-time sweetheart`s killer, killing anyone, even the police, that gets in his way of finding and killing her murderer. The Big Fat Kill focuses on an everyman getting caught in a street war between a group of prostitutes and a group of mercenaries, the police and the mob. That Yellow Bastard follows an aging police officer who protects a young woman from a grotesquely disfigured serial killer. The intro and outro of the film are based on the short story "The Customer is Always Right" which is collected in Booze, Broads & Bullets, the sixth book in the comic series.')
;

COMMIT;
