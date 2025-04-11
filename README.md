# simple_paster

I AM NOT A PROGRAMMER <br>
I do not code for a living or as a hobby and this tool (barely even a script) was made with <br>
heavy input from ChatGPT (free). Even if I defined what each section was meant to do, tested, <br>
checked functionality with success/fail statements, etc., I do not anticipate this being <br>
relevant to other people and have no intention of adding new functionalities or somesuch. <br>
 <br>
This is a product of SPITE <br>
This tool exists because I'm sick of the aspects of my life which require repeating the same <br>
information over and over. That's literally it. <br>
 <br>
SIMPLE AND PRIVATE were the only two things I cared about <br>
The whole purpose was to have the effect of a text expander without any of the weird and <br>
(to the layman like me) obscured or at least translucent keylogger-ish behaviours normal <br>
text expanders come with. <br>
 <br>
BEHAVIOUR: <br>
1. In this version the .exe and .csv must be in the same folder, but folder location doesn't <br>
matter. I like to stick a shortcut on my desktop and add a hotkey combination in properties. <br>
If you want a specific path for the .csv instead I suggest something like: <br>
 <br>
 usr, err := user.Current() <br>
	if err != nil { <br>
		return nil, fmt.Errorf("ERROR MESSAGE: %v", err) <br>
	} <br>
	csvPath := filepath.Join(usr.HomeDir, "[DIRECTORY1, eg. Documents]", "[DIRECTORY2]", "snippets.csv") <br>
 <br>
Note: this requires "os/user" to be imported too. <br>
 <br> 
3. .exe opens the Windows CLI, loads the .csv and performs regex search of column 1 for the <br>
CLI input. Search can be performed as often as desired (eg. if there's a typo, etc). <br>
4. Upon selecting a result (max. of 5 results displayed) the corresponding text from column 2 <br>
is copied to clipboard and the CLI closes. <br>
 <br>
Aside from the text being copied nothing is stored, nothing is executed.
