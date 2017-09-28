package dev

import "testing"

func TestCommand(t *testing.T) {
	tmpout := newtempfile()
	tmperr := newtempfile()
	defer tmpout.remove()
	defer tmperr.remove()

	t.Run("valid command should be error free", func(t *testing.T) {
		c := &command{
			name:   "uptime",
			stdout: tmpout.f,
			stderr: tmperr.f,
		}
		if err := c.run(); err != nil {
			t.Error(err)
		}
	})

	t.Run("non-existing command should error.", func(t *testing.T) {
		c := &command{
			name:   "noexist",
			stdout: tmpout.f,
			stderr: tmperr.f,
		}
		if err := c.run(); err == nil {
			t.Error(err)
		}
	})

	t.Run("command with incorrect args, should error.", func(t *testing.T) {
		c := &command{
			name:   "ls",
			args:   []string{"foo"},
			stdout: tmpout.f,
			stderr: tmperr.f,
		}
		if err := c.run(); err == nil {
			t.Error(err)
		}
	})
}

func TestCommandParsing(t *testing.T) {
	t.Run("should split command name and args correctly", func(t *testing.T) {
		cmd := "tail -f README.md"
		name, args := parseCmd(cmd)
		if name != "tail" {
			t.Error("Expected the command name to be tail")
		}
		if len(args) != 2 {
			t.Error("Expected the args length to be 2")
		}

		if args[0] != "-f" || args[1] != "README.md" {
			t.Error("Expected the arguments to be in order")
		}
	})

	t.Run("commands with no args should return an empty array", func(t *testing.T) {
		cmd := "dev"
		name, args := parseCmd(cmd)
		if name != "dev" {
			t.Error("Expected name to be dev")
		}
		if len(args) > 0 {
			t.Error("Expected args length to be 0")
		}
	})

	t.Run("an empty string for a command should give a blank name and args len of 0", func(t *testing.T) {
		cmd := ""
		name, args := parseCmd(cmd)
		if name != "" {
			t.Error("Expected a blank name")
		}
		if len(args) != 0 {
			t.Error("Expected length of args to be 0")
		}
	})

}
