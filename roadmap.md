# script-kitty

essentially a shell / task-runner & aliasing tool

organize aliases into projects
- default project is the global project (but can change the default prohect to be whatever)
- creates dedicated script-kitty-aliases.sh file in home directory and can be imported into you configured shell interpreters config file
    - essentially acts as a .bash_aliases file
- once the default project changes that new project overwites the script-kitty-aliases.sh file w that projects aliases


all aliases stored in a yaml file at the user level

use like
```
sk coh train
sk list projects
sk lp
sk list
sk l
```

list command shud list all available aliases to user: global and for project (if activated)

create a new project

```
sk create project
```

optional arg to specify directory

optional arg to specify name

will take on cwd for both by default

```
Enter a list of commands to run to activate the project:

(leave last command empty to exit or first command empty to run zero commands)

1. source venv

2.

Enter a list if commands to run to deactivate the project

1.
```

program activates that project like 

sk project-name

program will only deactivate the project if a new project is chose to actuvate or explicit deactivate command is given


sk deactivate project-name

activate also supported but not necessary

sk activate project-name



stretch-goal: support nested projects

initially shud deactivate an active project if create cmd invoked