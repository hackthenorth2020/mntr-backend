--Select user
----select profile
SELECT * FROM profiles WHERE uid = $1;
----select edu
SELECT * FROM education WHERE uid = $1;
----select jobs
SELECT * FROM jobs WHERE uid = $1;
----select links
SELECT * FROM links WHERE uid = $1;
----select schedule
SELECT * FROM schedule WHERE uid = $1;

--create user
----create profile
INSERT INTO profiles (uid, fname, lname, email, bday, interests, bio)
VALUES ($1, $2, $3, $4, $5, $6, $7);
----create edu
INSERT INTO education (uid, school, start_date, end_date, major, city, country)
VALUES ($1, $2, $3, $4, $5, $6, $7);
----create jobs
INSERT INTO jobs (uid, company, start_year, end_year, position, city, country)
VALUES ($1, $2, $3, $4, $5, $6, $7); 
----create links
INSERT INTO links (uid, linkedin, twitter, facebook, github, other)
VALUES ($1, $2, $3, $4, $5, $6);
----create schedule?
INSERT INTO schedules (uid, monday, tuesday, wednesday, thursday, friday, saturday, sunday)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

--update user
----update profile
UPDATE profiles SET (fname, lname, email, bday, interests)
= ($2, $3, $4, $5, $6)
WHERE uid = $1;
----update edu
UPDATE education SET (school, start, end, major, city, country)
= ($2, $3, $4, $5, $6, $7)
WHERE uid = $1 AND start = $2 AND major = $4;
----update jobs
UPDATE jobs SET (company, start, end, position, city, country)
= ($2, $3, $4, $5, $6, $7)
WHERE uid = $1 AND company = $2 AND start = $3;
----update links
UPDATE links SET (linkedin, twitter, facebook, github, other)
= ($2, $3, $4, $5, $6)
WHERE uid = $1;
----update schedule?
UPDATE schedules SET (monday, tuesday, wednesday, thursday, friday, saturday, sunday)
= ($2, $3, $4, $5, $6, $7, $8)
WHERE uid = $1;

--DELETE
----update edu
DELETE FROM education WHERE uid = $1, school = $2, major = $3;
----update jobs
DELETE FROM jobs WHERE uid = $1, company = $2, position = $3;

--PAIRING
----select pairing
SELECT * FROM pairings WHERE mentor_id = $1 OR mentee_id = $2;
----create pairing
INSERT INTO pairings VALUES ($1, $2);
----delete pairing
DELETE FROM pairings WHERE mentor_id = $1 AND mentee_id = $2;

--POINTS
----update points
UPDATE points SET points = $2 WHERE id = $1

--MOST SIMILAIR INTERESTS
with c as (SELECT uid, count(1) as common
      FROM (SELECT uid, unnest(profiles.interests) as intr
      FROM profiles)
x WHERE intr = any(SELECT unnest(interests) FROM profiles WHERE uid = $1) AND uid <> $1
GROUP BY uid) 
SELECT * FROM c WHERE common = (select max(common) from c);