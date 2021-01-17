create table profiles(
    uid text primary key, 
    fname text not null,
    lname text not null,
    email text not null,
    bday text not null,
    interests text[],
    bio text
);

create table education (
    uid text,
    school text not null,
    start_date text not null,
    end_date text not null,
    major text,
    city text,
    country text,
    PRIMARY KEY (uid, school, major),
    CONSTRAINT fk_uid FOREIGN KEY (uid) REFERENCES profiles(uid)
    ON DELETE CASCADE
);

create table jobs (
    uid text,
    company text not null,
    start_date text not null,
    end_date text,
    position text not null,
    city text,
    country text,
    PRIMARY KEY (uid, company, start_date),
    CONSTRAINT fk_uid FOREIGN KEY (uid) REFERENCES profiles(uid)
    ON DELETE CASCADE
);

create table links (
    uid text primary key,
    linkedin text,
    twitter text,
    facebook text,
    github text,
    other text,
    CONSTRAINT fk_uid FOREIGN KEY (uid) REFERENCES profiles(uid)
    ON DELETE CASCADE
);

create table pairings (
    mentor_id text not null,
    mentee_id text not null,
    status int not null,
    PRIMARY KEY (mentor_id, mentee_id),
    CONSTRAINT not_self check (mentor_id <> mentee_id),
    CONSTRAINT fk_mentor FOREIGN KEY (mentor_id) REFERENCES profiles(uid),
    CONSTRAINT fk_mentee FOREIGN KEY (mentee_id) REFERENCES profiles(uid)
    ON DELETE CASCADE
);

create table schedules (
    uid text primary key,
    monday text,
    tuesday text,
    wednesday text,
    thursday text,
    friday text,
    saturday text,
    sunday text,
    CONSTRAINT fk_uid FOREIGN KEY (uid) REFERENCES profiles(uid)
    ON DELETE CASCADE
);

CREATE TABLE points (
    uid text primary key,
    points int,
    CONSTRAINT fk_uid FOREIGN KEY (uid) REFERENCES profiles(uid)
    ON DELETE CASCADE
);

CREATE TABLE messages (
    uuid UUID PRIMARY KEY,
    from text,
    to text,
    message text,
    time text
    CONSTRAINT fk_from_uid FOREIGN KEY (from) REFERENCES profiles(uid),
    CONSTRAINT fk_from_uid FOREIGN KEY (to) REFERENCES profiles(uid)
)