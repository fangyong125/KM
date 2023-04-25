package main

import (
	"bytes"
	"fmt"
	"github.com/go-vgo/robotgo"
	"os"
	"strconv"
	"strings"
	"time"
)

func mouseScroll(l int) {
	if l <= 3 {
		help()
		return
	}
	line, err := strconv.Atoi(os.Args[3])
	if err != nil {
		help()
		return
	}
	if line > 0 {
		robotgo.ScrollMouse(line, "down")
	} else {
		robotgo.ScrollMouse(-line, "up")
	}
}
func mouseMove(l int) {
	if l <= 4 {
		help()
		return
	}
	x, err := strconv.Atoi(os.Args[3])
	if err != nil {
		help()
		return
	}
	y, err := strconv.Atoi(os.Args[4])
	if err != nil {
		help()
		return
	}
	robotgo.Move(x, y)
}
func mouseDrag(l int){
	if l <= 4 {
		help()
		return
	}
	x, err := strconv.Atoi(os.Args[3])
	if err != nil {
		help()
		return
	}
	y, err := strconv.Atoi(os.Args[4])
	if err != nil {
		help()
		return
	}
	robotgo.Move(x, y)
}

func keyboardInput(l int) {
	if l <= 3 {
		help()
		return
	}
	text := os.Args[3]
	for _, ch := range text {
		s := string(ch)
		if ch >= 'A' && ch <= 'Z' {
			robotgo.KeyToggle("shift", "down")
			time.Sleep(5 * time.Millisecond)
		}
		robotgo.KeyToggle(s, "down")
		time.Sleep(50 * time.Millisecond)
		robotgo.KeyToggle(s, "up")
		if ch >= 'A' && ch <= 'Z' {
			robotgo.KeyToggle("shift", "up")
			time.Sleep(5 * time.Millisecond)
		}
	}
}

func keyboardHotKey(l int) {
	if l <= 3 {
		help()
		return
	}
	text := os.Args[3]
	keys := strings.Split(text, " ")
	kl := len(keys)
	for i := 0; i < kl; i++ {
		robotgo.KeyToggle(keys[i], "down")
		//fmt.Println(keys[i], "down")
		time.Sleep(10 * time.Millisecond)
	}
	for i := kl - 1; i >= 0; i-- {
		robotgo.KeyToggle(keys[i], "up")
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	var cmdBuilder bytes.Buffer
	l := len(os.Args)
	for i := 1; i < l; i++ {
		cmdBuilder.WriteString(os.Args[i])
		cmdBuilder.WriteString(" ")
	}
	cmd := cmdBuilder.String()

	if strings.HasPrefix(cmd, "mouse click left") {
		robotgo.Click(`left`, false)
		return
	}
	if strings.HasPrefix(cmd, "mouse click right") {
		robotgo.Click(`right`, false)
		return
	}
	if strings.HasPrefix(cmd, "mouse dbclick left") {
		robotgo.Click(`left`, true)
		return
	}
	if strings.HasPrefix(cmd, "mouse dbclick right") {
		robotgo.Click(`right`, true)
		return
	}
	if strings.HasPrefix(cmd, "mouse scroll") {
		mouseScroll(l)
		return
	}
	if strings.HasPrefix(cmd, "mouse position") {
		x, y := robotgo.GetMousePos()
		fmt.Println(x, y)
		return
	}
	if strings.HasPrefix(cmd, "mouse move") {
		mouseMove(l)
		return
	}
	if strings.HasPrefix(cmd, "mouse drag") {
		mouseDrag(l)
		return
	}
	if strings.HasPrefix(cmd, "keyboard input") {
		keyboardInput(l)
		return
	}
	if strings.HasPrefix(cmd, "keyboard hotkey") {
		keyboardHotKey(l)
		return
	}
	help()
}
func help() {
	fmt.Println("km mouse/keyboard options ")
	fmt.Println("km mouse click left             单击鼠标左键")
	fmt.Println("km mouse click right            单击鼠标右键")
	fmt.Println("km mouse dbclick left           双击鼠标左键")
	fmt.Println("km mouse dbclick right          双击鼠标右键")
	fmt.Println("km mouse scroll 5               向下滚动5行")
	fmt.Println("km mouse scroll -5              向上滚动5行")
	fmt.Println("km mouse position               打印当前坐标")
	fmt.Println("km mouse move 100 200           移动到(100,200)")
	fmt.Println("km mouse drag 100 200           拖动到(100,200)")
	fmt.Println("km keyboard input AZza,09       输入AZza,09")
	fmt.Println("km keyboard hotkey \"ctrl c\"     模拟复制ctrl+c")
	fmt.Println("km keyboard hotkey 热键")
	fmt.Println("cmd = Windows 系统中 win 键")
	fmt.Println("lcmd = Windows 系统中左边 win 键")
	fmt.Println("rcmd = Windows 系统中右边 win 键")
	fmt.Println("ctrl = ctrl")
	fmt.Println("lctrl = 左边 ctrl 键")
	fmt.Println("rctrl = 右边 ctrl 键")
	fmt.Println("shift = shift")
	fmt.Println("tab = tab")
	fmt.Println("alt = alt")
	fmt.Println("lalt = 左边 alt 键")
	fmt.Println("ralt = 右边 alt 键")
	fmt.Println("space = 空格")
	fmt.Println("up = ↑键")
	fmt.Println("down = ↓键")
	fmt.Println("right = →键")
	fmt.Println("left = ←键")
}
