CREATE TABLE `notifications` (
  uuid    CHAR(36) NOT NULL PRIMARY KEY,
  content JSON     NOT NULL,
  created DATETIME NOT NULL
) ENGINE = InnoDB DEFAULT CHAR SET = "utf8";
