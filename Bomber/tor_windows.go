// +build !withoutTor

package main

import (
	"net"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/fatih/color"
	proxy "github.com/valyala/fasthttp/fasthttpproxy"
)

func attachTor() {
	println("Запуск " + color.New(color.FgGreen).Sprint("Tor") + "...")

	var (
		err error

		i, status int

		n, port  string
		testsite = "http://example.com/"

		torpath = filepath.Join(execDir, "tor")
		// torpath, _  = filepath.Abs("./tor-for-tests")
		torDataPath = filepath.Join(torpath, "Data")

		torDataDir string

		tor = &exec.Cmd{
			Path: filepath.Join(torpath, "Tor", "tor.exe"),
			Args: []string{
				"tor.exe",
				"-f", filepath.Join(torpath, "torrc"),
				"--DataDirectory", "",
				"--SOCKSPort", "",
			},
			Dir: torpath,
		}

		ln net.Listener
	)

	for {
		if i == 9 {
			println("Не удалось запустить Tor! Попробуйте убить все процессы tor в диспетчере задач и/или перезапустите бомбер!")
			shutdown(true)
		}

		n = strconv.Itoa(i)
		port = "376" + n

		ln, err = net.Listen("tcp", ":"+port)
		if err != nil || ln == nil {
			i++
			continue
		}
		ln.Close()

		torDataDir = filepath.Join(torDataPath, "data"+n)

		tor.Args[4] = torDataDir
		tor.Args[6] = port

		err = tor.Start()
		if err != nil {
			println("Не удалось запустить Tor!")
			errCheck(err)
		}

		// Прописывание Tor-proxy в http-клиенте
		client.Dial = proxy.FasthttpSocksDialer("127.0.0.1:" + port)

		for b := 0; b < 3; b++ {
			status, _, err = client.Get(nil, testsite)
			if status == 200 && err == nil {
				break
			}
			time.Sleep(time.Second * 5)
		}
		status, _, err = client.Get(nil, testsite)
		if status == 200 && err == nil {
			break
		} else {
			tor.Process.Kill()
		}

		i++
	}

	println(color.New(color.FgGreen).Sprint("Tor") + " запущен! Вы в безопасности ;)")
}
