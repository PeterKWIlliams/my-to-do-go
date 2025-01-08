-- +goose Up
CREATE TABLE Projects (
  id TEXT PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE,
  working_directory TEXT,
  desired_work_time INT DEFAULT 0,
  total_work_time INT DEFAULT 0,
  project_no INT UNIQUE,
  last_worked_on TEXT,
  created_at TEXT NOT NULL DEFAULT DATE('now'),
  updated_at TEXT NOT NULL DEFAULT DATE('now')
);

-- +goose Down
DROP TABLE Projects;



