DROP VIEW ocenka;
CREATE VIEW ocenka
AS SELECT name, lesson, item
FROM exam 
WHERE ITEM <4 OR item IS NULL;
SELECT *
FROM ocenka
