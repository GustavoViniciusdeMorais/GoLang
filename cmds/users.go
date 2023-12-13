package cmds

type User struct {
	id    int64
	name  string
	email string
}

func All() string {
	users := "All users\n"
	return users
}
