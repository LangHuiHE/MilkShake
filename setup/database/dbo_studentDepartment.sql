create table studentDepartment
(
    typeID     int auto_increment
        primary key,
    definition varchar(20) null
);

INSERT INTO dbo.studentDepartment (typeID, definition) VALUES (1, 'Undefined');
INSERT INTO dbo.studentDepartment (typeID, definition) VALUES (2, 'Accounting');
INSERT INTO dbo.studentDepartment (typeID, definition) VALUES (3, 'Art');
INSERT INTO dbo.studentDepartment (typeID, definition) VALUES (4, 'Criminal Justice');
INSERT INTO dbo.studentDepartment (typeID, definition) VALUES (5, 'Communication');
INSERT INTO dbo.studentDepartment (typeID, definition) VALUES (6, 'Computing and Design');
INSERT INTO dbo.studentDepartment (typeID, definition) VALUES (7, 'Dental Hygiene');
INSERT INTO dbo.studentDepartment (typeID, definition) VALUES (8, 'Education');
INSERT INTO dbo.studentDepartment (typeID, definition) VALUES (9, 'Engineering');
INSERT INTO dbo.studentDepartment (typeID, definition) VALUES (10, 'English');
INSERT INTO dbo.studentDepartment (typeID, definition) VALUES (11, 'Mathematics');
INSERT INTO dbo.studentDepartment (typeID, definition) VALUES (12, 'Media Studies');
