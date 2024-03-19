-- INSERT TABLE PARENT
INSERT INTO tbl_parent (id, name) VALUES
(1, 'Wahyu'),
(2, 'Ilham'),
(3, 'Irwan');

-- INSERT TABLE CHILD
INSERT INTO tbl_child (id, name, parent_id) VALUES
(1, 'Zaki', 2),
(2, 'Ilham', NULL),
(3, 'Irwan', 2),
(4, 'Arka', 3);

-- SELECT QUERY
SELECT C.id, C.name, P.name AS parent_name
FROM tbl_child C
LEFT JOIN tbl_parent P ON C.parent_id = P.id