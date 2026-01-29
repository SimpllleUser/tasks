package main

import "fmt"

// Task 3.8: Deep Copy Struct

// type User struct {
// Name    string
// Friends []*User
// }
// Напиши функцію DeepCopy(user *User) *User яка робить глибоку копію включно з усіма Friends.
type User struct {
	Name    string
	Friends []*User
}

func DeepCopyStruct(user *User) *User {
	visited := make(map[*User]*User)

	var clone func(u *User) *User
	clone = func(u *User) *User {
		if u == nil {
			return nil
		}

		if c, ok := visited[u]; ok {
			return c
		}

		c := &User{
			Name:    u.Name,
			Friends: make([]*User, len(u.Friends)),
		}
		visited[u] = c

		for i, f := range u.Friends {
			c.Friends[i] = clone(f)
		}

		return c
	}

	return clone(user)
}

func main() {

	root := User{
		Name: "Root",
		Friends: []*User{
			{
				Name: "Level1-A",
				Friends: []*User{
					{
						Name: "Level2-A1",
						Friends: []*User{
							{Name: "Level3-A1a"},
							{Name: "Level3-A1b"},
						},
					},
					{
						Name: "Level2-A2",
						Friends: []*User{
							{Name: "Level3-A2a"},
						},
					},
				},
			},
			{
				Name: "Level1-B",
				Friends: []*User{
					{
						Name: "Level2-B1",
						Friends: []*User{
							{Name: "Level3-B1a"},
							{Name: "Level3-B1b"},
							{Name: "Level3-B1c"},
						},
					},
				},
			},
		},
	}

	result := DeepCopyStruct(&root)
	root.Name = "!"
	fmt.Println("Root ", root.Name)
	fmt.Println("Result", result.Name)
}
