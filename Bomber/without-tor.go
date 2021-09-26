// +build withoutTor

package main

import "github.com/fatih/color"

func attachTor() {
	tor := color.New(color.FgGreen).Sprint("Tor")
	println(tor + " не будет запущен! Вы НЕ в безопасности :[\nЧтобы " + tor + " использовался, необходимо скачать соответствующую версию программы!")
}
