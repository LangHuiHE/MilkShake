create table paymentMethod
(
    typeID     int auto_increment
        primary key,
    definition varchar(20) null
);

INSERT INTO dbo.paymentMethod (typeID, definition) VALUES (1, 'Cash');
INSERT INTO dbo.paymentMethod (typeID, definition) VALUES (2, 'Check');
INSERT INTO dbo.paymentMethod (typeID, definition) VALUES (3, 'E-Check');
INSERT INTO dbo.paymentMethod (typeID, definition) VALUES (4, 'Credit Card');
INSERT INTO dbo.paymentMethod (typeID, definition) VALUES (5, 'Debit');
INSERT INTO dbo.paymentMethod (typeID, definition) VALUES (6, 'Scholarship');
