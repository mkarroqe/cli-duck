# cli-duck
Virtual rubber duck that lives in your terminal to support you when you need it. When called, `cli-duck` presents you with an ASCII duck. 
This project is an exercise in creating command line tools in Go using Cobra.

<img src="https://raw.githubusercontent.com/RubberDuckDebugging/rubberduckdebugging.github.io/master/images/rubberducky.png" style="width: 50px" /> Further rubber duck reading: https://rubberduckdebugging.com/

# Usage
`go run main.go cli-duck <command> [--flags]`

### Commands 
| Command | Resource | Flags | Description |
| - | - | - | - |
| hi | - | `--cool`, `--count`, `--size`  | Rubber duck will appear. It's ready to listen. |
| say | nothing | `--cool`, `--size` | The duck says.. nothing. But the tension of a speech bubble is there. |
| say | hello | `--cool`, `--lang`, `--size` | The duck speaks! It says hello. |

> Accepting requests for more!

### Flags
| Flag | Description |
| - | - |
| `--cool` | Rubber duck will be wearing sunglasses. |
| `--count` | Takes int value for number of ducks that appear. Default value is 1. |
| `--help` | Prints Usage and flags. |
| `--lang`  | Changes duck's speaking language. Takes string values `albanian` (duck), `duckspeak` (duck), `english` (human), and `french` (duck). Default is `duckspeak`. |
| `--size` | Takes string value `small` or `large`. Default value is `small`. |

# Examples
```
marykarroqe@Marys-MBP cli-duck % go run main.go hi                    
    __
___( o)>
\ <_. )
 `---'

marykarroqe@Marys-MBP cli-duck % go run main.go hi --cool
    __
___(⌐■)>
\ <_. )
 `---'

marykarroqe@Marys-MBP cli-duck % go run main.go say hello --lang=french 

    __     /coin/
___( o)>  /
\ <_. )
 `---'

marykarroqe@Marys-MBP cli-duck % go run main.go hi --size=large

		              ██████                                    
                    ██      ██                                  
                  ██          ██                                
                  ██      ██   ██                                
                  ██        ░░░░██                              
                    ██      ████                                
      ██              ██  ██                                    
    ██  ██        ████    ██                                    
    ██    ████████          ██                                  
    ██                        ██                                
      ██              ██      ██                                
      ██    ██      ██        ██                                
        ██    ████████      ██                                  
        ██                  ██                                  
          ████          ████                                    
              ██████████      
              
marykarroqe@Marys-MBP cli-duck % go run main.go hi --size=large --cool

		              ██████                                    
                    ██      ██                                  
                  ██          ██                                
                  ██     ⌐-██--██                                
                  ██        ░░░░██                              
                    ██      ████                                
      ██              ██  ██                                    
    ██  ██        ████    ██                                    
    ██    ████████          ██                                  
    ██                        ██                                
      ██              ██      ██                                
      ██    ██      ██        ██                                
        ██    ████████      ██                                  
        ██                  ██                                  
          ████          ████                                    
              ██████████        
              
marykarroqe@Marys-MBP cli-duck % go run main.go hi --size=small --count=2
    __
___( o)>
\ <_. )
 `---'
    __
___( o)>
\ <_. )
 `---'
marykarroqe@Marys-MBP cli-duck % go run main.go hi --size=megagigahuge   
A megagigahuge duck is too much to handle.
```

# Credits
- Small: https://textart.io/art/OUmq7JexuhjpvBnTpL_HAQeF/duckling-swimming
- Large: https://textart.sh/topic/duck
- Cool Mod: yours truly                        
