create table eventClock
(
    eventClockID   int auto_increment
        primary key,
    eventID        int      null,
    userID         int      null,
    hourType       int      null,
    startTimeStamp datetime null,
    endTimeStamp   datetime null,
    note           text     null
);

INSERT INTO dbo.eventClock (eventClockID, eventID, userID, hourType, startTimeStamp, endTimeStamp, note) VALUES (13, 7963417, 394203, 0, '2022-04-11 19:55:34', null, '');
INSERT INTO dbo.eventClock (eventClockID, eventID, userID, hourType, startTimeStamp, endTimeStamp, note) VALUES (17, 7962622, 394203, 0, '2022-04-11 23:57:15', null, '');
INSERT INTO dbo.eventClock (eventClockID, eventID, userID, hourType, startTimeStamp, endTimeStamp, note) VALUES (18, 7962622, 12345678, 0, '2022-04-12 02:04:47', null, '');
