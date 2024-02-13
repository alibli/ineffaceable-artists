DROP TABLE IF EXISTS artist;
CREATE TABLE artist (
  id  int AUTO_INCREMENT not null PRIMARY KEY,
  name VARCHAR(128) not null,
  national_code VARCHAR(32) not null
);

DROP TABLE IF EXISTS album;
CREATE TABLE album (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  style      VARCHAR(128),
  -- artist     VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  artist_id INT,
  FOREIGN KEY (artist_id) REFERENCES artist(id),
  PRIMARY KEY (`id`)
);

INSERT INTO artist
  (name, national_code)
VALUES
  ('Ali Bohlooli', 0022789654),
  ('Havva Alimbadi', 0077459215),
  ('Zahra Golmohammadi', 0088564840);

INSERT INTO album
  (title, style, price, artist_id)
VALUES
  ('title one', 'Pop', 56.99, 1),
  ('turkish title', 'Turkish', 63.99, 2),
  ('ghorbani 1', 'Classic', 17.99, 3),
  ('ghorbani 2', 'Classic', 133.4, 3);
