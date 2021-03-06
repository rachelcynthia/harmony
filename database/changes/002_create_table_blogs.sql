CREATE TABLE BLOGS (
    ID SERIAL PRIMARY KEY,
    PUBLISHED_DATE TEXT NOT NULL,
    TITLE TEXT NOT NULL,
    CONTENT TEXT NOT NULL,
    PUBLIC BOOLEAN NOT NULL,
    GROUPS TEXT,
    USER_ID INT NOT NULL,
    CONSTRAINT fk_Blogs  
    FOREIGN KEY(USER_ID)   
    REFERENCES USERS(ID)  
);