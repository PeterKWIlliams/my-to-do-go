-- name: CreateProjects :exec
INSERT INTO Projects (
  id, name,working_directory,desired_work_time,
  total_work_time,project_no,created_at, updated_at
  )
VALUES (?, ?, ?, ?, ?, ?, ?, ?);
--

-- name: GetAllProjects :many

SELECT * FROM Projects;


--

