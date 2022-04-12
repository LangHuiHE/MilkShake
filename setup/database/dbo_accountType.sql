create table accountType
(
    typeID     int auto_increment
        primary key,
    definition varchar(20) null
);

INSERT INTO dbo.accountType (typeID, definition) VALUES (1, 'Tuition');
INSERT INTO dbo.accountType (typeID, definition) VALUES (2, 'Meal Plan');
INSERT INTO dbo.accountType (typeID, definition) VALUES (3, 'PrintPoint');
