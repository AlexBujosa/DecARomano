# Index

1. **About the algorithm**
2. **Installation**
3. **Documentation, Requirments, Acceptance Criteria, and Test Cases**
4. **Questionnarie**
5. **Arabic numerals to roman numerals**

---

## 1. **About the algorithm**

- This algorithm has functionality to convert Arabic numerals to Roman numerals, which are the numbers we use when we talk about century and other things, this number not include the zero one. This algoriths works from a broad case to base case using recursion and adding the appropiate letter. When we talk about thousands of number we use from 4000 (IV|) using the sidebar (|) to express thousands. The exception of thousands are 1, 2, 3 in Roman numerals are M, MM , MMM and the millions are M|, MM|, MMM|, ... MMMMMMMMMM|. The limit number is 9999999.

## 2. **Installation**

Firstable, you need to install the lastest version of Go. Using this command

- Go to downloads in the terminal and the you write wget https://go.dev/dl/go1.18.2.linux-amd64.tar.gz
- Then you run tar -xzvf go1.18.2.linux-amd64.tar.gz
- Optional (if you want to folders inside your go folder yo can use ll go, the meaning of ll is list all <'Name of folder '>)
- To see all the commands that go gives you, you write .go/bin/go in your terminal
- Now to use all the commands that you saw in the step before, everywhere you are. You need to move your go folder to /usr/local/. Using the next command to move: sudo mv go /usr/local/ and then you try using go in your terminal if the command is not recognized you need access to env-variable
- - write in your terminal -> vi ~/.bashrc
- - then press shift + g at the same time, this redirect you to bottom of the file and press O and then enter
- - Now you need to write #Add go path (this is a comment)
- - then you write export PATH=$PATH:/usr/local/bin, then press control c to get out and then you write :wq and press enter
- - then you write source ~/.bashrc and press enter
- - and finally you write go, and then this commands displays all the go commands, you can use

**For using the programm you need to run the next commands**

- cd Documents, to locate inside Documents folder then sudo mkdir (name_folder)
- then cd name_folder
- then you write git clone https://github.com/AlexBujosa/DecimalARomano
- cd decimalAromano
- then write go run decimalAromano.go
- Finally this go to display the console application

## 3. **Documentation, Requirments, Acceptance Criteria, and Test Cases**

[Requirements, Acceptance, and Test Cases](Docs/ReqCriCas.md)

## 4. **Questionnaire**

[Questionnaire](Docs/cuestionario.md)

## 5. **Arabic numerals to roman numerals**

[ArabicToroman](decimalAromano.go)
