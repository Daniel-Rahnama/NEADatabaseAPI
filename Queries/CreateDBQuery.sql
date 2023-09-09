CREATE DATABASE NEAGameDB;

CREATE TABLE Leaderboard (
	Username varchar(255) NOT NULL,
    Time1 int,
    Time2 int,
    Time3 int,
    PRIMARY KEY (Username)
);