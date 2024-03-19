-- CREATE TABLE PARENT
CREATE TABLE tbl_parent (
    id INT PRIMARY KEY,
    name VARCHAR(255)
);

-- CREATE TABLE CHILD
CREATE TABLE tbl_child (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    parent_id INT,
    FOREIGN KEY (parent_id) REFERENCES tbl_parent(id)
);