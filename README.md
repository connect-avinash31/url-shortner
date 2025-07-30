# url-shortner
This repo contains the code related to url shortner
1. this repo has 3 layer of code 
	1. REST Interface which has exposed the functionlity of shortner , fetch original and metrics
		- here help of internet has been taken to know the converison lirbary and syntax 
	2. Shortner services which is internally handling of ShortenURL , OriginalURL and Metrics 
		- no internet help has been taken 
	3. Hashing mechanism 
		- written a small has function which every time multiply currenbt value with 31 and add new value of string characters
		- here no guarantee of clossion avoidance but still less collision  will occure
2. Docker file has follwoing details
	1. An alpine base image of ubuntu where go1.20 is extracted
        2. Then we are setting few envioments just to make sure go can run successfully
        3. then creating binary of go to run
        4. giving CMD so that same can start at docker run called.
3. Added Unit Test to check UrlShortner behaviour
 
