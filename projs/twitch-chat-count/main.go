package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"
)

type ircMessage struct {
	prefix
	command string
	params  string
}

func (m ircMessage) String() string {
	return m.prefix.String() + " " + m.command + " :" + m.params
}

type prefix struct {
	nick string
	user string
	host string
}

type chatter struct {
	username string
	messages set
}

type set map[string]struct{}

func (s set) String() string {
	var sb strings.Builder
	for k := range s {
		sb.WriteString(k)
	}
	return sb.String()
}

func (c chatter) String() string { return c.username }

func (p prefix) String() string { return p.nick + "!" + p.user + "@" + p.host }

func main() {
	chn := flag.String("channel", "jadez", "twitch irc channel to connect to.")
	flag.Parse()
	conn := JoinTwitchChat("justinfan000", *chn)
	errLog := setupLogging()
	f, err := os.Create("chatters.log")
	if err != nil {
		log.Fatal(err)
	}

	var chatters []chatter
	usersInfo := make(map[string]chatter)
	line := make([]byte, 512)
	for start := time.Now(); time.Since(start) < (1*time.Hour + 5*time.Minute); {
		n, err := conn.Read(line)
		if err != nil {
			log.Fatal(err)
		}
		data := string(line[:n])
		switch {
		case strings.HasPrefix(data, "PING"):
			conn.Write(ircMsg([]byte("PONG :tmi.twitch.tv")))
			log.Println("Ponged Twitch")
			continue
		case strings.HasPrefix(data, ":tmi.twitch.tv") || !strings.Contains(data, "PRIVMSG"):
			continue
		}

		fields := strings.SplitN(data, " ", 4)
		if len(fields) < 4 {
			errLog.Printf("data: %q\nfields: %+v\n", data, fields)
			continue
		}
		nm, usrhost, ok := strings.Cut(fields[0][1:], "!")
		if !ok {
			continue
		}
		name := string(nm)
		usr, _, ok := strings.Cut(usrhost, "@")
		if !ok {
			continue
		}
		switch name {
		case "streamelements", "fossabot", "nightbot":
			continue
		case usr:
			// nop
		default:
			log.Printf("nickname: %q is NOT the same as username: %q\n", name, usr)
		}

		if c, ok := usersInfo[name]; ok {
			c.messages[fields[3]] = struct{}{}
		} else {
			usersInfo[name] = chatter{
				username: name,
				messages: map[string]struct{}{string(fields[3]): {}},
			}
			fmt.Println("added chatter:", name)
		}

		if len(usersInfo) < 2 {
			continue
		}
		fmt.Println("Message", string(fields[3]))

		for _, c := range usersInfo {
			if contains(chatters, c.username) {
				continue
			}
			chatters = append(chatters, c)
		}

		sort.Slice(chatters, func(i, j int) bool {
			return len(chatters[i].messages) > len(chatters[j].messages)
		})

		for i := 0; i < len(chatters) && i < 5; i++ {
			fmt.Printf("User: %q has sent %d messages\n", chatters[i].username, len(chatters[i].messages))
		}
	}

	for i := 0; i < len(chatters) && i < 5; i++ {
		fmt.Printf("User: %q\n", chatters[i].username)
		for m := range chatters[i].messages {
			fmt.Printf("message: %q\n", m)
		}
	}

	for _, c := range chatters {
		f.Write([]byte(c.username + "\n"))
	}

	cmd := exec.Command("notify-send", "-u", "critical", "Completed watching chat")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	cmd = exec.Command("aplay", "/home/jay/dls/mixkit-positive-notification-951.wav")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func JoinTwitchChat(nick, channel string) net.Conn {
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Write(ircMsg([]byte("NICK " + nick)))
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Write(ircMsg([]byte("JOIN #" + channel)))
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func ircMsg(a []byte) []byte {
	return append(a, []byte("\r\n")...)
}

func setupLogging() *log.Logger {
	f, err := os.Create("bad_tings.log")
	if err != nil {
		log.Fatal(err)
	}
	return log.New(f, "", log.LstdFlags)
}

func contains(a []chatter, username string) bool {
	for _, c := range a {
		if c.username == username {
			return true
		}
	}
	return false
}
