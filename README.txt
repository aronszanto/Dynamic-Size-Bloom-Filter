Instructions (for OS X)

Install Go using Homebrew. 

Set up your Go environment as detailed by https://golang.org/doc/install. You can see our working directory (https://github.com/aszanto9/Blumo) and a screenshot of our environment in ../report. 

Blumo on our laptops is found at the following directory: $USER/Projects/src/github.com/aszanto9. 

“Projects” is the directory with contains all other Go projects, and the folder which contains bitset is found here: $USER/Projects/src/github.com/willf. You will need to edit some of the “include” statements in the .go files if you store folders differently. Once you’ve set up the above, open your terminal and navigate to the main directory. 

Copy and paste the following into the terminal to run our program: 

go install 
go build 
./main

Please note than instead of placing our dictionaries in the “data” file as suggested in the project outline, we have stored it in Dictionaries. 