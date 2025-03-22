CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    age INT
);

INSERT INTO users (name, age) VALUES
('Taro', 25),
('Hanako', 30),
('Jiro', 22);