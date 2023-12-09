CREATE TABLE IF NOT EXISTS Course (
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

CREATE TABLE IF NOT EXISTS ContributorTeam (
    ID BIGINT PRIMARY KEY,
    -- Add other fields for ContributorTeam as needed
);

-- Assuming you are using PostgreSQL or a similar database that supports UUID
-- If you are using MySQL, you might want to use AUTO_INCREMENT instead
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

ALTER TABLE Course ADD COLUMN IF NOT EXISTS ContributorTeamID UUID DEFAULT uuid_generate_v4();

-- Add foreign key constraint
ALTER TABLE Course
    ADD CONSTRAINT FK_ContributorTeam
        FOREIGN KEY (ContributorTeamID)
            REFERENCES ContributorTeam(ID);
