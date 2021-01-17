package profiles

const (
	SELECT_PROFILE      = "SELECT * FROM profiles WHERE uid = $1;"
	SELECT_ALL_PROFILES = "SELECT * FROM profiles LIMIT 50"
	SELECT_EDUCATION    = "SELECT * FROM education WHERE uid = $1;"
	SELECT_JOBS         = "SELECT * FROM jobs WHERE uid = $1;"
	SELECT_LINKS        = "SELECT * FROM links WHERE uid = $1;"
	SELECT_SCHEDULE     = "SELECT * FROM schedule WHERE uid = $1;"

	CREATE_PROFILE   = "INSERT INTO profiles (uid, fname, lname, email, bday, interests, bio) VALUES ($1, $2, $3, $4, $5, $6, $7);"
	CREATE_EDUCATION = "INSERT INTO education (uid, school, start_date, end_date, major, city, country)	VALUES ($1, $2, $3, $4, $5, $6, $7);"
	CREATE_JOBS      = "INSERT INTO jobs (uid, company, start_date, end_date, position, city, country) VALUES ($1, $2, $3, $4, $5, $6, $7);"
	CREATE_LINKS     = "INSERT INTO links (uid, linkedin, twitter, facebook, github, other)	VALUES ($1, $2, $3, $4, $5, $6);"
	CREATE_SCHEDULES = "INSERT INTO schedules (uid, monday, tuesday, wednesday, thursday, friday, saturday, sunday)	VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"

	UPDATE_PROFILE   = "UPDATE profiles SET (fname, lname, email, bday, interests, bio) = ($2, $3, $4, $5, $6, $7) WHERE uid = $1;"
	UPDATE_EDUCATION = "UPDATE education SET (school, start_date, end_date, major, city, country) = ($2, $3, $4, $5, $6, $7) WHERE uid = $1;"
	UPDATE_JOBS      = "UPDATE jobs SET (company, start_date, end_date, position, city, country)	= ($2, $3, $4, $5, $6, $7)	WHERE uid = $1, company = $2, start_date = $3;"
	UPDATE_LINKS     = "UPDATE links SET (linkedin, twitter, facebook, github, other)	= ($2, $3, $4, $5, $6) WHERE uid = $1;"
	UPDATE_SCHEDULES = "UPDATE schedules SET (monday, tuesday, wednesday, thursday, friday, saturday, sunday) = ($2, $3, $4, $5, $6, $7, $8) WHERE uid = $1;"

	DELETE_EDUCATION = "DELETE FROM education WHERE uid = $1 AND school = $2 AND major = $3;"
	DELETE_JOBS      = "DELETE FROM jobs WHERE uid = $1 AND company = $2 AND position = $3;"

	SELECT_PAIRING        = "SELECT * FROM pairings WHERE (mentor_id = $1 OR mentee_id = $2) AND status <> 0;"
	SELECT_MENTORS        = "SELECT * FROM pairings WHERE mentee_id = $1 AND status <> 0;"
	SELECT_MENTEES        = "SELECT * FROM pairings WHERE mentor_id = $1 AND status <> 0;"
	SELECT_MENTOR_REQS    = "SELECT uid, fname, lname, email, bday, interests, bio FROM (SELECT mentee_id FROM pairings WHERE mentor_id = $1 AND status = 1) JOIN profiles on mentee_id=uid;"
	UPDATE_PAIR_MATCHED   = "UPDATE pairings SET status = 2 WHERE mentor_id = $1 AND mentee_id = $2"
	UPDATE_PAIR_UNMATCHED = "UPDATE pairings SET status = 0 WHERE mentor_id = $1 AND mentee_id = $2"
	INSERT_PAIRING        = "INSERT INTO pairings  VALUES ($1, $2, 1);"
	DELETE_PAIRING        = "DELETE FROM pairings WHERE mentor_id = $1 AND mentee_id = $2;"

	UPDATE_POINTS = "UPDATE points SET points = points + $2 WHERE id = $1"

	GET_MOST_SIMILAR_INTERESTS = "with c as (SELECT uid, count(1) as common FROM (SELECT uid, unnest(profiles.interests) as intr FROM profiles) " +
		"x WHERE intr = any(SELECT unnest(interests) FROM profiles WHERE uid = $1) AND uid <> $1	GROUP BY uid) " +
		"SELECT * FROM c WHERE common = (select max(common) from c);"

	GET_LIMIT_SIMILAR_INTERESTS = "with cte as (SELECT uid, count(1) as common " +
		"FROM (SELECT uid, unnest(profiles.interests) as intr " +
		"FROM profiles) " +
		"x WHERE intr = any(SELECT unnest(interests) FROM profiles WHERE uid = $1) AND uid <> $1 " +
		"GROUP BY uid)  " +
		"SELECT cte.uid, fname, lname, email, bday, interests, bio FROM cte INNER JOIN profiles p on cte.uid = p.uid ORDER BY common DESC LIMIT $2;"

	GET_MESSAGES = "SELECT * FROM messages WHERE (to = $2 AND from = $1) OR (to = $1 AND from = $2)"

	INSERT_MESSAGE = "INSERT INTO MESSAGES (from, to, message, time) VALUES ($1, $2, $3, $4)"
)
