# simple_paster

I AM NOT A PROGRAMMER
I do not code for a living or as a hobby and this tool (barely even a script) was made with
heavy input from ChatGPT (free). Even if I defined what each section was meant to do, tested,
checked functionality with success/fail statements, etc., I do not anticipate this being
relevant to other people and have no intention of adding new functionalities or somesuch.

This is a product of SPITE
This tool exists because I'm sick of the aspects of my life which require repeating the same
information over and over. That's literally it.

SIMPLE AND PRIVATE were the only two things I cared about
The whole purpose was to have the effect of a text expander without any of the weird and
(to the layman like me) obscured or at least translucent keylogger-ish behaviours normal
text expanders come with.

BEHAVIOUR:
1. In this version the .exe and .csv must be in the same folder, but folder location doesn't
matter. I like to stick a shortcut on my desktop and add a hotkey combination in properties.
If you want a specific path for the .csv instead I suggest something like:

 usr, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("ERROR MESSAGE: %v", err)
	}
	csvPath := filepath.Join(usr.HomeDir, "[DIRECTORY1, eg. Documents]", "[DIRECTORY2]", "snippets.csv")

Note: this requires "os/user" to be imported too.
 
3. .exe opens the Windows CLI, loads the .csv and performs regex search of column 1 for the
CLI input. Search can be performed as often as desired (eg. if there's a typo, etc).
4. Upon selecting a result (max. of 5 results displayed) the corresponding text from column 2
is copied to clipboard and the CLI closes.

Aside from the text being copied nothing is stored, nothing is executed.
