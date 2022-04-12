create table paymentRuns
(
    paymentRunID  int auto_increment
        primary key,
    paymentID     int      null,
    userID        int      null,
    amount        float    null,
    paymentMethod int      null,
    processDate   datetime null,
    constraint paymentRuns_paymentMethod_typeID_fk
        foreign key (paymentMethod) references paymentMethod (typeID),
    constraint paymentRuns_payment_paymentID_fk
        foreign key (paymentID) references payment (paymentID)
);

INSERT INTO dbo.paymentRuns (paymentRunID, paymentID, userID, amount, paymentMethod, processDate) VALUES (1, 1, 394300, 100, 6, '2022-01-21 15:10:06');
