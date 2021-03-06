CREATE TABLE GROUPS (
    ID SERIAL PRIMARY KEY,
    NAME TEXT NOT NULL,
    DATE_CREATED DATE NOT NULL,
    ADMIN_NAME TEXT NOT NULL,
    USER_ID INT NOT NULL,
    CONSTRAINT fk_Favourites  
    FOREIGN KEY(USER_ID)   
    REFERENCES USERS(ID)  
);