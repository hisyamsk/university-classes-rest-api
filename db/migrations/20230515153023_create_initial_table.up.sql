CREATE TABLE student(
  id SERIAL PRIMARY KEY,
  name VARCHAR(254) NOT NULL,
  email VARCHAR(254) NOT NULL,
  active BOOLEAN NOT NULL,
  semester INTEGER NOT NULL
);

CREATE TABLE class_detail(
  id SERIAL PRIMARY KEY,
  name VARCHAR(254) NOT NULL,
  start_at TIME NOT NULL,
  end_at TIME NOT NULL
);

CREATE TABLE class(
  id SERIAL PRIMARY KEY,
  student_id INTEGER,
  class_detail_id INTEGER,
  FOREIGN KEY (student_id) REFERENCES student,
  FOREIGN KEY (class_detail_id) REFERENCES class_detail
);