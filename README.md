# Web Analyzer

This Project provides the general information of webpage. 

###Description: 

Take URL and input and give the following insight of its content

* **HTML Version**
    
    Identify Html Version by looking at doctype tag. Can identify HTML 5, HTML 4.01 Strict, HTML 4.01 Transitional, HTML 4.01 Frameset, XHTML 1.0 Strict, XHTML 1.0 Transitional, XHTML 1.0 Frameset, and XHTML 1.1 
    
*  **Page Tile**
   
    Identify from title tage

*  **Heading Tag**
    
    Gives the count of all heading tag by level.

*  **Count of Internal link**
    
    Sum of relative links and all link with same domain name as input URL

*  **Count of External links**
    
    Sum of links whose domain does not match with input url domain

*  **Count of inaccessible links**
*  **Login Form**
    
    Detect if page have a login from by looking for input type password

### Prerequisites

Go Version 1.13 or later Installed 

### Steps

*  Clone the git repo 
*  Go the project folder and type ```go run .```
*  App is running on port 8000 you can access if from your browser by going to http://localhost:8000/
 
