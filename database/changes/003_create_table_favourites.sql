CREATE TABLE FAVOURITES (
    ID SERIAL PRIMARY KEY,
    LIST TEXT NOT NULL,
    USER_ID INT NOT NULL,
    CONSTRAINT fk_Favourites  
    FOREIGN KEY(USER_ID)   
    REFERENCES USERS(ID)  
);