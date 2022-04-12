create table userType
(
    typeID     int auto_increment
        primary key,
    definition varchar(20) null
);

INSERT INTO dbo.userType (typeID, definition) VALUES (1, 'Admin');
INSERT INTO dbo.userType (typeID, definition) VALUES (2, 'Faculty');
INSERT INTO dbo.userType (typeID, definition) VALUES (3, 'Student');
INSERT INTO dbo.userType (typeID, definition) VALUES (4, 'Visitor');
