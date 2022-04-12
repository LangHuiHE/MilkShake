create table eventHourType
(
    typeID     int auto_increment
        primary key,
    definition varchar(20) null
);

INSERT INTO dbo.eventHourType (typeID, definition) VALUES (1, 'General');
INSERT INTO dbo.eventHourType (typeID, definition) VALUES (2, 'Service');
INSERT INTO dbo.eventHourType (typeID, definition) VALUES (3, 'Participate');
