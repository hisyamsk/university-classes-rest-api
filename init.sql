CREATE DATABASE university_classes_db;

\c university_classes_db;

CREATE TABLE student(
  id SERIAL PRIMARY KEY,
  name VARCHAR(254) NOT NULL,
  email VARCHAR(254) NOT NULL,
  active BOOLEAN NOT NULL,
  semester INTEGER NOT NULL
);

CREATE TABLE class(
  id SERIAL PRIMARY KEY,
  name VARCHAR(254) NOT NULL,
  start_at TIME NOT NULL,
  end_at TIME NOT NULL
);

CREATE TABLE enrolled_class(
  id SERIAL PRIMARY KEY,
  student_id INTEGER,
  class_id INTEGER,
  FOREIGN KEY (student_id) REFERENCES student,
  FOREIGN KEY (class_id) REFERENCES class
);