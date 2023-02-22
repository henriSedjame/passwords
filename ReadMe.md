
_**Passwords** is a cli tool to store passwords on a local disk_

### HOW TO USE PASSWORDS ?

* Clone the github project
     
        git clone https://github.com/henriSedjame/passwords.git

* Build the project in the directory of your choice

      go build -o <out-dir>/passwords <project-dir>/src

* Add the cli to your path and enjoy 

  * To see all operations that be achieved with password cli, run :
      
        passwords --help

  * To add a new password, run :

        passwords -add -label="The paassword label" -value="The password value" 
  
  * To update a stored password, run :

        passwords -update -label="The paassword label" -value="The new password value"

  * To remove a stored password, run :

        passwords -delete -label="The password label"
  
  * To show a stored password, run:
     
        passwords -show -label="The password label"
  
  * To show all passwords, run :
  
        passwords -list
