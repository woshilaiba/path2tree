package common

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type pathmeta struct {
	Level           int
	Path_name       string
	Next_path       map[string]*pathmeta
	Next_path_slice []*pathmeta
}

var (
	rootdir    pathmeta
	level_sign []string
)

const (
	empty_prefix  = "    "
	normal_prefix = "│   "
	middle_prefix = "├── "
	last_prefix   = "└── "
	path_sep      = "/"
)

func init() {
	rootdir.Level = 0
	rootdir.Path_name = "root"
	rootdir.Next_path = make(map[string]*pathmeta, 100)
	level_sign = make([]string, 20)
}
func Find_or_make_pathmeta(path string) {
	//must get rid of "\n", or will show a parent path twice
	if path[len(path)-1] == '\n' {
		path = path[:len(path)-1]
	}
	paths := strings.Split(path, path_sep)
	if paths != nil {
		var parent *pathmeta
		var ok bool
		emptypath_cnt := 0
		for i, pp := range paths {
			if pp == "" {
				emptypath_cnt++
				continue
			}
			var son *pathmeta
			if i-emptypath_cnt == 0 {
				//this is the root
				parent = &rootdir
			}
			son, ok = parent.Next_path[pp]
			if !ok {
				//not exist ,create it
				tt := &pathmeta{Level: i - emptypath_cnt + 1, Path_name: pp, Next_path: make(map[string]*pathmeta)}
				parent.Next_path[pp] = tt
				parent.Next_path_slice = append(parent.Next_path_slice, tt)
				son = tt
			}
			parent = son
		}
	}
}

//recursive call for traversing
func print_pathmeta(p *pathmeta) {
	//print current node
	for i := 0; i < p.Level; i++ {
		fmt.Print(level_sign[i])
	}
	//current node name
	fmt.Println(p.Path_name)
	//traverse the subnodes, set self level_sign
	if len(level_sign) <= p.Level {
		level_sign = append(level_sign, "")
	}
	//prepair for printing subnodes
	var orig_sign string
	//set level_sign for subnodes printing. it's each line's prefix, such as  "|   " "├──"
	level_sign[p.Level] = middle_prefix
	if len(p.Next_path_slice) != 0 && p.Level > 0 {
		orig_sign = level_sign[p.Level-1]
		if orig_sign == middle_prefix {
			level_sign[p.Level-1] = normal_prefix
		} else if orig_sign == last_prefix {
			level_sign[p.Level-1] = empty_prefix
		}
	}
	for i, v := range p.Next_path_slice {
		if i+1 == len(p.Next_path_slice) {
			//the last of the subnodes，change level_sign[p.Level] to "└── "
			level_sign[p.Level] = last_prefix
		}
		print_pathmeta(v)
	}
	//after printing subnodes，restore the level_sign
	if len(p.Next_path_slice) != 0 && p.Level > 0 {
		level_sign[p.Level-1] = orig_sign
	}
}

func Print_pathmeta() {
	print_pathmeta(&rootdir)
}

//read a file for input， or STDIN
func Readfile(filepaht string) {
	var f *os.File
	var er error
	if filepaht == "" {
		f = os.Stdin
	} else {
		f, er = os.Open(filepaht)
		if er != nil {
			fmt.Println(er.Error())
			panic(er)
		}
	}
	r := bufio.NewReader(f)
	for {
		buf, err := r.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				//abnormal
				fmt.Println(err.Error())
				panic(err)
			} else if len(buf) == 0 {
				//finished reading
				return
			}
		}
		Find_or_make_pathmeta(buf)
	}
}
