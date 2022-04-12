create table payment
(
    paymentID       int auto_increment
        primary key,
    userID          int                      not null,
    totalBalance    float                    null,
    paymentPlanType int                      null,
    semester        int                      null,
    accountType     int                      null,
    processDate     datetime                 null,
    paymentMethod   int                      null,
    location        varchar(30) charset utf8 null
);

INSERT INTO dbo.payment (paymentID, userID, totalBalance, paymentPlanType, semester, accountType, processDate, paymentMethod, location) VALUES (1, 12345678, -100, null, 13, 1, '2022-02-19 17:04:41', null, 'University');
INSERT INTO dbo.payment (paymentID, userID, totalBalance, paymentPlanType, semester, accountType, processDate, paymentMethod, location) VALUES (2, 12345678, 70, null, 13, 1, '2022-02-19 17:38:28', 1, 'Cashier');
INSERT INTO dbo.payment (paymentID, userID, totalBalance, paymentPlanType, semester, accountType, processDate, paymentMethod, location) VALUES (3, 12345678, 200, null, 13, 2, '2022-02-19 18:39:37', 2, 'Dinning Service');
INSERT INTO dbo.payment (paymentID, userID, totalBalance, paymentPlanType, semester, accountType, processDate, paymentMethod, location) VALUES (4, 12345678, -2, null, 13, 3, '2022-02-19 18:40:00', null, 'Library');
INSERT INTO dbo.payment (paymentID, userID, totalBalance, paymentPlanType, semester, accountType, processDate, paymentMethod, location) VALUES (5, 12345678, 2.3, null, 13, 1, '2022-02-01 18:36:00', 3, 'Mobile');
INSERT INTO dbo.payment (paymentID, userID, totalBalance, paymentPlanType, semester, accountType, processDate, paymentMethod, location) VALUES (6, 12345678, 5.01, null, 11, 1, '2021-02-20 18:37:43', 3, 'Mobile');
INSERT INTO dbo.payment (paymentID, userID, totalBalance, paymentPlanType, semester, accountType, processDate, paymentMethod, location) VALUES (7, 12345678, 3.1, null, 13, 2, '2022-02-19 18:38:34', 2, 'Dinning Service');
INSERT INTO dbo.payment (paymentID, userID, totalBalance, paymentPlanType, semester, accountType, processDate, paymentMethod, location) VALUES (8, 12345678, 5.6, null, 13, 2, '2022-02-19 18:38:44', 2, 'Dinning Service');
