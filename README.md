
# ASCII - ART
* `ayessenb` 
* `amashpiy` 


## Second big project in go lang (testing 3 months period)

We created the program which receives as arguments the string and/or name of a file without ".txt" extension, containing some banner with a specific graphical template representation using ASCII (eg. standard, shadow or thinkertoy) will be given. 

* go.mod - defines a module, which is a collection of Go packages. It contains the data about the module, such as its name, version, and dependencies on other modules.
* go.sum - contain a list of all the downloaded modules, their versions, and their hash.
* main.go - ascii art main function which executes all the functions related to ascii from funcs folder as well as the CheckFileHashing, which do handle the case when one of the banner files were changed from the original.
* additional.go - checks all the validity, creates the map where the run is mapped to the array of string for each banner letter, clearning the files for thinkertone case, since it contains the 13 ascii value, whihc is out of range for us, after the word was inserted the function GetWord does the mapping and returns the ascii art image of the input string.

#### Description:


* The program reads the banner file where the height of each letter is 8 lines
* The first line is not considered
* Each character in banner file is separated by one '\n'
* Since Banner contains the ascii values between 32 and 126 then the other runes are not accepted.
* If the Banner file will be changed the program will give the Error message
* If the program will get less than one or more than two arguments, the program will give the Error message
* If the second argument is not "standard" or "shadow" or "thinkertoy" then the program will give the Error messages
* If the input string argument doesn't contain ascii value then the program gives the Error message.(between 32 and 126 (both included))


#### Improved skills:
* go lang programming skills 
* usage of string libraries of Go
* handling new line char.
* usage of many functions and files

## Usage/Examples
Cloning storage to your host
```CMD/Terminal 
git clone git@git.01.alem.school:ayessenb/ascii-art.git
```
Go to the downloaded repository:

```CMD/Terminal 
cd ascii-art
```
Run a program:
```CMD/Terminal 
go run . "Salem Alem" | cat -e
```


Output: __Salem Alem__
+ standard.txt
```bash
  _____           _                                       _                     $
 / ____|         | |                              /\     | |                    $
| (___     __ _  | |   ___   _ __ ___            /  \    | |   ___   _ __ ___   $
 \___ \   / _` | | |  / _ \ | '_ ` _ \          / /\ \   | |  / _ \ | '_ ` _ \  $
 ____) | | (_| | | | |  __/ | | | | | |        / ____ \  | | |  __/ | | | | | | $
|_____/   \__,_| |_|  \___| |_| |_| |_|       /_/    \_\ |_|  \___| |_| |_| |_| $
                                                                                $
                                                                                $
```


Run a program:
```CMD/Terminal 
go run . "Salem Alem" standard | cat -e
```


Output: __Salem Alem__
+ standard.txt
```bash
  _____           _                                       _                     $
 / ____|         | |                              /\     | |                    $
| (___     __ _  | |   ___   _ __ ___            /  \    | |   ___   _ __ ___   $
 \___ \   / _` | | |  / _ \ | '_ ` _ \          / /\ \   | |  / _ \ | '_ ` _ \  $
 ____) | | (_| | | | |  __/ | | | | | |        / ____ \  | | |  __/ | | | | | | $
|_____/   \__,_| |_|  \___| |_| |_| |_|       /_/    \_\ |_|  \___| |_| |_| |_| $
                                                                                $
                                                                                $
```

Run a program:
```CMD/Terminal 
go run . "Salem Alem" thinkertoy | cat -e
```


Output: __Salem Alem__
+ thinkertoy.txt
```bash
                                               $
 o-o       o                   O   o           $
|          |                  / \  |           $
 o-o   oo  | o-o o-O-o       o---o | o-o o-O-o $
    | | |  | |-' | | |       |   | | |-' | | | $
o--o  o-o- o o-o o o o       o   o o o-o o o o $
                                               $
                                               $
```


Run a program:
```CMD/Terminal 
go run . "Salem Alem" shadow | cat -e
```

Output: __Salem Alem__
+ shadow.txt
```bash
                                                                                       $
  _|_|_|          _|                                 _|_|   _|                         $
_|         _|_|_| _|   _|_|   _|_|_|  _|_|         _|    _| _|   _|_|   _|_|_|  _|_|   $
  _|_|   _|    _| _| _|_|_|_| _|    _|    _|       _|_|_|_| _| _|_|_|_| _|    _|    _| $
      _| _|    _| _| _|       _|    _|    _|       _|    _| _| _|       _|    _|    _| $
_|_|_|     _|_|_| _|   _|_|_| _|    _|    _|       _|    _| _|   _|_|_| _|    _|    _| $
                                                                                       $
                                                                                       $
```


## Feedback

If you liked our project, we would be grateful if you could add `Star` to the repository.

Alem Student
10.05.2023.