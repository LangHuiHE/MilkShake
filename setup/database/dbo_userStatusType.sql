create table userStatusType
(
    typeID     int auto_increment
        primary key,
    definition varchar(20) null
);

INSERT INTO dbo.userStatusType (typeID, definition) VALUES (1, 'Active');
INSERT INTO dbo.userStatusType (typeID, definition) VALUES (2, 'Suspend');
INSERT INTO dbo.userStatusType (typeID, definition) VALUES (3, 'Terminate');
