
-- work command.
-- source /home/anxiu/project/campus-hub/campus-hub-server/deploy/sql/init.sql
--

CREATE TABLE IF NOT EXISTS course (
    ID BIGINT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    ContributorTeamID BIGINT, -- Assuming ContributorTeam is a separate table with ID as primary key
    Discipline VARCHAR(255),
    License VARCHAR(255),
    Origination VARCHAR(255),
    Version VARCHAR(255),
    CreatedAt BIGINT,
    UpdatedAt BIGINT,
    DeletedAt BIGINT,
    ResourceAddr VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS contributorTeam (
    ID BIGINT PRIMARY KEY,
    Name VARCHAR(255)
    -- Add other fields for ContributorTeam as needed
);


-- Insert data into ContributorTeam
INSERT INTO contributorTeam (ID, Name) VALUES
    (1, 'Team1');

-- -- Insert data into Contributor
-- INSERT INTO Contributor (ID, Name, Org, Role) VALUES
--     (1, 'Contributor1', 'Org1', 'Instructor'),
--     (2, 'Contributor2', 'Org2', 'TeachingAssistant'),
--     (3, 'Contributor3', 'Org3', 'CommunityContributor');

-- Insert data into Course
-- MySQL Version, only work in MySQL since use the variable feature.
SET @timenow := UNIX_TIMESTAMP(NOW());
INSERT INTO course (ID, Name, ContributorTeamID, Discipline, License, Origination, Version, CreatedAt, UpdatedAt, DeletedAt, ResourceAddr) VALUES
    (1, 'Course1', 1, 'Discipline1', 'MIT', 'Origination1', '23 fall', @timenow, @timenow, NULL, 'cdn.campus-hub.online/resource/course/1'),
    (2, 'Course2', 1, 'Discipline2', 'Apache 2.0', 'Origination2', '24 spring', @timenow, @timenow, NULL, 'cdn.campus-hub.online/resource/course/2'),
    (3, 'Course3', 1, 'Discipline3', 'MIT', 'Origination3', '1.0', @timenow, @timenow, NULL, 'cdn.campus-hub.online/resource/course/3');

-- Insert data into Course
-- PostgreSQL Version, use DO block.
-- -- Assuming you are using PostgreSQL syntax with $1, $2, etc. as placeholders
-- -- You may need to adjust syntax based on the database system you are using
--
-- -- Declare a variable to hold the current time
-- DO $$
-- DECLARE
-- timenow BIGINT;
-- BEGIN
--     timenow := EXTRACT(EPOCH FROM NOW());
--
--     -- Insert data into Course
-- INSERT INTO Course (ID, Name, ContributorTeamID, Discipline, License, Origination, Version, CreatedAt, UpdatedAt)
-- VALUES
--     (1, 'Course1', 1, 'Discipline1', 'License1', 'Origination1', 'Version1', timenow, timenow),
--     (2, 'Course2', 1, 'Discipline2', 'License2', 'Origination2', 'Version2', timenow, timenow),
--     (3, 'Course3', 1, 'Discipline3', 'License3', 'Origination3', 'Version3', timenow, timenow);
-- END $$;


-- only for postgreSQL uuid create feature.
-- -- Assuming you are using PostgreSQL or a similar database that supports UUID
-- -- If you are using MySQL, you might want to use AUTO_INCREMENT instead
-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
--
-- ALTER TABLE Course ADD COLUMN IF NOT EXISTS ContributorTeamID UUID DEFAULT uuid_generate_v4();
--
-- -- Add foreign key constraint
-- ALTER TABLE Course
--     ADD CONSTRAINT FK_ContributorTeam
--         FOREIGN KEY (ContributorTeamID)
--             REFERENCES ContributorTeam(ID);
