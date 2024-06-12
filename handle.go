package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (c *Config) CompileHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !Run[m.Author.ID] {
		c.CompileRun(s, m)
	} else {
		s.ChannelMessageSend(m.ChannelID, "```you are running...```")
	}
}

func (c *Config) CompileRun(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, ";compile") {
		if len(m.Content) > 11 {
			parse := m.Content[8+3:]
			parse = parse[:len(parse)-3]
			if strings.HasPrefix(parse, "go") {
				Run[m.Author.ID] = true
				if is, err := c.InfiniteLoop(parse[2:]); !is && c.isSafe(parse[2:]) {
					if out, err := Compile(parse[2:]); err == nil && out != "" {
						s.ChannelMessageSend(m.ChannelID, "```"+string(out)+"```")
					} else {
						fmt.Println(err)
					}
				} else {
					if is && err == nil {
						s.ChannelMessageSend(m.ChannelID, "```no loop```")
					} else if err != nil {
						fmt.Println(err)
					} else {
						s.ChannelMessageSend(m.ChannelID, "```you are send banned code```")
					}
				}
				delete(Run, m.Author.ID)
			}
		}
	}
}
