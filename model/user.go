package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	MangerId int    `json:"manager_id"`
	CoachId  int    `json:"coach_id"`
}

type Teams struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ManagerId int    `json:"manager_id"`
}

type UserTeam struct {
	ID     int `json:"id"`
	UserId int `json:"user_id"`
	TeamId int `json:"team_id"`
}

type Skill struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserSkill struct {
	ID         int `json:"id"`
	UserId     int `json:"user_id"`
	SkillId    int `json:"skill_id"`
	SkillLevel int `json:"skill_level"`
}
