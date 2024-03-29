package cli

import (
	"bufio"
	"os"
	"strconv"
	"time"

	api "heid9/downtime/api/cli"
)

type Command struct {
	args  []string
	input []string
}

func NewCommand() api.Command {
	c := &Command{}
	c.input = c.parseStdin()
	c.args = os.Args[1:]
	return c
}

func (c *Command) Urls() []string {
	return c.input
}

func (c *Command) DurationArg() (dur time.Duration, ok bool) {
	if len(c.args) != 2 {
		return
	}
	key, val := c.args[0], c.args[1]
	if key != "-i" {
		return
	}
	res, err := strconv.Atoi(val)
	if err != nil {
		return
	}
	return time.Duration(res) * time.Second, true
}

func (c *Command) parseStdin() (lines []string) {
	file := os.Stdin
	stat, err := file.Stat()
	if err != nil {
		return []string{}
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			txt := scanner.Text()
			lines = append(lines, txt)
		}
	}
	return
}
