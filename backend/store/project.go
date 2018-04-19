package store

import (
	"log"
)

type Project struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Weeks  float64 `json:"weeks,omitempty"`
	Person string  `json:"person"`
}

func CreateProject(project Project) error {
	var person *Person
	person, err := GetOrCreatePerson(project.Person)
	if err != nil {
		log.Panic(err)
		return err
	}
	_, err = db.Exec("INSERT INTO projects (name, weeks, person) VALUES ($1, $2, $3)", project.Name, project.Weeks, person.ID)
	if err != nil {
		log.Panic(err)
		return err
	}
	return nil
}

func UpdateProject(project Project) error {
	var person *Person
	person, err := GetOrCreatePerson(project.Person)
	if err != nil {
		log.Panic(err)
		return err
	}
	_, err = db.Exec("UPDATE projects SET name = $1, weeks = $2, person = $3 WHERE id = $4", project.Name, project.Weeks, person.ID, project.ID)
	if err != nil {
		log.Panic(err)
		return err
	}
	return nil
}

func GetProjects() ([]Project, error) {
	rows, err := db.Query("select projects.id, projects.name, projects.weeks, people.name from projects LEFT JOIN people ON (projects.person = people.id)")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var projects []Project
	for rows.Next() {
		var project Project
		err = rows.Scan(&project.ID, &project.Name, &project.Weeks, &project.Person)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func DeleteProject(id int) error {
	_, err := db.Exec("DELETE FROM projects WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
