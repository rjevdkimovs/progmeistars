DROP VIEW from10to12 ;
CREATE VIEW from10to12
AS SELECT NAME, lesson, item
FROM lessons
WHERE item <4 OR IS NULL;
SELECT *
FROM from10to12
