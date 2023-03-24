CREATE TABLE IF NOT EXISTS Todos  (
    ID INT NOT NULL UNIQUE AUTO_INCREMENT,
    Title VARCHAR (255),
    Description VARCHAR (255),
    UserID INT,
    FOREIGN KEY (UserID) REFERENCES Users(ID),
    PRIMARY KEY (ID)
)