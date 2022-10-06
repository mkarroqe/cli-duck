# cli-duck
Virtual rubber duck that lives in your terminal to support you when you need it. When called, `cli-duck` presents you with an ASCII duck. 
This project is an exercise in creating command line tools in Go using Cobra.

<img src="https://raw.githubusercontent.com/RubberDuckDebugging/rubberduckdebugging.github.io/master/images/rubberducky.png" style="width: 50px" /> Further rubber duck reading: https://rubberduckdebugging.com/

# Usage
`go run main.go cli-duck <command> [--flags]`

### Commands 
| Command | Description |
| - | - |
| hi | Rubber duck will appear. It's ready to listen. |

> Accepting requests for more!

### Flags
| Flag | Example | Description |
| - | - | - |
| `--cool` | `cli-duck hi --cool` | Rubber duck will be wearing sunglasses. |
| `--count` | `cli-duck hi --count=2` | Takes int value for number of ducks that appear. Default value is 1. |
| `--help` | `cli-duck hi --help` | Prints Usage and flags. |
| `--size` |  `cli-duck hi --size large` | Takes string value `small` or `large`. Default value is `small`. |

# Demo
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
marykarroqe@Marys-MBP cli-duck % go run main.go hi --size=large

		      ██████                                    
                    ██      ██                                  
                  ██          ██                                
                  ██      ██  ██                                
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
Sorry, we couldn't find any duckies of size megagigahuge.
```

# Credits
- Small: https://textart.io/art/OUmq7JexuhjpvBnTpL_HAQeF/duckling-swimming
- Large: https://textart.sh/topic/duck
- Cool Mod: yours truly                        
