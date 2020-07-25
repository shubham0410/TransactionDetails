
-- create the database named test
CREATE DATABASE test;

-- query to create table transaction. ParentId is the foreign key to self table
CREATE TABLE `transaction` (
  `id` bigint(20) NOT NULL,
  `amount` double DEFAULT NULL,
  `type` varchar(90) DEFAULT NULL,
  `parentId` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_transaction_1_idx` (`parentId`),
  CONSTRAINT `fk_transaction_1` FOREIGN KEY (`parentId`) REFERENCES `transaction` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


--/procedure to get sum of parents of the children. 
DELIMITER $$
CREATE PROCEDURE calctotal(
   IN number INT,
   OUT total INT
)
BEGIN
   DECLARE parent_ID INT DEFAULT NULL ;
   DECLARE tmptotal INT DEFAULT 0;
   DECLARE tmptotal2 INT DEFAULT 0;

   SELECT id   FROM transaction   WHERE parentId = number INTO parent_ID;   
   SELECT amount   FROM transaction   WHERE id = number INTO tmptotal;     

   IF parent_ID IS NULL
    THEN
    SET total = tmptotal;
   ELSE     
    CALL calctotal(parent_ID, tmptotal2);
    SET total = tmptotal2 + tmptotal;   
   END IF;
END$$
DELIMITER ;
