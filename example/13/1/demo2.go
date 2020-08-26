package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func hasOverlay(ubSuffix, src []byte) (int, bool) {
	overlayLength := len(src) - 1
	for i := overlayLength; i > 0; i-- {
		overlay := src[:i]
		if yes := bytes.HasSuffix(ubSuffix, overlay); yes {
			return i, true
		}
	}
	return -1, false
}

var (
	src string = "你好"
	dst string = "hello"
)

func bigFile(file string) bool {
	fi, err := os.Stat(file)
	if err != nil {
		// 判断文件是否存在, 如果文件不存在, 报错并退出
		if os.IsNotExist(err) {
			log.Fatalf("%v, File isn't existed\n", err)
		} else {
			log.Fatal(err)
		}
	}
	fileSize := fi.Size()
	// 假设超过 1000 bytes 即为大文件
	if fileSize > 1000 {
		return true
	}
	return false
}

func main() {
	// 判断大小
	if big := bigFile("big.txt"); !big {
		fmt.Println("parsing small file...")
		os.Exit(0)
	}
	// 获得文件句柄
	fp, err := os.Open("big.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	// 按段读取
	targetByteNum := len(src)

	buf := bufio.NewReader(fp)
	frag := make([]byte, 100)
	f, err := os.OpenFile("result.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for {
		n, err := buf.Read(frag)
		frag = frag[0:n] // 确保 frag 就是我们读取的内容
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}

		// 特例判断
		for {
			ubSuffix := frag[len(frag)-targetByteNum-1 : len(frag)]
			overlay, yes := hasOverlay(ubSuffix, []byte(src))
			if yes {
				expandSize := targetByteNum - overlay
				expandFrag := make([]byte, expandSize)

				// 读取新的值
				m, err := buf.Read(expandFrag)
				expandFrag = expandFrag[0:m]
				if err != nil {
					if err == io.EOF {
						break
					} else {
						log.Fatal(err)
					}
				}
				// 扩展段的值
				frag = append(frag, expandFrag...)
			} else {
				break
			}
		}

		// 查找和替换
		fragResult := bytes.ReplaceAll(frag, []byte(src), []byte(dst))

		// 写入结果
		if _, err := f.Write(fragResult); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("All content has been parsed!")
}
