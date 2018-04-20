CREATE ROLE timescale;
CREATE TABLE people (id serial NOT NULL PRIMARY KEY, name text);
CREATE TABLE users (username text PRIMARY KEY, password text NOT NULL);
CREATE TABLE projects (id serial NOT NULL PRIMARY KEY, name text, weeks real, person integer, FOREIGN KEY (person) REFERENCES people(id));
GRANT ALL ON people TO timescale;
GRANT ALL ON people_id_seq TO timescale;
GRANT ALL ON projects TO timescale;
GRANT ALL ON projects_id_seq TO timescale;
GRANT ALL ON users TO timescale;
