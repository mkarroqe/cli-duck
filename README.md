# cli-duck
Virtual rubber duck that lives in your terminal to support you when you need it. When called, `cli-duck` presents you with an ASCII duck. 
This project is an exercise in creating command line tools in Go using Cobra.

<img src="https://raw.githubusercontent.com/RubberDuckDebugging/rubberduckdebugging.github.io/master/images/rubberducky.png" style="width: 50px" /> Further rubber duck reading: https://rubberduckdebugging.com/

# Usage
`cli-duck [--flags] <command>`

### Commands
| Command | Description |
| - | - |
| help | Prints out usage and options. |
| hi | Rubber duck will appear. It's ready to listen. |

### Options
| Flag | Example | Description |
| - | - | - |
| `--cool` | `cli-duck hi --cool` | Rubber duck will be wearing sunglasses. |
| `--count` | `cli-duck hi --count=2` | Takes int value for number of ducks that appear. Default value is 1. |
| `--size` |  `cli-duck hi --size large` | Takes string value `small` or `large`. Default value is `small`. |

# Credits
Small: https://textart.io/art/OUmq7JexuhjpvBnTpL_HAQeF/duckling-swimming
Large: https://textart.sh/topic/duck

```
        __
    ___( o)>
    \ <_. )
     `---'   
        __
    ___(⌐■)>
    \ <_. )
     `---' 
     
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
              
              
                      ██████                                    
                    ██      ██                                  
                  ██          ██                                
                  ██    ⌐-██--██                                
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

 ```                                               
                                                         
