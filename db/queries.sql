CREATE TABLE users (
user_id INT AUTO_INCREMENT PRIMARY KEY,
username VARCHAR(40) NOT NULL UNIQUE KEY,
password VARCHAR(40) NOT NULL,
email VARCHAR(50) NOT NULL,
firstname VARCHAR(50) NOT NULL,
lastname VARCHAR(50)
);

CREATE TABLE tweets (
tweet_id INT AUTO_INCREMENT PRIMARY KEY,
tweet TEXT,
publication_date TIMESTAMP,
user_id INT
);

CREATE TABLE followers (
user_id INT,
follower_id INT
);

# Index helps in performance
CREATE INDEX tweet_index ON tweets (user_id);
CREATE INDEX user_index ON followers (user_id);
CREATE INDEX follower_index ON followers (follower_id);


SELECT username, firstname, lastname, tweet, publication_date
FROM users, tweets where users.user_id = tweets.user_id AND users.user_id in  (SELECT follower_id from followers f, users u where
f.user_id = u.user_id and u.username = 'mark') ORDER BY tweets.publication_date DESC LIMIT 30

# using tweet created timestamp to get the last one minute tweets instead of using offset approach.
# using timestamp, you can get duplicate tweets if one of Marks followers tweeted between the time the two queries ran

SELECT username, firstname, lastname, tweet, publication_date
FROM users, tweets where users.user_id = tweets.user_id AND publication_date >= NOW() - INTERVAL 1 MINUTE AND users.user_id in  (SELECT follower_id from followers f, users u where
f.user_id = u.user_id and u.username = 'mark') ORDER BY tweets.publication_date DESC LIMIT 30