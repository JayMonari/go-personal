package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"sort"
	"strings"
	"text/template"
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
	ch := flag.String("channel", "jadez", "twitch irc channel to connect to.")
	dur := flag.Duration("duration", 15*time.Minute, "The amount of time for the program to run.")
	flag.Parse()

	conn := JoinTwitchChat("justinfan000", *ch)
	errLog := setupLogging()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	var chatters []chatter
	chatters = []chatter{{username: "based space", messages: set{"has been": {}, "what not": {}}}}
	chattersInfo := make(map[string]chatter)
	line := make([]byte, 512)
out:
	for start := time.Now(); time.Since(start) < *dur; {
		select {
		case <-sig:
			break out
		default:
			// nop
		}
		fmt.Println("Time remaining:", *dur-time.Since(start).Truncate(time.Second))
		n, err := conn.Read(line)
		if err != nil {
			log.Fatal(err)
		}
		data := string(line[:n])
		switch {
		case strings.HasPrefix(data, "PING"):
			conn.Write(ircMsg([]byte("PONG :tmi.twitch.tv")))
			errLog.Println("Ponged Twitch")
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
			errLog.Printf("nickname: %q is NOT the same as username: %q\n", name, usr)
		}

		if c, ok := chattersInfo[name]; ok {
			c.messages[fields[3]] = struct{}{}
		} else {
			chattersInfo[name] = chatter{
				username: name,
				messages: map[string]struct{}{string(fields[3]): {}},
			}
			fmt.Println("added chatter:", name)
		}

		if len(chattersInfo) < 2 {
			continue
		}
		fmt.Println("Message", string(fields[3]))

		for _, c := range chattersInfo {
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

	write(chatters)
	notify()
}

func write(chatters []chatter) {
	tmpl := template.Must(template.ParseFiles("template.js"))

	jsf, err := os.Create("gift.js")
	if err != nil {
		log.Fatal(err)
	}
	if err := tmpl.Execute(jsf, chatters); err != nil {
		log.Fatal(err)
	}

	msgf, err := os.Create("messages.log")
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range chatters {
		msgf.WriteString(fmt.Sprintf("User: %q\n", c.username))
		for m := range c.messages {
			msgf.WriteString(fmt.Sprintf("message: %q\n", m))
		}
		msgf.WriteString("\n")
	}
}

func notify() {
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
