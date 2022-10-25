/*
Copyright © 2022 Mary Karroqe <mkarroqe@gmail.com>
*/
package cmd

type Duck struct {
	Head string
	Eyes string
	Bill string
	Body string

	Size    string
	Cool    bool
	Speaks  bool
	Message string
}

func (d Duck) CreateLittleDuck() string {
	d.Head = "\n    __"
	d.Eyes = "___( o)>"
	d.Body = "\\ <_. )\n `---'\n"
	var fullDuck string

	if d.Speaks {
		d.Head += "     /" + d.Message + "/"
		if d.Cool {
			d.Eyes = "___(⌐■)>"
		}
		d.Eyes += "  /"
	}

	fullDuck += d.Head + "\n" + d.Eyes + "\n" + d.Body
	return fullDuck
}

func (d Duck) CreateLargeDuck() string {
	d.Head = "\n		      ██████                                    \n                    ██      ██                                  \n                  ██          ██"
	d.Eyes = "                  ██       ██  ██"
	d.Bill = "                  ██        ░░░░██"
	d.Body = "                    ██      ████                                \n      ██              ██  ██                                    \n    ██  ██        ████    ██                                    \n    ██    ████████          ██                                  \n    ██                        ██                                \n      ██              ██      ██                                \n      ██    ██      ██        ██                                \n        ██    ████████      ██                                  \n        ██                  ██                                  \n          ████          ████                                    \n              ██████████        \n"
	var fullDuck string

	if d.Speaks {
		d.Head += "        __________"
		if d.Cool {
			d.Eyes = "                  ██     ⌐-██--██"
		}
		d.Eyes += "      /  " + d.Message + "  /"
		d.Bill += "   < _________/"
	}

	fullDuck += d.Head + "\n" + d.Eyes + "\n" + d.Bill + "\n" + d.Body
	return fullDuck
}
