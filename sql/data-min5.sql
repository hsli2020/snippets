insert into data_min5 (ts, val) values
('2020-09-01 00:00:00', 5),
('2020-09-01 00:05:00', 1),
('2020-09-01 00:10:00', 1),
('2020-09-01 00:15:00', 1),
('2020-09-01 00:20:00', 2),
('2020-09-01 00:25:00', 2),
('2020-09-01 00:30:00', 2),
('2020-09-01 00:35:00', 3),
('2020-09-01 00:40:00', 3),
('2020-09-01 00:45:00', 3),
('2020-09-01 00:50:00', 4),
('2020-09-01 00:55:00', 4),
('2020-09-01 01:00:00', 4);

-- 00:00:00 as start of min15
SELECT ts, COUNT(*), AVG(val)
FROM data_min5
GROUP BY UNIX_TIMESTAMP(ts) DIV 900


-- 00:00:00 as end of min15
SELECT ts, COUNT(*), AVG(val)
FROM data_min5
GROUP BY (UNIX_TIMESTAMP(ts)-1) DIV 900

--
SELECT FROM_UNIXTIME(((UNIX_TIMESTAMP(ts)-1) DIV 900)*900 + 900) AS tt, COUNT(*), AVG(val)
FROM data_min5
GROUP BY (UNIX_TIMESTAMP(ts)-1) DIV 900;
